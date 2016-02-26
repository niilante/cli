package arukas

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/manyminds/api2go/jsonapi"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"sort"
	"strings"
	"time"
)

const VERSION = "0.1.0.beta-7"

type Client struct {
	ApiURL     *url.URL
	HTTP       *http.Client
	Username   string
	Password   string
	UserAgent  string
	Debug      bool
	Output     func(...interface{})
	OutputDest io.Writer
	Timeout    time.Duration
}

var (
	client *Client
)

func (c *Client) PrintHeaderln(values ...interface{}) {
	fmt.Fprint(c.OutputDest, ToTSV(values[1:]), "\n")
}

func (c *Client) Println(values ...interface{}) {
	fmt.Fprint(c.OutputDest, ToTSV(values[1:]), "\n")
}

func (c *Client) Get(v interface{}, path string) error {
	return c.APIReq(v, "GET", path, nil)
}

func (c *Client) Patch(v interface{}, path string, body interface{}) error {
	return c.APIReq(v, "PATCH", path, body)
}

func (c *Client) Post(v interface{}, path string, body interface{}) error {
	return c.APIReq(v, "POST", path, body)
}

func (c *Client) Put(v interface{}, path string, body interface{}) error {
	return c.APIReq(v, "PUT", path, body)
}

func (c *Client) Delete(path string) error {
	return c.APIReq(nil, "DELETE", path, nil)
}

func NewClientWithOsExitOnErr() *Client {
	client, err := NewClient()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	return client
}

func NewClient() (*Client, error) {
	debug := false
	if os.Getenv("ARUKAS_DEBUG") != "" {
		debug = true
	}
	apiURL := "https://app.arukas.io/api/"
	if os.Getenv("ARUKAS_JSON_API_URL") != "" {
		apiURL = os.Getenv("ARUKAS_JSON_API_URL")
	}
	client := new(Client)
	parsedURL, err := url.Parse(apiURL)
	if err != nil {
		fmt.Println("url err")
		return nil, err
	}
	parsedURL.Path = strings.TrimRight(parsedURL.Path, "/")

	client.ApiURL = parsedURL
	client.UserAgent = "Arukas CLI (" + VERSION + ")"
	client.Debug = debug
	client.OutputDest = os.Stdout
	client.Timeout = 30 * time.Second

	if username := os.Getenv("ARUKAS_JSON_API_TOKEN"); username != "" {
		client.Username = username
	} else {
		return nil, errors.New("ARUKAS_JSON_API_SECRET is not set.")
	}

	if password := os.Getenv("ARUKAS_JSON_API_SECRET"); password != "" {
		client.Password = password
	} else {
		return nil, errors.New("ARUKAS_JSON_API_TOKEN is not set.")
	}

	return client, nil
}

func (c *Client) NewRequest(method, path string, body interface{}) (*http.Request, error) {
	var ctype string
	var rbody io.Reader

	switch t := body.(type) {
	case nil:
	case string:
		rbody = bytes.NewBufferString(t)
	case io.Reader:
		rbody = t
	case []byte:
		rbody = bytes.NewReader(t)
		ctype = "application/vnd.api+json"
	default:
		v := reflect.ValueOf(body)
		if !v.IsValid() {
			break
		}
		if v.Type().Kind() == reflect.Ptr {
			v = reflect.Indirect(v)
			if !v.IsValid() {
				break
			}
		}

		j, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		rbody = bytes.NewReader(j)
		ctype = "application/json"
	}
	requestURL := *c.ApiURL // shallow copy
	requestURL.Path += path
	if c.Debug {
		fmt.Printf("Requesting: %s %s %s\n", method, requestURL, rbody)
	}
	req, err := http.NewRequest(method, requestURL.String(), rbody)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/vnd.api+json")
	req.Header.Set("User-Agent", c.UserAgent)
	if ctype != "" {
		req.Header.Set("Content-Type", ctype)
	}
	req.SetBasicAuth(c.Username, c.Password)

	return req, nil
}

func (c *Client) APIReq(v interface{}, method, path string, body interface{}) error {
	var marshaled []byte
	var err1 error
	var req *http.Request
	if body != nil {
		var err error
		_, ok := body.(jsonapi.MarshalIdentifier)
		if ok {
			marshaled, err = jsonapi.Marshal(body)
		} else {
			marshaled, err = json.Marshal(body)
		}
		if err != nil {
			ExitCode = 1
			return err
		}

		if c.Debug {
			fmt.Println("json: ", string(marshaled))
		}
		req, err1 = c.NewRequest(method, path, marshaled)
	} else {
		req, err1 = c.NewRequest(method, path, body)
	}

	if err1 != nil {
		return err1
	}
	return c.DoReq(req, v)
}

func (c *Client) DoReq(req *http.Request, v interface{}) error {

	httpClient := c.HTTP
	if httpClient == nil {
		httpClient = &http.Client{
			Timeout: c.Timeout,
		}
	}

	res, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if c.Debug {
		fmt.Println("Status:", res.StatusCode)
		headers := make([]string, len(res.Header))
		for k, _ := range res.Header {
			headers = append(headers, k)
		}
		sort.Strings(headers)
		for _, k := range headers {
			if k != "" {
				fmt.Println(k+":", strings.Join(res.Header[k], " "))
			}
		}
		fmt.Println(string(body))
	}

	if err = checkResponse(res); err != nil {
		return err
	}

	switch t := v.(type) {
	case nil:
	case io.Writer:
		_, err = io.Copy(t, res.Body)
	default:
		err = jsonapi.Unmarshal(body, v)
		if err != nil {
			err = json.Unmarshal(body, v)
		}
	}
	return err
}

func checkResponse(res *http.Response) error {
	if res.StatusCode == 404 {
		return fmt.Errorf("The resource does not found on the server: %s", res.Request.URL)
	} else if res.StatusCode >= 400 {
		return fmt.Errorf("server returned %d", res.StatusCode)
	}
	return nil
}

func PrintTsvln(values ...interface{}) {
	fmt.Println(ToTSV(values))
}

func ToTSV(values []interface{}) string {
	str := make([]string, 0)
	for _, s := range values {
		if v, ok := s.(string); ok {
			str = append(str, string(v))
		} else {
			str = append(str, "<byte>")
		}

	}
	return strings.Join(str, "\t")
}

func SplitTSV(str string) []string {
	splitStr := strings.Split(str, "\t")
	var trimmed []string
	for _, v := range splitStr {
		trimmed = append(trimmed, strings.Trim(v, "\n"))
	}
	return trimmed
}

func removeFirstLine(str string) string {
	lines := strings.Split(str, "\n")
	return strings.Join(lines[1:], "\n")
}

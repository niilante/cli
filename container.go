package arukas

import (
	"errors"
	"fmt"
	"github.com/manyminds/api2go/jsonapi"
	"strconv"
	"strings"
	"time"
)

type PortMapping struct {
	ContainerPort int    `json:"container_port"`
	ServicePort   int    `json:"service_port"`
	Host          string `json:"host"`
}
type TaskPorts []PortMapping
type PortMappings []TaskPorts

type Env struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Envs []Env

type Port struct {
	Protocol string `json:"protocol"`
	Number   int    `json:"number"`
}

type Ports []Port

type Container struct {
	Envs         Envs         `json:"envs"`
	Ports        Ports        `json:"ports"`
	PortMappings PortMappings `json:"port_mappings,omitempty"`
	StatusText   string       `json:"status_text,omitempty"`
	ID           string
	ImageName    string    `json:"image_name"`
	CreatedAt    JSONTime  `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	App          *App
	Mem          int    `json:"mem"`
	AppId        string `json:"app_id"`
	Instances    int    `json:"instances"`
	IsRunning    bool   `json:"is_running,omitempty"`
	Cmd          string `json:"cmd"`
	Name         string `json:"name"`
	Endpoint     string `json:"end_point,omitempty"`
}

func (c Container) GetID() string {
	return string(c.ID)
}

func (c *Container) SetID(ID string) error {
	c.ID = ID
	return nil
}

func (c Container) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: "apps",
			Name: "app",
		},
	}
}

func (c Container) GetReferencedIDs() []jsonapi.ReferenceID {
	result := []jsonapi.ReferenceID{}

	if c.AppId != "" {
		result = append(result, jsonapi.ReferenceID{ID: c.AppId, Name: "app", Type: "apps"})
	}

	return result
}

func (c Container) GetReferencedStructs() []jsonapi.MarshalIdentifier {
	result := []jsonapi.MarshalIdentifier{}
	if c.App != nil {
		result = append(result, c.App)
	}
	return result
}

func (c *Container) SetToOneReferenceID(name, ID string) error {
	if name == "app" {
		if ID == "" {
			c.App = nil
		} else {
			c.App = &App{ID: ID}
		}

		return nil
	}

	return errors.New("There is no to-one relationship with the name " + name)
}

func ParseEnv(envs []string) (Envs, error) {
	var parsedEnvs Envs
	for _, env := range envs {
		kv := strings.Split(env, "=")
		parsedEnvs = append(parsedEnvs, Env{Key: kv[0], Value: kv[1]})
	}
	return parsedEnvs, nil
}

func ParsePort(ports []string) (Ports, error) {
	var parsedPorts Ports
	for _, port := range ports {
		kv := strings.Split(port, ":")
		num, err := strconv.Atoi(kv[0])
		if err != nil {
			return nil, fmt.Errorf("Port number must be numeric. Given: %s", kv[0])
		}
		parsedPorts = append(parsedPorts, Port{Number: num, Protocol: kv[1]})
	}
	return parsedPorts, nil
}

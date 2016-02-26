package arukas

import (
	"encoding/json"
	"github.com/manyminds/api2go/jsonapi"
)

type TmpJSON map[string][]map[string]interface{}

type AppSet struct {
	App       App
	Container Container
}

func (as AppSet) MarshalJSON() ([]byte, error) {
	var (
		app           []byte
		appJson       map[string]map[string]interface{}
		container     []byte
		containerJson map[string]map[string]interface{}
		marshaled     []byte
		err           error
	)

	if app, err = jsonapi.Marshal(as.App); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(app, &appJson); err != nil {
		return nil, err
	}

	if container, err = jsonapi.Marshal(as.Container); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(container, &containerJson); err != nil {
		return nil, err
	}

	data := map[string][]map[string]interface{}{
		"data": []map[string]interface{}{
			appJson["data"],
			containerJson["data"],
		},
	}

	if marshaled, err = json.Marshal(data); err != nil {
		return nil, err
	}

	return marshaled, nil
}

func SelectResources(data TmpJSON, resourceType string) TmpJSON {
	resources := make([]map[string]interface{}, 0)

	for _, v := range data["data"] {
		if v["type"] == resourceType {
			resources = append(resources, v)
		}
	}

	filtered := map[string][]map[string]interface{}{
		"data": resources,
	}
	return filtered
}

func (as *AppSet) UnmarshalJSON(bytes []byte) error {
	var (
		appBytes       []byte
		containerBytes []byte
		err            error
		data           TmpJSON
	)
	if err = json.Unmarshal(bytes, &data); err != nil {
		return err
	}

	apps := SelectResources(data, "apps")
	containers := SelectResources(data, "containers")

	if appBytes, err = json.Marshal(apps); err != nil {
		return err
	}

	if containerBytes, err = json.Marshal(containers); err != nil {
		return err
	}

	var parsedApps []App
	if err = jsonapi.Unmarshal(appBytes, &parsedApps); err != nil {
		return err
	}

	var parsedContainers []Container
	if err = jsonapi.Unmarshal(containerBytes, &parsedContainers); err != nil {
		return err
	}

	as.App = parsedApps[0]
	as.Container = parsedContainers[0]
	return nil
}

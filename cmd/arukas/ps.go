package main

import (
	"log"

	arukas "github.com/arukasio/cli"
)

func listContainers(listAll bool, quiet bool) {
	var parsedContainer []arukas.Container
	var filteredContainer []arukas.Container
	client := arukas.NewClientWithOsExitOnErr()

	if err := client.Get(&parsedContainer, "/containers"); err != nil {
		log.Println(err)
		ExitCode = 1
		return
	}

	if listAll {
		filteredContainer = parsedContainer
	} else {
		for _, container := range parsedContainer {
			if container.StatusText == "running" {
				filteredContainer = append(filteredContainer, container)
			}
		}
	}

	if quiet {
		for _, container := range filteredContainer {
			client.Println(nil, container.ID)
		}
	} else {
		client.PrintHeaderln(nil, "CONTAINER ID", "IMAGE", "COMMAND", "CREATED", "STATUS", "ARUKAS DOMAIN", "ENDPOINT")
		for _, container := range filteredContainer {
			client.Println(nil, container.ID, container.ImageName, container.Cmd, container.CreatedAt.String(),
				container.StatusText, container.ArukasDomain, container.Endpoint)
		}
	}
}

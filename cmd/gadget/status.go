package main

import (
	"fmt"
)

// Process the build arguments and execute build
func GadgetStatus(args []string, g *GadgetContext) {
	g.loadConfig()
	ensureKeys()

	client, err := gadgetLogin(gadgetPrivKeyLocation)

	if err != nil {
		panic(err)
	}

	fmt.Println("[GADGT]  Retrieving status:")
	
	stagedContainers,_ := findStagedContainers(args, append(g.Config.Onboot, g.Config.Services...))
	
	for _, container := range stagedContainers {
		commandFormat := `docker ps -a --filter=ancestor=%s --format "{{.Image}} {{.Command}} {{.Status}}"`
		cmd := fmt.Sprintf(commandFormat, container.ImageAlias)
		runRemoteCommand(client, cmd)
		if err != nil {
			panic(err)
		}
	}
}

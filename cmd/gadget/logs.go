package main

import (
	"fmt"
)

// Process the build arguments and execute build
func GadgetLogs(args []string, g *GadgetContext) {
	g.loadConfig()
	ensureKeys()

	client, err := gadgetLogin(gadgetPrivKeyLocation)

	if err != nil {
		panic(err)
	}

	fmt.Println("[GADGT]  Retrieving logs:")
	
	for _, onboot := range g.Config.Onboot {
		commandFormat := `docker logs %s`
		cmd := fmt.Sprintf(commandFormat, onboot.Alias)
		runRemoteCommand(client, cmd)
		if err != nil {
			panic(err)
		}
	}
}

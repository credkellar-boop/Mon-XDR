package main

import (
	"log"

	"github.com/credkellar-boop/Mon-XDR/pkg/action"
	"github.com/gorilla/websocket"
)

func ListenForCommands(serverURL string) {
	conn, _, err := websocket.DefaultDialer.Dial(serverURL, nil)
	if err != nil {
		log.Fatalf("Connection to backend failed: %v", err)
	}
	defer conn.Close()

	for {
		var cmd struct {
			Action string `json:"action"` // e.g., "QUARANTINE", "KILL_PROCESS"
			Path   string `json:"path"`
		}

		// Blocking read: wait for command from the server
		err := conn.ReadJSON(&cmd)
		if err != nil {
			log.Println("Connection dropped:", err)
			break
		}

		// Execute the command locally using the exported public actions
		switch cmd.Action {
		case "QUARANTINE":
			action.QuarantineFile(cmd.Path)
		case "KILL_PROCESS":
			action.KillProcess(cmd.Path)
		}
	}
}

package main

import (
    "github.com/gorilla/websocket"
    "log"
)

func ListenForCommands(serverURL string) {
    conn, _, err := websocket.DefaultDialer.Dial(serverURL, nil)
    if err != nil {
        log.Fatal("Connection to backend failed:", err)
    }

    for {
        var cmd struct {
            Action string `json:"action"` // e.g., "QUARANTINE"
            Path   string `json:"path"`
        }
        
        // Blocking read: wait for command from the server
        err := conn.ReadJSON(&cmd)
        if err != nil {
            log.Println("Connection dropped:", err)
            break
        }

        // Execute the command locally
        switch cmd.Action {
        case "QUARANTINE":
            quarantineFile(cmd.Path)
        case "KILL_PROCESS":
            killProcess(cmd.Path)
        }
    }
}

package action

import "log"

// QuarantineFile must be capitalized to be exported!
func QuarantineFile(path string) {
	log.Printf("[Action] Quarantining file or payload at: %s", path)
	// Add your cross-platform quarantine logic here
}

// KillProcess must be capitalized to be exported!
func KillProcess(processName string) {
	log.Printf("[Action] Killing process: %s", processName)
	// Add your process termination logic here
}

package pyscript

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
)

func runScript() {
	input := map[string]interface{}{
		"name": "Peter",
		"age":  31,
	}

	// Marshal to JSON
	jsonData, _ := json.Marshal(input)

	// Call Python script
	cmd := exec.Command("python3", "script.py")

	// Send JSON to Python via stdin
	cmd.Stdin = bytes.NewReader(jsonData)

	// Capture Python output
	out, _ := cmd.Output()

	// Parse JSON response
	var result map[string]interface{}
	json.Unmarshal(out, &result)

	fmt.Println("Python responded:", result)
}

package main

import (
	"fmt"
	"strings"
)

func processLogs(logs []string) map[string]int {
	errorEvents := make(map[string]int)
	for _, log := range logs {
		tokens := strings.Split(log, " ")
		if len(tokens) < 3 {
			continue
		}
		ip := tokens[2]
		response := tokens[len(tokens)-1]
		if response[0] == '4' || response[0] == '5' {
			event := fmt.Sprintf("[%s] %s", response, ip)
			errorEvents[event]++
		}
	}

	return errorEvents

}

func main() {
	//example data
	logs := []string{
		"[2023/02/06 13:00:00] 192.168.3.92 : POST HTTP/1.1 404",
		"[2023/02/06 13:00:00] 192.168.4.22 : GET HTTP/1.1 500"}
	errorsFromLogs := processLogs(logs)
	fmt.Printf("Number of errors: %d\n", len(errorsFromLogs))
	for event, _ := range errorsFromLogs {
		fmt.Println(event)
	}

}

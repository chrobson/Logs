## Log Error Finder

This is a simple Go function that takes a string array of log lines as input and returns a map of error events in the format `[<Response_Code>] <IP_address>`.

### Function Signature

```go
func processLogs(logs []string) map[string]int
```

### Inputs

The function takes a single argument:

* `logs`: A string array containing log lines in the following format:
  ```
  [<timestamp>] <IP_address> : <HTTP_method> HTTP/1.1 <Response_Code>
  ```

### Outputs

The function returns a map `errorEvents` containing the error events in the following format:

* Key: `[<Response_Code>] <IP_address>`
* Value: The number of occurrences of the error event in the input logs.

### Example Usage
Program can be accesed from Go playground: https://go.dev/play/p/5w8gMR8bbRo

Here's an example usage of the function:

```go
package main

import (
    "fmt"
)

func main() {
    logs := []string{
        "[2023/02/06 13:01:34] 192.168.3.72 : POST HTTP/1.1 403",
        "[2023/02/06 13:01:35] 192.168.3.72 : GET HTTP/1.1 200",
        "[2023/02/06 13:01:36] 192.168.3.72 : PUT HTTP/1.1 500",
    }
    errorEvents := processLogs(logs)
    fmt.Printf("Number of errors events: %d\n", len(errorEvents))
    for event, _ := range errorEvents {
        fmt.Printf("%s\n", event)
    }
}
```

In this example, the function is called with an array `logs` containing three log lines. The output of the program is as follows:

```
Number of errosrs events: 2
[403] 192.168.3.72
[500] 192.168.3.72
```

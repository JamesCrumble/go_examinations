package parts

import (
	"fmt"
	"io"
	"net/http"
)

// TODO: Implement another methods with POST request

func read_from_response_buffer(response *http.Response) string {

	content, _ := io.ReadAll(response.Body)

	fmt.Printf("CONTENT TYPE as original \"content\" (io.ReadAll) => %T\n", content)
	fmt.Printf("CONTENT TYPE as casted to string \"(string)(content)\" => %T\n", (string)(content))
	fmt.Printf("CONTENT TYPE as argumented to string \"string(content)\" => %T\n", string(content))

	return string(content)
}

func HttpPartExecute(url *string) {
	printInfo("HTTP PART")

	response, err := http.Get(*url)
	if err != nil {
		fmt.Printf("ERROR: => \"%s\"\n", err)
	} else {
		fmt.Printf(
			"CONTENT FROM \"%s\" URL => %s\nSTATUS CODE: \"%d\"\n",
			*url,
			read_from_response_buffer(response),
			response.StatusCode,
		)

	}

	printInfo("HTTP PART")
}

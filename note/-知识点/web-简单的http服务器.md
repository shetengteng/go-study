```go
package main

import (
	"fmt"
	"io"
	"net/http"
)

func firstPage(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>hello</h1>")
}

func main() {

	fmt.Println("start")
	http.HandleFunc("/", firstPage)
	http.ListenAndServe(":8080", nil)
}
```


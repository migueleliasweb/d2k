package docker141

import (
	"fmt"
	"net/http"
	"strings"
)

func Docker141Handler(rw http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path, r.Method)

	splittedPath := strings.Split(r.URL.Path, "/")

	if splittedPath[1] == "_ping" || splittedPath[1] == "ping" || (len(splittedPath) >= 3 && strings.HasPrefix(splittedPath[1], "v1.") && splittedPath[1] != "v1.41") {
		splittedPath[1] = "v1.41"

		fmt.Println("rewritten as:", strings.Join(splittedPath, "/"))
		r.URL.Path = strings.Join(splittedPath, "/")
	}

	// docker cli uses this header to figure out the version of the daemon
	// instead of using `/version` *sigh*
	rw.Header().Add("Api-Version", "v1.41")
	fmt.Println("-----------------------------------")
}

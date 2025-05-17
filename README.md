# go-mixi2
[![Go Reference](https://pkg.go.dev/badge/github.com/matsuu/go-mixi2.svg)](https://pkg.go.dev/github.com/matsuu/go-mixi2)

Unofficial mixi2 client in Go.

## Usage

* Access https://mixi.social/ in your browser and log in.
* After logging in, open the Developer Tools (usually by pressing F12) and navigate to the Network tab.
* While keeping the Network tab open, make a test post.
* Find and select the CreatePost request in the Network tab list.
* In the Request Headers section for the selected CreatePost request, find the following:
    * The value of auth\_token within the Cookie header.
    * The value of the X-Auth-Key header.
    * The value of the X-Mercury-User-Agent header.
* Pass the values you have identified as follows:

```go
package main

import (
	"context"
	"fmt"
	"os"

	"connectrpc.com/connect"

	"github.com/matsuu/go-mixi2"
	"github.com/matsuu/go-mixi2/gen/com/mixi/mercury/api"
)

func createPost(ctx context.Context, authKey, authToken, userAgent, text string) error {
	client := mixi2.NewClient(mixi2.WithAuth(authKey, authToken, userAgent))
	post := api.CreatePostRequest{
		Text: text,
	}
	resp, err := client.CreatePost(ctx, connect.NewRequest(&post))
	if err != nil {
		return fmt.Errorf("failed to CreatePost: %w", err)
	}
	fmt.Println(resp.Msg.String())

	return nil
}

func main() {
    authKey := os.Getenv("MIXI2_AUTH_KEY")
    authToken := os.Getenv("MIXI2_AUTH_TOKEN")
    ua := os.Getenv("MIXI2_USER_AGENT")
    text := "42"
	if err := createPost(context.Background(), authKey, authToken, ua, text); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
```

## Build

[Install Connect tools](https://connectrpc.com/docs/go/getting-started/)

```
buf generate
```

## References

* [mixi2](https://mixi.social/)

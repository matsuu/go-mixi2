package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"

	"connectrpc.com/connect"

	"github.com/matsuu/go-mixi2"
	"github.com/matsuu/go-mixi2/gen/com/mixi/mercury/api"
)

func createPost(ctx context.Context, authKey, authToken, userAgent, text string, community string) error {
	client := mixi2.NewClient(mixi2.WithAuth(authKey, authToken, userAgent))
	post := api.CreatePostRequest{
		Text: text,
	}
	if community != "" {
		post.CommunityId = &community
	}
	resp, err := client.CreatePost(ctx, connect.NewRequest(&post))
	if err != nil {
		return fmt.Errorf("failed to CreatePost: %w", err)
	}
	fmt.Println(resp.Msg.String())

	return nil
}

func main() {
	authKey := flag.String("key", "", "X-Auth-Key in Request Header")
	authToken := flag.String("token", "", "auth_token in Cookie")
	ua := flag.String("ua", "", "X-Mercury-User-Agent in Request Header")
	text := flag.String("text", "", "post text")
	community := flag.String("community", "", "community id")
	flag.Parse()

	var missings []string
	if *authKey == "" {
		missings = append(missings, "key")
	}
	if *authToken == "" {
		missings = append(missings, "token")
	}
	if *ua == "" {
		missings = append(missings, "ua")
	}
	if *text == "" {
		missings = append(missings, "text")
	}
	if len(missings) > 0 {
		fmt.Fprintf(os.Stderr, "no %s found\n", strings.Join(missings, ", "))
		flag.PrintDefaults()
		os.Exit(1)
	}

	if err := createPost(context.Background(), *authKey, *authToken, *ua, *text, *community); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

package mixi2

import (
	"encoding/base64"
	"testing"

	"github.com/matsuu/go-mixi2/gen/com/mixi/mercury/api"

	"google.golang.org/protobuf/proto"
)

func TestUnmarshal(t *testing.T) {
	b64 := []byte(`Cr0DCiQwZWY5NThkOS03NGY5LTRkZjctOGFhYy0zZTBkODQwMjU4MmQSEzE3NDM1ODc3MDk3Mzk3NDYzODgaBgj9krS/BiIkMDhiYjViMDItMGM5Mi00OTM3LWIzNmQtOTE2ZmI2MzUzYjA3OssCCiQwNTI1NTA4OS03NzYxLTRhZTUtOWE5ZC1iMDcyMGQxNWZhMjEYAiADMp4CCmxodHRwczovL21lZGlhLm1peGkuc29jaWFsL2ltYWdlcy8wOGJiNWIwMi0wYzkyLTQ5MzctYjM2ZC05MTZmYjYzNTNiMDcvbC8wNTI1NTA4OS03NzYxLTRhZTUtOWE5ZC1iMDcyMGQxNWZhMjESCmltYWdlL3dlYnAYgAQggAQqbGh0dHBzOi8vbWVkaWEubWl4aS5zb2NpYWwvaW1hZ2VzLzA4YmI1YjAyLTBjOTItNDkzNy1iMzZkLTkxNmZiNjM1M2IwNy9zLzA1MjU1MDg5LTc3NjEtNGFlNS05YTlkLWIwNzIwZDE1ZmEyMTIKaW1hZ2Uvd2VicDiABECABEocTEdPNDl6TXt0N0lVfnFheVJqeHUtO3h1b2ZXQmIEdGVzdA==`)
	raw := make([]byte, base64.StdEncoding.DecodedLen(len(b64)))
	n, err := base64.StdEncoding.Decode(raw, b64)
	if err != nil {
		t.Fatal("Failed to decode response", err)
	}
	raw = raw[:n]

	postResponse := &api.CreatePostResponse{}
	if err := proto.Unmarshal(raw, postResponse); err != nil {
		t.Fatal("Failed to read response", err)
	}
	println(postResponse.String())
}

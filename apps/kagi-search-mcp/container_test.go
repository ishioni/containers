package main

import (
	"context"
	"net/http"
	"testing"

	"github.com/ishioni/containers/testhelpers"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func Test(t *testing.T) {
	ctx := context.Background()
	image := testhelpers.GetTestImage("ghcr.io/home-operations/kagi-search-mcp:rolling")

	const port = "8080/tcp"

	c, err := testcontainers.Run(ctx, image,
		testcontainers.WithEnv(map[string]string{"KAGI_API_KEY": "dummy"}),
		testcontainers.WithCmd("--http", "--host", "0.0.0.0", "--port", "8080"),
		testcontainers.WithExposedPorts(port),
		testcontainers.WithWaitStrategy(
			wait.ForListeningPort(port),
			wait.ForHTTP("/mcp").WithPort(port).WithStatusCodeMatcher(func(status int) bool {
				return status == http.StatusNotAcceptable
			}),
		),
	)
	testcontainers.CleanupContainer(t, c)
	require.NoError(t, err)
}

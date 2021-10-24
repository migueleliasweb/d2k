package docker141

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func TestFooBar(t *testing.T) {
	ctx := context.Background()

	cli, err := client.NewClient("127.0.0.1", "", http.DefaultClient, map[string]string{})
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		fmt.Println(container.ID)
	}

	cli.ServerVersion(ctx)
}

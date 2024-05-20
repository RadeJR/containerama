package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func main() {
	client, _ := client.NewClientWithOpts(client.FromEnv)
	reader, err := client.ContainerLogs(context.Background(), "ebf6a08b863a32e21d49b3e574cf6acf82b6cddddfa3b113040a1461ccb7ad52", container.LogsOptions{ShowStdout: true, ShowStderr: true, Follow: true})
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(os.Stdout, reader)
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}
}

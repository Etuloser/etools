package example

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func getPublicIp() (string, error) {
	resp, err := http.Get("http://ifconfig.me")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func List() ([]ContainerInfo, error) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}
	defer cli.Close()

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		return nil, err
	}
	containerInfo := []ContainerInfo{}
	publicIp, err := getPublicIp()
	if err != nil {
		return nil, err
	}
	for _, container := range containers {
		for _, port := range container.Ports {
			if strings.Contains(port.IP, "0.0.0.0") {
				uri := fmt.Sprintf("http://%s:%d", publicIp, port.PublicPort)
				info := ContainerInfo{
					Name: strings.Split(container.Names[0], "/")[1],
					URI:  uri,
				}
				containerInfo = append(containerInfo, info)
			}
		}
	}
	return containerInfo, nil
}

type ContainerInfo struct {
	Name string
	URI  string
}

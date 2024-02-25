package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/docker/cli/cli/connhelper"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

func main() {
	host := os.Getenv("DOCKER_HOST")
	helper, err := connhelper.GetConnectionHelper(host)

	if err != nil {
		panic(err)
	}

	httpClient := &http.Client{
		// No tls
		// No proxy
		Transport: &http.Transport{
			DialContext: helper.Dialer,
		},
	}

	var clientOpts []client.Opt

	clientOpts = append(clientOpts,
		client.WithHTTPClient(httpClient),
		client.WithHost(helper.Host),
		client.WithDialContext(helper.Dialer),
	)

	version := os.Getenv("DOCKER_API_VERSION")

	if version != "" {
		clientOpts = append(clientOpts, client.WithVersion(version))
	} else {
		clientOpts = append(clientOpts, client.WithAPIVersionNegotiation())
	}
	cli, err := client.NewClientWithOpts(clientOpts...)
	if err != nil {
		panic(err)
	}

	ctx, _ := context.WithCancel(context.Background())
	log.Printf("starting proxy for host %s", host)

	filter := filters.NewArgs()
	filter.Add("type", "network")
	filter.Add("event", "connect")
	filter.Add("event", "disconnect")
	filter.Add("type", "container")
	filter.Add("event", "attach")
	filter.Add("event", "export")
	filter.Add("event", "start")
	filter.Add("event", "die")

	msg, errChan := cli.Events(ctx, types.EventsOptions{
		Filters: filter,
	})

	for {
		select {
		case err := <-errChan:
			panic(err)
		case event := <-msg:
			fmt.Printf("event %+v\n", event)
			fmt.Println("reloading")
			reload()
		}
	}
}

// Listen for docker events
func reload() {
	fmt.Println("…done")
}

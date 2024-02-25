# Docker Proxy

This project should implement a proxy, that listens to a remote or local docker host.  
Everytime, a new container opens a port on that remote host, the local proxy will open that port as well, so that you alway can can access the remote service via a local port (e.g. nginx on `http://HOST:2023` will be accessible via `http://localhost:2023` as well.).

We are already able to connect to that remote docker host an receiving events.  
Next step is to find out, how to detect the newly opened ports.  
Then we have to setup new proxy service in demand.  

## RUN

* Copy `.envrc.examlpe` to `.envrc` and update accordingly.
* Run the service with 
  > go run cmd/proxy/main.go


## Links

* https://blog.programster.org/use-remote-docker-host-with-ssh
* https://docs.docker.com/reference/cli/docker/system/events/
* https://stackoverflow.com/questions/52816890/uptodate-list-of-running-docker-containers-stated-in-an-exported-golang-variable
* https://medium.com/@nathanpbrophy/write-a-sample-port-forwarder-in-golang-2748309c1e80

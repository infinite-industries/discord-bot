set dotenv-load

binary := "discord-bot"

tag := `grep ^FROM build/Dockerfile | cut -d: -f2`
image := "infinite-industries/" + binary

registry := "ghcr.io"
registry_user := "jswank"

# this environment variable will be passed to podman login
pass_var := "REGISTRY_PASSWORD"  

# compile and run
all: build run

# display help
help: 
  just -l

# remove binaries + container
@clean: clean-bin clean-container

# build cmd
@build cmd=binary:
  mkdir -p bin
  go build -o bin/{{cmd}} ./cmd/{{cmd}}

# runs cmd (default: bin/{{ binary }}). 
run cmd=binary:
  bin/{{cmd}} --loglevel=debug

@clean-bin:
  rm -f bin/*

# build a docker image
build-image:
  podman build -t {{image}} -f build/Dockerfile .

# run a container via podman
run-container: 
  podman run --rm -P --env-file .env --name {{ binary }} {{ image }}

# remove the container
@clean-container:
  podman rm -f 

# publish the image to the default registry
publish: _publish

# publish the image to ghcr.io
publish-ghcr: (_publish "ghcr.io")

# publish the image to quay.io
publish-quay: (_publish "quay.io")

# publish the image to docker.io
publish-docker: (_publish "docker.io")

_publish registry=registry user=registry_user alt_tag=tag:
  @ echo "${{ pass_var }}" | podman login {{registry}} -u {{user}} --password-stdin 
  @ podman tag {{image}}:{{tag}} {{registry}}/{{image}}:{{alt_tag}}
  @ podman push {{registry}}/{{image}}:{{alt_tag}}
  @ podman logout {{registry}}

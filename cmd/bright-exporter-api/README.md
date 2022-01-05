# `bright-export-api`

This is a simple Prometheus exporter for the Bright API.

You can either build the binary with `go build` and run it, or use the provided `Dockerfile`.

## Configuration

The only configuration required is provided by environment variables:

* `BRIGHT_USERNAME` - your username for the Bright API.
* `BRIGHT_PASSWORD` - your password for the Bright API.
* `PORT` - An optional port to run the exporter on, defaults to `9998` if not provided.

## Docker

### Building

To build the container simply run `make` in this directory.

### Running

If the container has built successfully, you can run it with something like:

    docker run \
        --rm 
        -e BRIGHT_USERNAME="<your bright API username>" \
        -e BRIGHT_PASSWORD="<your bright API password>" \
        -p 9998:9998 \
        bright-exporter-api:latest
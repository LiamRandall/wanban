# HTTP Hello World

This is a simple TinyGo Wasm example that responds with a "Hello World" message for each request.

## Prerequisites

- `go` 1.23
- `tinygo` 0.33
- [`wash`](https://wasmcloud.com/docs/installation) 0.35.0
- `wasmtime` 25.0.0 (if running with wasmtime)

## Building

```bash
wash build
```

## wanban Example API Usage

This code is currently re-entrant, so it is essentially stateless - but here are the API Commands supported at the moment:


Create a board:
```
curl -X POST -d '{"title":"My Board"}' -H "Content-Type: application/json"  http://localhost:8000/boards
```

List all boards:
```
    curl http://localhost:8000/boards
```

Create a list in a board:
```
# Replace BOARD_ID with the ID you got from creating a board
curl -X POST -d '{"title":"To Do"}' -H "Content-Type: application/json" http://localhost:8000/boards/BOARD_ID/lists
```

Add a card to a list:
```
# Replace BOARD_ID and LIST_ID accordingly
curl -X POST -d '{"title":"Finish writing code"}' -H "Content-Type: application/json"  http://localhost:8000/boards/BOARD_ID/lists/LIST_ID/cards
```

## Running with wasmtime

You must have wasmtime 25.0.0 for this to work. Make sure to follow the build step above first.

```bash
wasmtime serve -Scommon ./build/http_hello_world_s.wasm
```

## Running with wasmCloud

Make sure to follow the build steps above, and replace the file path in [the wadm manifest](./wadm.yaml) with the absolute path to your local built component.

```shell
wash up -d
wash app deploy ./wadm.yaml
curl http://localhost:8000
```

## Adding Capabilities

To learn how to extend this example with additional capabilities, see the [Adding Capabilities](https://wasmcloud.com/docs/tour/adding-capabilities?lang=tinygo) section of the wasmCloud documentation.

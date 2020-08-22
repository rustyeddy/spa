# spa - The Single Page App Server

spa was built to server up Single Page Applications (SPA) on small
embedded devices, for example works great serving up dashboards on
devices controlled by a Raspberry PI.

## Build

```bash
% go build
```

## Serve HTML Websites 

To run the static website directly from memory do this:

```
% spa
```

If you want to run a specific website that contains a directory with
an index.html file in it, we can serve that up like this:

```
% spa -pub examples/static
```

That should produce a website that looks like: (Todo put a screen
shoot of our super simple static website).

## REST API

- GET /api/health


## Web Socket

We have a websocket connected between our app and the server. We'll merge this branch then work on adding components that communicate over the websocket.

All websocket communications are JSON formatted for easy debugging and processing. In heavy scenarios a binary version may fair better such as Protocol Buffers or TLV.

### Supported message types

- echo (client sends string to server, server echos string back to client)

### TODO Implement

- time (server sends client time structs periodically)
- quote (client enters quote in form, quote sent to server and recorded)

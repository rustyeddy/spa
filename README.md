# spa - The Single Page App Server

spa was built to server up Single Page Applications (SPA) on small
embedded devices, for example works great serving up dashboards on
devices controlled by a Raspberry PI.

## Build

```bash
% go build
```

## Run

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


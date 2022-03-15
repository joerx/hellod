# hellod

Hello World web server. For testing random shit.

## Building

Docker login:

```sh
echo $GITHUB_TOKEN | docker login ghcr.io -u USERNAME --password-stdin
```

Build and publish:

```sh
make push
```

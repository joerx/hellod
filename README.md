# hellod

Hello World web server. For testing random shit.

## Building && Running

```
make && out/hellod
```

## Docker

Docker login:

```sh
echo $GITHUB_TOKEN | docker login ghcr.io -u USERNAME --password-stdin
```

Docker build and publish:

```sh
make docker-push
```

# Kfn

POC of a FaaS build tool to create knative services starting from functions

## Building

Look at https://github.com/containers/buildah/blob/master/install.md for dep dependency installation

```bash
make build
```

## Example

Try to deploy `example-fn.js`

```bash
./kfn run fn.js
```

## Kfn image

To build the image, run:

```bash
make image
```

To run kfn image (assuming in your workdir you have a directory `mycontainer` and the function in the directory `fn`):

```bash
podman run --net=host --security-opt label=disable --security-opt seccomp=unconfined --privileged --device /dev/fuse:rw -v $(pwd)/mycontainer:/var/lib/containers:Z -v "./fn:/home" kfn --verbose run /home/example-fn.js
```




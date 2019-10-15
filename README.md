# Kfn

POC of a FaaS build tool to create knative services starting from functions

## Building

Look at https://github.com/containers/buildah/blob/master/install.md for dep dependency installation

```shell script
make build
```

## Install

```shell script
make install
```

## Example

Try to deploy `example-fn.js`

```shell script
./kfn run fn.js
```

## Kfn image

To build the image, run:

```shell script
make image
```

To run kfn image (assuming in your workdir you have a directory `mycontainer` and the function in the directory `fn`):

```shell script
podman run --net=host --security-opt label=disable --security-opt seccomp=unconfined --privileged --device /dev/fuse:rw -v $(pwd)/mycontainer:/var/lib/containers:Z -v "./fn:/home" kfn --verbose run /home/example-fn.js
```

## Commands available

* `kfn init`: Init a function in the working directory
* `kfn clean [function]` or `kfn clean --global`: Clean target directory of specified function or all `.kfn` directory
* `kfn edit [function] [editor]`: Edit the function with the specified editor
* `kfn build`: Build the specified function and push to the specified registry
* `kfn run`: Build, push and run the specified function

## Functions documentation

### Dependencies

To add dependencies, add a comment:

```
// kfn:dependency primal 0.2.3
```

Kfn will add the dependency in the specific build manifest

### Rust

#### Requirements

For Rust, you'll need [musl libc](https://www.musl-libc.org/how.html) and the corresponding target for rustc. 
This target is required to static link the libc, reducing the effective image size to just the required libraries.

You can install the target with rustup:

```shell script
rustup target add x86_64-unknown-linux-musl
```

#### Build profile

To speedup during the development the cargo profile from release to debug, add this comment in your function:

```
// kfn:build-dev true
```

#### Compilation cache

If you want to speedup the compilation across various functions, you can use [sccache](https://github.com/mozilla/sccache). To enable it when you compile your functions, add the comment:

```
// kfn:build-env RUSTC_WRAPPER=sccache
```

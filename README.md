# Kfn

POC of a FaaS build tool to create knative services starting from functions

## Building

Look at https://github.com/containers/buildah/blob/master/install.md for dep dependency installation

```bash
go build -o kfn main.go
```

## Example

Try to deploy `example-fn.js`

```bash
./kfn run fn.js
```



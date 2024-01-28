# Monoone

This is _the one_ monorepo that combines:

- Protobuf definitions
- Go backend implementation
- gRPC gateway
- gRPC Typescript Client
- React App

## Running this setup

To run everything - gateway, backend, react app - you need to run three different commands:

1. Start the backend:
    ```
    mage run:api
    ```
1. Start the gateway:
    ```
    mage run:gateway
    ```
1. Start the app (supports hot reloading via Vite):
    ```
    mage run:app
    ```

## Tooling

This repository is set up with the following overall tooling:

- Mage for running make-like commands
- Bazel as build system for Go / Protobuf
- Buf to generate code from Protobuf files
- Yarn for frontend code

## Development

### Initial Setup

Ensure the following tools are installed locally:
1. Bazelisk:
    ```
    brew install bazelisk
    ```
2. Buf:
    ```
    brew install bufbuild/buf/buf
    ```
3. Yarn v2 is enabled:
    ```
    corepack enable
    ```

Make sure to install all recommended VSCode plugins when using this repository. Also for doing any changes to the frontend code make sure you have setup
the Yarn SDKs correctly by running:
```
yarn dlx @yarnpkg/sdks vscode
```

This should also automatically configure TypeScript correctly. After running this command re-open VSCode and open any TypeScript file. You should be prompted
to select your "workspace" TypeScript installation.

### Protobuf

All protobuf definitions are stored inside the [protobufs](/protobufs/) directory.

When adding dependencies to external protobufs make sure to reference them in the Buf config file [buf.yaml](/protobufs/buf.yaml).

Whenver you changed the protobuf definitions:
1. Make sure all Bazel build files are up to date:
    ```
    mage bazel:gazelle
    ```
2. Generate new code for the updated protos:
    ```
    mage buf:generate
    ```

The protobuf generation does the following:
1. Generate corresponding Go code for the protobuf messages (in [gen/proto/go](/gen/proto/go/)).
2. Generate gRPC service stubs for Go (in [gen/proto/go](/gen/proto/go/)).
3. Generate gRPC Gateway code in Go (in [gen/proto/go](/gen/proto/go/)).
4. Generate an OpenAPI v2 specification (in [gen/proto/openapi](/gen/proto/openapi/)).
5. Generate a TypeScript client from the OpenAPI spec (in [packages/api/src/api.ts](/packages/api/src/api.ts)) with `swagger-typescript-api`.

> We use `swagger-typescript-api` via the OpenAPI intermediate spec as the default gRPC web client generators don't really
> support option / required fields and validation rules so that the generated TypeScript types are harder to work with.

### Go

There are two main directories for Go code:
- [gateway](/gateway/) for the gRPC gateway server
- [apiserver](/apiserver/) for the regular business logic implementation

After any change to a Go file just make sure to run
```
mage bazel:gazelle
```
to ensure that all Bazel build files are up-to-date.

When adding new dependencies for Go, just add the dependency as usual:
```
go get <dependency>
```
Afterwards just make sure to run:
```
mage bazel:gazelle
```
This will also automatically run `go mod tidy`.

### Frontend

We use Yarn workspaces to have multiple packages in this repository. The [packages/api](/packages/api/) package is purely to contain the
generated TypeScript client and should not need manual intervention.

The React app is fully contained in [packages/app](/packages/app/) and is a default React app setup with Vite.

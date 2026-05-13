# KV Store

## Summary

This is a simple implemetation of a command line key value storage server (kind of like redis).
The program itself acts as a server that can store string keys and values in memory and serve them
concurrently to multiple clients connecting via tcp.

## Building

Run the following command in the root of the project.

```bash
make build
```

It will create an executable called `kv-store` in a directory called `out/`.

## Usage

The most simple way to run the app is the following:

```bash
kv-store
```

This will start the server at a default port. The port of the server can be explicitely
specified with:

```bash
kv-store --port 8080
```

From another terminal connect using:

```bash
nc localhost <port>
```

Now you can send the commands to the server. All commands are separated by newlines (\n).

## Available Commands

- `SET <key> <value>`: sets the key and value in the storage
- `GET <key>`: prints the value for the specified key or (nil)
- `DEL <key>`: deletes the value for the specified key
- `EXISTS <key>`: checks if the specified key exists in the storage
- `KEYS`: lists all keys present in the storage
- `QUIT`: quits the connection

# Go Challenge: Tiny Redis-like Key-Value Server

## Goal

Build a small Redis-like key-value server in Go.

The program should run as a TCP server, accept multiple client connections, parse simple text commands, and store key-value data in memory.

This project is designed to teach practical Go concepts:

- Goroutines

- Mutexes

- TCP networking

- Command parsing

- Error handling

- Maps

- Structs and methods

- Interfaces

- Unit testing

- Graceful shutdown

- Command-line flags

---

## Example Usage

Start the server:

```bash

go run ./cmd/server --port 6380

```

Connect with `netcat`:

```bash

nc localhost 6380

```

Example client session:

```text

PING

PONG

SET language go

OK

GET language

go

EXISTS language

true

DEL language

OK

GET language

(nil)

QUIT

BYE

```

---

## Suggested Project Structure

```text

kv-server/

  go.mod

  README.md

  CHALLENGE.md

  cmd/

    server/

      main.go

  internal/

    store/

      store.go

      store_test.go

    protocol/

      parser.go

      parser_test.go

    server/

      server.go

      server_test.go

```

---

## Core Requirements

### 1. Server Startup

The program must start a TCP server.

Requirements:

- The server listens on a configurable port.

- The default port should be `6380`.

- The port should be configurable with a command-line flag.

Example:

```bash

go run ./cmd/server --port 6380

```

The server should print a startup message similar to:

```text

Listening on :6380

```

---

### 2. TCP Client Handling

The server must accept multiple client connections.

Requirements:

- Each client connection is handled independently.

- Multiple clients can connect at the same time.

- One blocked or slow client must not block other clients.

- Each connection should be handled using a goroutine.

- The server should close client connections cleanly when the client sends `QUIT` or disconnects.

---

### 3. In-Memory Store

Implement a thread-safe in-memory key-value store.

The store should support string keys and string values.

Suggested type:

```go

type Store struct {

    mu   sync.RWMutex

    data map[string]string

}

```

Required methods:

```go

Set(key, value string)

Get(key string) (string, bool)

Delete(key string) bool

Exists(key string) bool

Keys() []string

```

Requirements:

- The store must be safe for concurrent use.

- Reads should not corrupt writes.

- Writes should not corrupt reads.

- Access to the map must be protected with a mutex.

---

### 4. Command Protocol

The server should use a simple line-based protocol.

Each command is sent as one line of text.

Example:

```text

SET name Hakon

GET name

```

Commands are separated by newline characters.

Input should be parsed into a command structure.

Suggested type:

```go

type Command struct {

    Name string

    Args []string

}

```

Requirements:

- Command names should be case-insensitive.

- Extra whitespace should be ignored.

- Empty lines should return an error or be ignored consistently.

- Invalid commands should return a clear error response.

- The server must not crash on malformed input.

---

## Required Commands

### `PING`

Checks whether the server is alive.

Input:

```text

PING

```

Output:

```text

PONG

```

---

### `SET`

Stores a key-value pair.

Input:

```text

SET <key> <value>

```

Example:

```text

SET language go

```

Output:

```text

OK

```

Requirements:

- The key must not be empty.

- The value must not be empty.

- Existing keys should be overwritten.

- For the core version, values may be limited to a single word.

---

### `GET`

Returns the value for a key.

Input:

```text

GET <key>

```

Example:

```text

GET language

```

Output when the key exists:

```text

go

```

Output when the key does not exist:

```text

(nil)

```

---

### `DEL`

Deletes a key.

Input:

```text

DEL <key>

```

Example:

```text

DEL language

```

Output when the key existed:

```text

OK

```

Output when the key did not exist:

```text

(nil)

```

---

### `EXISTS`

Checks whether a key exists.

Input:

```text

EXISTS <key>

```

Output when the key exists:

```text

true

```

Output when the key does not exist:

```text

false

```

---

### `KEYS`

Returns all stored keys.

Input:

```text

KEYS

```

Example output:

```text

language

name

session

```

Requirements:

- If there are no keys, return an empty response or `(empty)`.

- Keys do not need to be sorted in the first version.

- Sorting keys is acceptable as an improvement.

---

### `QUIT`

Closes the client connection.

Input:

```text

QUIT

```

Output:

```text

BYE

```

Requirements:

- The server should close only the current client connection.

- Other connected clients must remain connected.

---

## Error Handling

The server should return error messages for invalid input.

Example invalid commands:

```text

GET

SET onlykey

UNKNOWN command

DEL key extra

```

Example error response:

```text

ERR invalid command

```

or:

```text

ERR usage: SET <key> <value>

```

Requirements:

- Errors should be returned to the client.

- The server must continue running after bad input.

- A bad command from one client must not affect other clients.

- The program should avoid panics for normal user input.

---

## Testing Requirements

Use Go's built-in testing package.

Run all tests with:

```bash

go test ./...

```

### Store Tests

Test the in-memory store.

Required test cases:

- `Set` stores a value.

- `Get` returns an existing value.

- `Get` returns `false` for a missing key.

- `Set` overwrites an existing key.

- `Delete` removes an existing key.

- `Delete` returns `false` for a missing key.

- `Exists` returns the correct result.

- `Keys` returns stored keys.

### Parser Tests

Test command parsing.

Required test cases:

- Valid `PING`.

- Valid `SET`.

- Valid `GET`.

- Valid `DEL`.

- Valid `EXISTS`.

- Valid `KEYS`.

- Valid `QUIT`.

- Lowercase commands are normalized.

- Extra whitespace is handled.

- Empty input is handled.

- Invalid argument counts return errors.

### Concurrency Tests

Test concurrent access to the store.

Required test cases:

- Multiple goroutines can call `Set`.

- Multiple goroutines can call `Get`.

- Mixed reads and writes do not panic.

- The test should pass with the race detector.

Run with:

```bash

go test -race ./...

```

---

## Non-Functional Requirements

### Reliability

- The server should not crash from malformed client input.

- A disconnected client should not crash the server.

- Network errors should be handled gracefully.

### Simplicity

- Use the Go standard library first.

- Avoid unnecessary third-party dependencies.

- Keep packages small and focused.

### Readability

- Use clear package names.

- Use explicit error handling.

- Prefer simple code over clever code.

- Add comments only where they clarify non-obvious behavior.

### Concurrency Safety

- The shared store must be protected.

- The program should pass `go test -race ./...`.

- Avoid global mutable state unless it is protected.

---

## Milestones

### Milestone 1: Basic Store

Implement the store package.

Deliverables:

- `Store` type

- `Set`

- `Get`

- `Delete`

- `Exists`

- `Keys`

- Store unit tests

Completion criteria:

```bash

go test ./internal/store

```

passes.

---

### Milestone 2: Command Parser

Implement the protocol parser.

Deliverables:

- `Command` type

- Parser function

- Parser unit tests

Example function:

```go

func Parse(line string) (Command, error)

```

Completion criteria:

```bash

go test ./internal/protocol

```

passes.

---

### Milestone 3: Single-Client TCP Server

Implement a TCP server that handles one client at a time.

Deliverables:

- Server startup

- Basic command handling

- Client read/write loop

- Support for `PING`, `SET`, `GET`, and `QUIT`

Completion criteria:

A user can connect with `nc` and run:

```text

PING

SET language go

GET language

QUIT

```

---

### Milestone 4: Multi-Client Support

Update the server to handle multiple clients concurrently.

Deliverables:

- Goroutine per client connection

- Shared thread-safe store

- Clean connection shutdown

Completion criteria:

Two or more clients can connect at the same time and interact with the same store.

---

### Milestone 5: Full Core Command Set

Implement all required commands.

Deliverables:

- `PING`

- `SET`

- `GET`

- `DEL`

- `EXISTS`

- `KEYS`

- `QUIT`

Completion criteria:

All required commands work through a TCP client.

---

### Milestone 6: Race Detector Clean Run

Run the test suite with the race detector.

Completion criteria:

```bash

go test -race ./...

```

passes without data race warnings.

---

## Optional Extensions

### 1. TTL Support

Add expiring keys.

New commands:

```text

SETEX <key> <seconds> <value>

TTL <key>

```

Example:

```text

SETEX session 60 abc123

TTL session

```

Requirements:

- `SETEX` stores a key that expires after the given number of seconds.

- Expired keys should not be returned by `GET`.

- `TTL` returns the remaining time to live.

- Expired keys should eventually be removed from memory.

---

### 2. Persistence

Save the store to disk and load it on startup.

Example:

```bash

go run ./cmd/server --data ./data.json

```

Requirements:

- Store data in a JSON file.

- Load existing data on startup.

- Save data on shutdown.

- Optional: save data after every write.

---

### 3. Authentication

Require clients to authenticate before using store commands.

New command:

```text

AUTH <password>

```

Requirements:

- Password is configured with a command-line flag.

- Clients cannot use `GET`, `SET`, `DEL`, `EXISTS`, or `KEYS` before authenticating.

- `PING` and `QUIT` may remain available without authentication.

---

### 4. Pub/Sub

Add a basic publish-subscribe system.

New commands:

```text

SUBSCRIBE <channel>

PUBLISH <channel> <message>

```

Requirements:

- Clients can subscribe to a channel.

- Published messages are sent to subscribed clients.

- One client can publish while others receive messages.

- Use channels or a synchronized subscriber registry.

---

### 5. HTTP API

Expose the same store through HTTP.

Example endpoints:

```text

PUT    /keys/{key}

GET    /keys/{key}

DELETE /keys/{key}

GET    /keys

```

Requirements:

- TCP and HTTP APIs should use the same store.

- HTTP handlers should return proper status codes.

- Use `net/http`.

---

## Suggested Implementation Rules

- Do not use Redis libraries.

- Do not use a web framework for the core version.

- Use only the Go standard library unless an extension clearly needs otherwise.

- Keep the wire protocol simple.

- Write tests before or alongside each package.

- Run `gofmt` on all Go files.

- Run `go test ./...` frequently.

- Run `go test -race ./...` before considering the project complete.

---

## Final Completion Checklist

The project is complete when:

- [ ] The server starts from the command line.

- [ ] The port is configurable.

- [ ] Multiple clients can connect at once.

- [ ] The store is concurrency-safe.

- [ ] `PING` works.

- [ ] `SET` works.

- [ ] `GET` works.

- [ ] `DEL` works.

- [ ] `EXISTS` works.

- [ ] `KEYS` works.

- [ ] `QUIT` works.

- [ ] Invalid commands return errors.

- [ ] Malformed input does not crash the server.

- [ ] Store tests pass.

- [ ] Parser tests pass.

- [ ] Concurrency tests pass.

- [ ] `go test ./...` passes.

- [ ] `go test -race ./...` passes.

- [ ] The code is formatted with `gofmt`.

"""

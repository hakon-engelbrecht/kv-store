package command

import (
	"fmt"
	"strings"

	"github.com/hakon-engelbrecht/kv-store/internal/store"
)

type Command struct {
	Name string
	Args []string
}

// ParseCommand parses the provided message to a command.
// This function only produces valid commands.
// If the message contains no valid command, this function
// returns nil and the error.
func ParseCommand(msg string) (*Command, error) {
	parts := strings.Split(strings.TrimSpace(msg), " ")
	if len(parts) < 1 {
		return nil, fmt.Errorf("empty message")
	}

	cmd := strings.ToUpper(parts[0])

	switch cmd {
	case "PING":
		if len(parts) > 1 {
			return produceArgCountError("PING", len(parts))
		}
		return &Command{Name: "PING", Args: make([]string, 0)}, nil
	case "SET":
		if len(parts) != 3 {
			return produceArgCountError("SET", len(parts))
		}
		return &Command{Name: "SET", Args: parts[1:]}, nil
	case "GET":
		if len(parts) != 2 {
			return produceArgCountError("GET", len(parts))
		}
		return &Command{Name: "GET", Args: parts[1:]}, nil
	case "DEL":
		if len(parts) != 2 {
			return produceArgCountError("DEL", len(parts))
		}
		return &Command{Name: "DEL", Args: parts[1:]}, nil
	case "EXISTS":
		if len(parts) != 2 {
			return produceArgCountError("EXISTS", len(parts))
		}
		return &Command{Name: "EXISTS", Args: parts[1:]}, nil
	case "KEYS":
		if len(parts) != 1 {
			return produceArgCountError("KEYS", len(parts))
		}
		return &Command{Name: "KEYS", Args: make([]string, 0)}, nil
	case "QUIT":
		if len(parts) != 1 {
			return produceArgCountError("QUIT", len(parts))
		}
		return &Command{Name: "QUIT", Args: make([]string, 0)}, nil
	default:
		return nil, fmt.Errorf("unknown command: %v", msg)
	}
}

func (c *Command) Execute(s store.Store) string {
	switch c.Name {
	case "PING":
		return "PONG"
	case "SET":
		s.Set(c.Args[0], c.Args[1])
		return "OK"
	case "GET":
		val, exists := s.Get(c.Args[0])
		if exists {
			return val
		}
		return "(nil)"
	case "DEL":
		ok := s.Delete(c.Args[0])
		if ok {
			return "OK"
		}
		return "(nil)"
	case "EXISTS":
		exists := s.Exists(c.Args[0])
		if exists {
			return "true"
		}
		return "false"
	case "KEYS":
		keys := s.Keys()
		res := strings.Join(keys, "\n")
		return res
	case "QUIT":
		return "BYE"
	default:
		return ""
	}
}

func produceArgCountError(cmd string, argc int) (*Command, error) {
	return nil, fmt.Errorf("invalid argument count for %s command: %d", cmd, argc)
}

package loopia

import (
    "fmt"
    "log"

    "github.com/jonlil/loopia-go"
)

// Config -
type Config struct {
    Username   string
    Password   string
}

// Client -
func (c *Config) Client() (*loopia.API, error) {
    client, err := loopia.New(c.Username, c.Password)
    if err != nil {
    return nil, fmt.Errorf("Error creating new Loopia client: %s", err)
    }
    log.Printf("[INFO] Loopia Client configured for user: %s", c.Username)
    return client, nil
}

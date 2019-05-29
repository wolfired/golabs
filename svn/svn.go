package svn

import (
	"os/exec"
)

// Client Svn客户端
type Client struct {
	Username string
	Root     string
}

// Exec 执行svn操作
func (c *Client) Exec(args ...string) ([]byte, error) {
	return Exec(c.Root, args...)
}

// Exec 执行svn操作
func Exec(root string, args ...string) ([]byte, error) {
	cmd := exec.Command("svn", args...)
	cmd.Dir = root
	output, err := cmd.Output()
	return output, err
}

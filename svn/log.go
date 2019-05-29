package svn

import (
	"path/filepath"
	"strings"
)

// Log 日志
type Log struct {
	Entries []LogEntry `xml:"logentry"`
}

// LogEntry 日志条目
type LogEntry struct {
	Revision uint   `xml:"revision,attr"`
	Author   string `xml:"author"`
	Date     string `xml:"date"`
	Msg      string `xml:"msg"`
	Paths    []Path `xml:"paths>path"`
}

// Path 路径
type Path struct {
	Kind   string `xml:"kind,attr"`
	Action string `xml:"action,attr"`
	VALUE  string `xml:",chardata"`
}

// URL 相对路径
func (p Path) URL(relative string) string {
	return strings.Replace(p.VALUE, strings.Replace(relative, "^", "", -1)+"/", "", 1)
}

// URI 工作路径
func (p Path) URI(root string, relative string) string {
	return filepath.Join(root, p.URL(relative))
}

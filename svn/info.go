package svn

// Info 信息
type Info struct {
	Entry Entry `xml:"entry"`
}

// Entry 条目
type Entry struct {
	Path            string          `xml:"path,attr"`
	Revision        uint            `xml:"revision,attr"`
	Kind            string          `xml:"kind,attr"`
	URL             string          `xml:"url"`
	RelativeURL     string          `xml:"relative-url"`
	Repository      Repository      `xml:"repository"`
	WorkingCopyInfo WorkingCopyInfo `xml:"wc-info"`
	Commit          Commit          `xml:"commit"`
}

// Repository 仓库
type Repository struct {
	Root string `xml:"root"`
	UUID string `xml:"uuid"`
}

// WorkingCopyInfo 工作副本信息
type WorkingCopyInfo struct {
	WorkingCopyRootABSPath string `xml:"wcroot-abspath"`
	Schedule               string `xml:"schedule"`
	Depth                  string `xml:"depth"`
}

// Commit 提交
type Commit struct {
	Revision uint   `xml:"revision,attr"`
	Author   string `xml:"author"`
	Date     string `xml:"date"`
}

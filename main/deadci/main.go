package main

/*
本工具限定用于封神项目
*/
import (
	"bytes"
	"encoding/xml"
	"flag"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/wolfired/golabs/svn"
)

var (
	help         bool
	svnCDN       string
	localCDN     string
	svnCLI       string
	localCLI     string
	localDep     string
	swfDstRePath string
	username     string
	password     string
	waitSec      int64
)

var info *svn.Info
var svnCDNClient *svn.Client
var svnCLIClient *svn.Client

func main() {
	flag.BoolVar(&help, "help", false, "帮助")
	flag.StringVar(&svnCDN, "svn_cdn", "", "项目CDN svn 链接")
	flag.StringVar(&localCDN, "local_cdn", "", "项目CDN本地目录")
	flag.StringVar(&svnCLI, "svn_cli", "", "项目源码 svn 链接")
	flag.StringVar(&localCLI, "local_cli", "", "项目源码本地目录")
	flag.StringVar(&localDep, "local_dep", "", "本地私服目录, 任意指定一个目录就可以了")
	flag.StringVar(&swfDstRePath, "swf_dst_repath", "", "swf资源文件目标相对路径")
	flag.StringVar(&username, "username", "", "svn 用户")
	flag.StringVar(&password, "password", "", "svn 密码")
	flag.Int64Var(&waitSec, "wait_sec", 4, "轮询秒数")
	flag.Parse()

	if help || "" == svnCDN || "" == localCDN || "" == svnCLI || "" == localCLI || "" == localDep || "" == username {
		flag.Usage()
		return
	}

	svnCDNClient = &svn.Client{Username: username, Root: filepath.Join(localCDN)}
	svnCLIClient = &svn.Client{Username: username, Root: filepath.Join(localCLI)}

	if _, err := os.Stat(svnCDNClient.Root); os.IsNotExist(err) {
		svn.Exec(".", "co", "--depth=immediates", svnCDN, svnCDNClient.Root)

		svnCDNClient.Exec("up", "--depth=empty", filepath.Join("res/config"))
		svnCDNClient.Exec("up", "--set-depth=infinity", filepath.Join("res/config/name"))
		svnCDNClient.Exec("up", "--set-depth=infinity", filepath.Join("res/config/system"))
		svnCDNClient.Exec("up", "--set-depth=infinity", filepath.Join("res/ui"))
	}

	if _, err := os.Stat(svnCLIClient.Root); os.IsNotExist(err) {
		svn.Exec(".", "co", "--depth=immediates", svnCLI, svnCLIClient.Root)

		svnCLIClient.Exec("up", "--depth=empty", filepath.Join("bin-debug"))
		svnCLIClient.Exec("up", "--depth=empty", filepath.Join("bin-debug/res"))
		svnCLIClient.Exec("up", "--depth=empty", filepath.Join("bin-debug/res/ui"))

		svnCLIClient.Exec("up", "--set-depth=infinity", filepath.Join("src"))
		svnCLIClient.Exec("up", "--set-depth=infinity", filepath.Join("swc"))

		buf := bytes.NewBuffer([]byte{})

		err = template.Must(template.ParseFiles(filepath.Join(svnCLIClient.Root, "deadci_init.bat.temp"))).Execute(buf, struct {
			ProjectHome string
			ResHome     string
			DeployHome  string
		}{svnCLIClient.Root, svnCDNClient.Root, filepath.Join(localDep)})

		ioutil.WriteFile(filepath.Join(svnCLIClient.Root, "deadci_init.bat"), buf.Bytes(), os.ModePerm)

		copyFile(filepath.Join(svnCLIClient.Root, "deadci_svn_up_src.bat"), filepath.Join(svnCLIClient.Root, "deadci_svn_up_src.bat.temp"))
		copyFile(filepath.Join(svnCLIClient.Root, "deadci_bld_release.bat"), filepath.Join(svnCLIClient.Root, "deadci_bld_release.bat.temp"))
	}

	info = &svn.Info{}

	output, _ := svnCLIClient.Exec("info", "--xml")
	xml.Unmarshal(output, info)

	for {
		select {
		case <-time.Tick(time.Second * time.Duration(waitSec)):
			{
				log := &svn.Log{}
				output, _ := svnCLIClient.Exec("log", "--xml", "-v", "-r", "BASE:HEAD")
				xml.Unmarshal(output, log)

				if 1 >= len(log.Entries) {
					fmt.Println("wait")
					continue
				}

				svnCDNClient.Exec("up")
				svnCLIClient.Exec("up")

				need := make([]svn.LogEntry, 0, len(log.Entries)-1)

				for i := 1; i < len(log.Entries); i++ {
					e := log.Entries[i]
					if e.Author != username {
						continue
					}
					need = append(need, e)
				}

				copySwf(need)

				cmd := exec.Command("cmd.exe", "deadci_bld_release.bat")
				cmd.Dir = svnCLIClient.Root

				outt, _ := cmd.Output()

				fmt.Println(string(outt))

				// svnCDNClient.Exec("ci", "-m", mergeComments(need))
			}
		}
	}
}

func copySwf(entries []svn.LogEntry) {
	m := make(map[string]svn.Path)

	for _, e := range entries {
		for _, p := range e.Paths {
			value := p.VALUE
			if ".swf" != filepath.Ext(value) {
				continue
			}
			m[value] = p
		}
	}

	for _, p := range m {
		svnCLIClient.Exec("up", p.URL(info.Entry.RelativeURL))

		copyFile(p.URI(svnCLIClient.Root, info.Entry.RelativeURL), filepath.Join(svnCDNClient.Root, swfDstRePath, filepath.Base(p.URL(info.Entry.RelativeURL))))
	}
}

func mergeComments(entries []svn.LogEntry) string {
	mm := make(map[string]map[string]string)

	for _, e := range entries {
		msgs := strings.Split(e.Msg, "\n")
		m := mm[msgs[0]]

		if nil == m {
			mm[msgs[0]] = make(map[string]string)
			m = mm[msgs[0]]
		}

		for i := 1; i < len(msgs); i++ {
			m[msgs[i]] = msgs[i]
		}
	}

	strs := make([]string, 0, 32)

	for k, m := range mm {
		strs = append(strs, k)

		for _, v := range m {
			strs = append(strs, v)
		}
	}

	return strings.Join(strs, "\n")
}

func copyFile(dst string, src string) {
	sf, _ := os.Open(src)
	defer sf.Close()

	df, _ := os.OpenFile(dst, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	defer df.Close()

	io.Copy(df, sf)
}

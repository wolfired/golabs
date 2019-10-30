package main

/*
	pre_wait/a.png     pre_work/a.png
	cur_wait/a.png     cur_work/
	if md5(pre_wait/a.png) == md5(cur_wait/a.png) {
		copy pre_work/a.png to cur_work/a.png
	}
*/
import (
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var (
	help    bool
	exts    string
	preWait string
	preWork string
	curWait string
	curWork string
)

func main() {
	flag.BoolVar(&help, "help", false, "帮助")
	flag.StringVar(&exts, "exts", ".txt", "文本文件后缀")
	flag.StringVar(&preWait, "prewait", "", "前一版本待转换目录")
	flag.StringVar(&preWork, "prework", "", "前一版本已转换目录")
	flag.StringVar(&curWait, "curwait", "", "当前版本待转换目录")
	flag.StringVar(&curWork, "curwork", "", "当前版本已转换目录")
	flag.Parse()
	if help || "" == preWait || "" == preWork {
		flag.Usage()
		return
	}

	foreach(preWork, func(subFile string) bool {
		ext := filepath.Ext(subFile)

		if !strings.Contains(exts, ext) {
			return false
		}

		curByteWait, err := ioutil.ReadFile(strings.Replace(subFile, preWork, curWait, 1))

		if nil != err {
			// fmt.Println(err)
			return false
		}

		curHashWait := md5.Sum(curByteWait)
		curStrWait := hex.EncodeToString(curHashWait[:])

		preByteWait, _ := ioutil.ReadFile(strings.Replace(subFile, preWork, preWait, 1))
		preHashWait := md5.Sum(preByteWait)
		preStrWait := hex.EncodeToString(preHashWait[:])

		if curStrWait != preStrWait {
			return false
		}

		rbs, _ := ioutil.ReadFile(subFile)

		ioutil.WriteFile(strings.Replace(subFile, preWork, curWork, 1), rbs, os.ModePerm)

		return false
	})
}

func foreach(targetDir string, callback func(subFile string) bool) {
	infos, err := ioutil.ReadDir(targetDir)

	if nil != err {
		fmt.Println(err)
		return
	}

	for _, info := range infos {
		sub := filepath.Join(targetDir, info.Name())

		if info.IsDir() {
			foreach(sub, callback)
			continue
		}

		if callback(sub) {
			break
		}
	}
}

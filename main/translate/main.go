package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var (
	help   bool
	opencc string
	cccfg  string
	exts   string
	inDir  string
	outDir string
)

func main() {
	flag.BoolVar(&help, "help", false, "帮助")
	flag.StringVar(&opencc, "opencc", "", "opencc所在目录")
	flag.StringVar(&cccfg, "cccfg", "s2tw", "opencc -c 参数")
	flag.StringVar(&exts, "exts", ".txt", "文本文件后缀")
	flag.StringVar(&inDir, "indir", "", "输入目录")
	flag.StringVar(&outDir, "outdir", "", "输出目录, 如果不填, 将直接覆盖源文件")
	flag.Parse()

	if help || "" == inDir || "" == opencc {
		flag.Usage()
		return
	}

	if "" == outDir {
		outDir = inDir
	}

	outfile := func(infile string) string {
		if inDir != outDir {
			outfile := strings.Replace(infile, inDir, outDir, 1)
			os.MkdirAll(filepath.Dir(outfile), os.ModePerm)
			return outfile
		}

		return infile
	}

	foreach(inDir, func(subFile string) bool {
		ext := filepath.Ext(subFile)

		if !strings.Contains(exts, ext) {
			return false
		}

		cmdOpencc := exec.Command(opencc, "-i", subFile, "-o", outfile(subFile), "-c", cccfg)
		_, err := cmdOpencc.Output()

		if nil != err {
			fmt.Println(err)
		}

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

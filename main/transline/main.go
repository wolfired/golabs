package main

/*
if indir == outdir {
	input file name format: file_name.src.file_ext
	output file name format: file_name.dst.file_ext
} else {
	opencc -i indir/input_file -o outdir/input_file
}

file content:
md5("str")=escape("str")
*/
import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
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

	tempdir, err := ioutil.TempDir("", "transline")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(tempdir)

	foreach(inDir, func(subFile string) bool {
		ext := filepath.Ext(subFile)

		if !strings.Contains(exts, ext) {
			return false
		}

		bs, _ := ioutil.ReadFile(subFile)
		lines := strings.Split(string(bs), "\n")

		newLines := make([]string, len(lines))

		for i, line := range lines {
			coms := strings.Split(line, "=")
			oldLine, _ := url.PathUnescape(coms[1])

			tmpfn := filepath.Join(tempdir, "tmpfile")
			if err := ioutil.WriteFile(tmpfn, []byte(oldLine), os.ModePerm); err != nil {
				log.Fatal(err)
			}

			cmdOpencc := exec.Command(opencc, "-i", tmpfn, "-o", tmpfn, "-c", cccfg)
			_, err := cmdOpencc.Output()

			if nil != err {
				fmt.Println(err)
			}

			newLine, _ := ioutil.ReadFile(tmpfn)
			newLines[i] = coms[0] + "=" + url.PathEscape(string(newLine))
		}

		ioutil.WriteFile(outfile(subFile), []byte(strings.Join(newLines, "\n")), os.ModePerm)

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

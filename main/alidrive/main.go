package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/http/cookiejar"
	"os"
	"path/filepath"
	"regexp"
	"time"

	cid "github.com/ipfs/go-cid"
	mh "github.com/multiformats/go-multihash"
	qrcode "github.com/yeqown/go-qrcode"
)

var (
	help   bool
	debug  bool
	jar    *cookiejar.Jar
	client http.Client
)

func main() {
	flag.BoolVar(&help, "h", false, "help")
	flag.BoolVar(&debug, "d", false, "debug")
	p4hash := flag.String("p", "", "path to hash")
	n4hash := flag.String("n", ".hashed.json", "file name for hash output")
	R4hash := flag.Bool("R", false, "recursion directory")
	flag.Parse()

	if help {
		flag.Usage()
		return
	}

	jar, _ = cookiejar.New(nil)
	client = http.Client{
		Jar: jar,
	}

	switch flag.Arg(0) {
	case "hash":
		{
			hash(p4hash, n4hash, R4hash)
		}
	case "login":
		{
			login()
		}
	}
}

type HashedFolder struct {
	Folder string       `json:"folder"`
	Files  []HashedFile `json:"files"`
}

type HashedFile struct {
	Title       string `json:"title"`
	Size        int    `json:"size"`
	MD5         string `json:"md5"`
	SHA1        string `json:"sha1"`
	SHA256      string `json:"sha256"`
	IPFS_CID_V1 string `json:"ipfs_cid_v1"`
}

func hash(path *string, name *string, recursion *bool) {
	hashed_folder := HashedFolder{filepath.Base(*path), make([]HashedFile, 0)}

	fileinfos, _ := ioutil.ReadDir(*path)

	cb1 := cid.V1Builder{
		Codec:    cid.Raw,
		MhType:   mh.SHA2_256,
		MhLength: 0,
	}

	for _, fileinfo := range fileinfos {
		if fileinfo.IsDir() && *recursion {
			sub_path := filepath.Join(*path, fileinfo.Name())
			hash(&sub_path, name, recursion)
		} else if *name != fileinfo.Name() {
			bs, _ := ioutil.ReadFile(filepath.Join(*path, fileinfo.Name()))
			md5ba := md5.Sum(bs)
			sha1ba := sha1.Sum(bs)
			sha256ba := sha256.Sum256(bs)
			ipfs_cid_v1, _ := cb1.Sum(bs)
			hash_file := HashedFile{
				fileinfo.Name(),
				len(bs),
				hex.EncodeToString(md5ba[:]),
				hex.EncodeToString(sha1ba[:]),
				hex.EncodeToString(sha256ba[:]),
				ipfs_cid_v1.String(),
			}
			hashed_folder.Files = append(hashed_folder.Files, hash_file)
		}
	}

	bs, _ := json.Marshal(hashed_folder)
	ioutil.WriteFile(filepath.Join(*path, *name), bs, os.ModePerm)
}

func login() {
	url := ""
	client_id := ""
	redirect_uri := ""
	appName := ""
	appEntrance := ""
	_csrf_token := ""
	umidToken := ""
	returnUrl := ""
	hsiz := ""
	fromSite := ""
	bizParams := ""
	codeContent := ""

	{
		req, _ := http.NewRequest("GET", "https://www.aliyundrive.com/sign/in", nil)
		req.Header.Set("Host", "www.aliyundrive.com")
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:89.0) Gecko/20100101 Firefox/89.0")
		resp, _ := client.Do(req)
		bs, _ := ioutil.ReadAll(resp.Body)

		client_id = dumpStrFromBytes(bs, []string{`client_id\s*:\s*['"].+?['"]`, `['"].+?['"]`, `[^'"]+`}[:])
		redirect_uri = dumpStrFromBytes(bs, []string{`redirect_uri\s*:\s*['"].+?['"]`, `['"].+?['"]`, `[^'"]+`}[:])

		if debug {
			ioutil.WriteFile(fmt.Sprintf("./%d_%s", time.Now().UnixNano(), "sign_in.html"), bs, os.ModePerm)
		}
	}

	{
		req, _ := http.NewRequest("GET", "https://auth.aliyundrive.com/v2/oauth/authorize", nil)
		q := req.URL.Query()
		q.Set("client_id", client_id)
		q.Set("redirect_uri", redirect_uri)
		q.Set("response_type", "code")
		q.Set("login_type", "custom")
		q.Set("state", `{"origin":"https://www.aliyundrive.com"}`)
		req.URL.RawQuery = q.Encode()
		req.Header.Set("Host", "www.aliyundrive.com")
		req.Header.Set("Referer", "https://www.aliyundrive.com/")
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:89.0) Gecko/20100101 Firefox/89.0")

		resp, _ := client.Do(req)
		bs, _ := ioutil.ReadAll(resp.Body)

		if debug {
			ioutil.WriteFile(fmt.Sprintf("./%d_%s", time.Now().UnixNano(), "authorize.html"), bs, os.ModePerm)
		}
	}

	{
		req, _ := http.NewRequest("GET", "https://passport.aliyundrive.com/mini_login.htm", nil)
		q := req.URL.Query()
		q.Set("lang", "zh_cn")
		q.Set("appName", "aliyun_drive")
		q.Set("appEntrance", "web")
		q.Set("styleType", "auto")
		q.Set("bizParams", "")
		q.Set("notLoadSsoView", "false")
		q.Set("notKeepLogin", "false")
		q.Set("isMobile", "false")
		q.Set("hidePhoneCode", "true")
		q.Set("rnd", fmt.Sprintf("%f", rand.Float32()))
		req.URL.RawQuery = q.Encode()
		req.Header.Set("Host", "passport.aliyundrive.com")
		req.Header.Set("Referer", "https://auth.aliyundrive.com/")
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:89.0) Gecko/20100101 Firefox/89.0")

		url = req.URL.String()

		resp, _ := client.Do(req)
		bs, _ := ioutil.ReadAll(resp.Body)

		appName = dumpStrFromBytes(bs, []string{`['"]loginFormData['"]\s*:\s*\{.+?\}`, `['"]appName['"]\s*:\s*['"].+?['"]`, `\s*:\s*['"].+?['"]`, `['"].*?['"]`, `[^'"]+`}[:])
		appEntrance = dumpStrFromBytes(bs, []string{`['"]loginFormData['"]\s*:\s*\{.+?\}`, `['"]appEntrance['"]\s*:\s*['"].+?['"]`, `\s*:\s*['"].+?['"]`, `['"].*?['"]`, `[^'"]+`}[:])
		_csrf_token = dumpStrFromBytes(bs, []string{`['"]loginFormData['"]\s*:\s*\{.+?\}`, `['"]_csrf_token['"]\s*:\s*['"].+?['"]`, `\s*:\s*['"].+?['"]`, `['"].*?['"]`, `[^'"]+`}[:])
		umidToken = dumpStrFromBytes(bs, []string{`['"]loginFormData['"]\s*:\s*\{.+?\}`, `['"]umidToken['"]\s*:\s*['"].+?['"]`, `\s*:\s*['"].+?['"]`, `['"].*?['"]`, `[^'"]+`}[:])
		returnUrl = dumpStrFromBytes(bs, []string{`['"]loginFormData['"]\s*:\s*\{.+?\}`, `['"]returnUrl['"]\s*:\s*['"].+?['"]`, `\s*:\s*['"].+?['"]`, `['"].*?['"]`, `[^'"]+`}[:])
		hsiz = dumpStrFromBytes(bs, []string{`['"]loginFormData['"]\s*:\s*\{.+?\}`, `['"]hsiz['"]\s*:\s*['"].+?['"]`, `\s*:\s*['"].+?['"]`, `['"].*?['"]`, `[^'"]+`}[:])
		fromSite = dumpStrFromBytes(bs, []string{`['"]loginFormData['"]\s*:\s*\{.+?\}`, `['"]fromSite['"]\s*:\s*['"].+?['"]`, `\s*:\s*['"].+?['"]`, `['"].*?['"]`, `[^'"]+`}[:])
		bizParams = dumpStrFromBytes(bs, []string{`['"]loginFormData['"]\s*:\s*\{.+?\}`, `['"]bizParams['"]\s*:\s*['"].+?['"]`, `\s*:\s*['"].+?['"]`, `['"].*?['"]`, `[^'"]+`}[:])

		if debug {
			ioutil.WriteFile(fmt.Sprintf("./%d_%s", time.Now().UnixNano(), "mini_login.html"), bs, os.ModePerm)
		}
	}

	{
		req, _ := http.NewRequest("GET", "https://passport.aliyundrive.com/newlogin/qrcode/generate.do", nil)
		q := req.URL.Query()
		q.Set("appName", appName)
		// q.Add("appName", appName)
		q.Set("fromSite", fromSite)
		// q.Add("fromSite", fromSite)
		q.Set("appEntrance", appEntrance)
		q.Set("_csrf_token", _csrf_token)
		q.Set("umidToken", umidToken)
		q.Set("isMobile", "false")
		q.Set("lang", "zh_CN")
		q.Set("returnUrl", returnUrl)
		q.Set("hsiz", hsiz)
		q.Set("bizParams", bizParams)
		// q.Set("_bx-v", "2.0.31")
		req.URL.RawQuery = q.Encode()
		req.Header.Set("Host", "passport.aliyundrive.com")
		req.Header.Set("Referer", url)
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:89.0) Gecko/20100101 Firefox/89.0")

		resp, _ := client.Do(req)
		bs, _ := ioutil.ReadAll(resp.Body)

		codeContent = dumpStrFromBytes(bs, []string{`['"]codeContent['"]\s*:\s*['"].+?['"]`, `\s*:\s*['"].+?['"]`, `['"].*?['"]`, `[^'"]+`}[:])

		if debug {
			ioutil.WriteFile(fmt.Sprintf("./%d_%s", time.Now().UnixNano(), "generate.json"), bs, os.ModePerm)
		}
	}

	{
		qrc, _ := qrcode.New(codeContent)
		qrc.Save("./codeContent.jpeg")
	}

	{

	}
}

func dumpStrFromBytes(bs []byte, patterns []string) string {
	for _, pattern := range patterns {
		regexpc, _ := regexp.Compile(pattern)
		bs = regexpc.Find(bs)
	}
	return string(bs)
}

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"regexp"
	"strings"
)

var (
	help       bool
	jar        *cookiejar.Jar
	client     http.Client
	proxy      string
	username   string
	password   string
	licensealf string
	licenseulf string
	debug      bool
)

const usage = `Unity手动授权自动化, 用法:
  cd path/to/unity
  ./Unity.exe -batchmode -createManualActivationFile
  u3dmate -proxy http://127.0.0.1 -username XXX -password XXX -licensealf ./Unity_v2019.4.6f1.alf -licenseulf ./Unity_v2019.x.ulf
  ./Unity.exe -batchmode -manualLicenseFile ./Unity_v2019.x.ulf
`

func main() {
	flag.BoolVar(&help, "help", false, "Help")
	flag.StringVar(&proxy, "proxy", "", "Proxy, if you are in China, you have to use a proxy. Example: http://127.0.0.1:1080")
	flag.StringVar(&username, "username", "", "Your Unity Username")
	flag.StringVar(&password, "password", "", "Your Unity Password")
	flag.StringVar(&licensealf, "licensealf", "", "License alf file you want to upload")
	flag.StringVar(&licenseulf, "licenseulf", "", "Save path for the license ulf file")
	flag.BoolVar(&debug, "debug", false, "Debug mode for dev")

	flag.Parse()

	if help || "" == username || "" == password || "" == licensealf || "" == licenseulf {
		fmt.Println(usage)
		flag.Usage()
		return
	}

	jar, _ = cookiejar.New(nil)

	if "" == proxy {
		client = http.Client{
			Jar: jar,
		}
	} else {
		proxyURL, _ := url.Parse(proxy)
		client = http.Client{
			Jar:       jar,
			Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)},
		}
	}

	_, _ = openGet("https://license.unity3d.com/manual", "")
	_, _ = openGet("https://license.unity3d.com/genesis/api_domain", "")
	_, _ = openGet("https://license.unity3d.com/genesis/activation/step", "")
	bs, resp := openGet("https://license.unity3d.com/genesis/oauth/logout_callback?lastPage=/manual", "")

	authenticityToken := dumpFromBytes(dumpFromBytes(dumpFromBytes(dumpFromBytes(bs, "name=\"csrf-token\"\\s+content=\".+?\""), "content=\".+?\""), "\".+?\""), "[^\"]+")
	bs, resp = openLogin(resp.Request.URL.String(), string(authenticityToken), "")

	needCode := dumpFromBytes(bs, "Enter your code")

	if "Enter your code" == string(needCode) {
		code := ""
		fmt.Print("Enter your code: ")
		fmt.Scanf("%s", &code)
		bs, resp = openVerify(resp.Request.URL.String(), string(authenticityToken), code, "")
	}

	redirectURL := dumpFromBytes(dumpFromBytes(dumpFromBytes(bs, "<a\\s+href=\\s*['\"].+?['\"]"), "['\"].+?['\"]"), "[^'\"]+")
	_, _ = openGet(string(redirectURL), "")
	_, _ = openGet("https://license.unity3d.com/genesis/api_domain", "")
	_, _ = openGet("https://license.unity3d.com/genesis/activation/step", "")
	_ = openCreateTransaction("https://license.unity3d.com/genesis/activation/create-transaction", "")
	_, _ = openGet("https://license.unity3d.com/genesis/activation/step", "")
	_ = openUpdateTransaction("https://license.unity3d.com/genesis/activation/update-transaction", "")
	_, _ = openGet("https://license.unity3d.com/genesis/activation/step", "")
	_ = openDownloadLicense("https://license.unity3d.com/genesis/activation/download-license", "")
}

func dumpFromFile(filename string, pattern string) []byte {
	bs, _ := ioutil.ReadFile(filename)
	return dumpFromBytes(bs, pattern)
}

func dumpFromBytes(bs []byte, pattern string) []byte {
	regexpc, _ := regexp.Compile(pattern)
	return regexpc.Find(bs)
}

func openGet(url string, saved2 string) ([]byte, *http.Response) {
	fmt.Println(url)

	req, _ := http.NewRequest("GET", url, nil)
	resp, _ := client.Do(req)

	bs, _ := ioutil.ReadAll(resp.Body)

	if debug && "" != saved2 {
		ioutil.WriteFile(saved2, bs, os.ModePerm)
	}

	return bs, resp
}

func openToNormalLogin(url string, authenticityToken string, saved2 string) ([]byte, *http.Response) {
	fmt.Println(url)

	data := make(map[string][]string)
	data["_method"] = make([]string, 0)
	data["_method"] = append(data["_method"], "post")

	data["authenticity_token"] = make([]string, 0)
	data["authenticity_token"] = append(data["authenticity_token"], authenticityToken)

	resp, _ := client.PostForm(url, data)

	bs, _ := ioutil.ReadAll(resp.Body)

	if debug && "" != saved2 {
		ioutil.WriteFile(saved2, bs, os.ModePerm)
	}

	return bs, resp
}

func openPreToMailLoginOptions(url string, saved2 string) ([]byte, *http.Response) {
	fmt.Println(url)

	req, _ := http.NewRequest("OPTIONS", url, nil)
	resp, _ := client.Do(req)

	bs, _ := ioutil.ReadAll(resp.Body)

	if debug && "" != saved2 {
		ioutil.WriteFile(saved2, bs, os.ModePerm)
	}

	return bs, resp
}

func openPreToMailLogin(url string, saved2 string) ([]byte, *http.Response) {
	fmt.Println(url)

	req, _ := http.NewRequest("POST", url, strings.NewReader("{\"event\":\"toMailLogin\"}"))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := client.Do(req)

	bs, _ := ioutil.ReadAll(resp.Body)

	if debug && "" != saved2 {
		ioutil.WriteFile(saved2, bs, os.ModePerm)
	}

	return bs, resp
}

func openLogin(url string, authenticityToken string, saved2 string) ([]byte, *http.Response) {
	fmt.Println(url)

	data := make(map[string][]string)
	data["utf8"] = make([]string, 0)
	data["utf8"] = append(data["utf8"], "✓")

	data["_method"] = make([]string, 0)
	data["_method"] = append(data["_method"], "put")

	data["authenticity_token"] = make([]string, 0)
	data["authenticity_token"] = append(data["authenticity_token"], authenticityToken)

	data["conversations_create_session_form[email]"] = make([]string, 0)
	data["conversations_create_session_form[email]"] = append(data["conversations_create_session_form[email]"], username)

	data["conversations_create_session_form[password]"] = make([]string, 0)
	data["conversations_create_session_form[password]"] = append(data["conversations_create_session_form[password]"], password)

	data["conversations_create_session_form[remember_me]"] = make([]string, 0)
	data["conversations_create_session_form[remember_me]"] = append(data["conversations_create_session_form[remember_me]"], "false")

	// data["conversations_create_session_form[remember_me]"] = make([]string, 0)
	// data["conversations_create_session_form[remember_me]"] = append(data["conversations_create_session_form[remember_me]"], "true")

	data["commit"] = make([]string, 0)
	data["commit"] = append(data["commit"], "Sign in")

	resp, _ := client.PostForm(url, data)

	bs, _ := ioutil.ReadAll(resp.Body)

	if debug && "" != saved2 {
		ioutil.WriteFile(saved2, bs, os.ModePerm)
	}

	return bs, resp
}

func openVerify(url string, authenticityToken string, code string, saved2 string) ([]byte, *http.Response) {
	data := make(map[string][]string)
	data["utf8"] = make([]string, 0)
	data["utf8"] = append(data["utf8"], "✓")

	data["_method"] = make([]string, 0)
	data["_method"] = append(data["_method"], "put")

	data["authenticity_token"] = make([]string, 0)
	data["authenticity_token"] = append(data["authenticity_token"], authenticityToken)

	data["conversations_email_tfa_required_form[code]"] = make([]string, 0)
	data["conversations_email_tfa_required_form[code]"] = append(data["conversations_email_tfa_required_form[code]"], code)

	data["commit"] = make([]string, 0)
	data["commit"] = append(data["commit"], "Verify")

	resp, _ := client.PostForm(url, data)

	bs, _ := ioutil.ReadAll(resp.Body)

	if debug && "" != saved2 {
		ioutil.WriteFile(saved2, bs, os.ModePerm)
	}

	return bs, resp
}

func openCreateTransaction(url string, saved2 string) []byte {
	fmt.Println(url)

	fr, _ := os.Open(licensealf)
	req, _ := http.NewRequest("POST", url, fr)
	req.Header.Set("Content-Type", "text/xml")
	resp, _ := client.Do(req)

	bs, _ := ioutil.ReadAll(resp.Body)

	if debug && "" != saved2 {
		ioutil.WriteFile(saved2, bs, os.ModePerm)
	}

	return bs
}

func openUpdateTransaction(url string, saved2 string) []byte {
	fmt.Println(url)

	req, _ := http.NewRequest("PUT", url, strings.NewReader("{\"transaction\":{\"serial\":{\"type\":\"personal\"}}}"))
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	resp, _ := client.Do(req)

	bs, _ := ioutil.ReadAll(resp.Body)

	if debug && "" != saved2 {
		ioutil.WriteFile(saved2, bs, os.ModePerm)
	}

	return bs
}

func openDownloadLicense(url string, saved2 string) []byte {
	fmt.Println(url)

	req, _ := http.NewRequest("POST", url, strings.NewReader("{}"))
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	resp, _ := client.Do(req)

	bs, _ := ioutil.ReadAll(resp.Body)

	if debug && "" != saved2 {
		ioutil.WriteFile(saved2, bs, os.ModePerm)
	}

	payload := struct {
		Name string `json:"name"`
		XML  string `json:"xml"`
	}{}
	json.Unmarshal(bs, &payload)

	ioutil.WriteFile(licenseulf, []byte(payload.XML), os.ModePerm)

	return bs
}

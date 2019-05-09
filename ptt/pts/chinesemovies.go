package pts

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

func newChineseMovies(account *Account) Siter {
	return &chineseMovies{Account: account}
}

type chineseMovies struct {
	*Account
}

func (site *chineseMovies) Login() {
	var page string

	// cookiesMap := make(map[string][]*http.Cookie)
	// u, _ := url.Parse(site.Account.Host)

	//读取cookies
	// bs, _ := ioutil.ReadFile("./cookies.json")
	// json.Unmarshal(bs, &cookiesMap)
	// cookies := cookiesMap[site.Account.Site]
	// cj.SetCookies(u, cookies)

	// if ModeOnline == site.Mode || ModeOnlineAndSave == site.Mode {
	// 	page = "login.php"
	// 	httpc.Get(site.Host + "/" + page)

	// 	kvs := make(url.Values)
	// 	kvs.Add("logintype", "username")
	// 	kvs.Add("username", site.Name)
	// 	kvs.Add("password", site.Pass)
	// 	kvs.Add("thispagewidth", "yes")

	// 	page = "takelogin.php"
	// 	httpc.PostForm(site.Host+"/"+page, kvs)
	// }

	var resp *http.Response
	var doc *goquery.Document

	page = "index.php"
	if ModeOffline == site.Mode {
		f, _ := os.Open("./" + site.Site + "_" + page)
		defer f.Close()

		doc = Body2HTML(f)
	} else if ModeOnline == site.Mode || ModeOnlineAndSave == site.Mode {
		resp, _ = httpc.Get(site.Host + "/" + page)

		if nil == resp {
			return
		}

		defer resp.Body.Close()

		bs, _ := ioutil.ReadAll(resp.Body)

		if ModeOnlineAndSave == site.Mode {
			ioutil.WriteFile("./"+site.Site+"_"+page, []byte(Body2String(bytes.NewReader(bs))), os.ModePerm)
		}

		doc = Body2HTML(bytes.NewReader(bs))
	}

	//保存cookies
	// cookiesMap[site.Account.Site] = cj.Cookies(u)
	// bs, _ := json.Marshal(cookiesMap)
	// ioutil.WriteFile("./cookies.json", bs, os.ModePerm)

	site.ExtractUID(doc)
}

func (site *chineseMovies) Update() {
	if 0 == site.Info.UID {
		return
	}

	var page string
	var resp *http.Response
	var doc *goquery.Document

	page = "userdetails.php"

	if ModeOffline == site.Mode {
		f, _ := os.Open("./" + site.Site + "_" + page)
		defer f.Close()

		doc = Body2HTML(f)
	} else if ModeOnline == site.Mode || ModeOnlineAndSave == site.Mode {
		resp, _ = httpc.Get(site.Host + "/" + page + "?id=" + strconv.Itoa(site.Info.UID))

		if nil == resp {
			return
		}

		defer resp.Body.Close()

		bs, _ := ioutil.ReadAll(resp.Body)

		if ModeOnlineAndSave == site.Mode {
			ioutil.WriteFile("./"+site.Site+"_"+page, []byte(Body2String(bytes.NewReader(bs))), os.ModePerm)
		}

		doc = Body2HTML(bytes.NewReader(bs))
	}

	site.ExtractUpload(doc)
	site.ExtractDownload(doc)
	site.ExtractShare(doc)
	site.ExtractMagic(doc)
	site.ExtractSeed(doc)
	site.ExtractBulk(doc)
	site.ExtractDay(doc)
}

func (site *chineseMovies) Print() {
	fmt.Println(site.Account.Info)
}

func (site *chineseMovies) ExtractUID(doc *goquery.Document) { //UID
	doc.Find("table#info_block tr td table tr td span span a").Each(func(i int, s *goquery.Selection) {
		rg, _ := regexp.Compile(`\d+`)
		if href, ok := s.Attr("href"); ok {
			site.Info.UID, _ = strconv.Atoi(rg.FindString(href))
		}
	})
}
func (site *chineseMovies) ExtractUpload(doc *goquery.Document) { //上传量
	doc.Find("table.mainouter td#outer table.main table td").Each(func(i int, s *goquery.Selection) {
		if 15 == i {
			rg, _ := regexp.Compile(`[0-9,.]+[TGMK]B`)
			site.Info.Upload = rg.FindAllString(s.Text(), 1)[0]
		}
	})
}
func (site *chineseMovies) ExtractDownload(doc *goquery.Document) { //下载量
	doc.Find("table.mainouter td#outer table.main table td").Each(func(i int, s *goquery.Selection) {
		if 16 == i {
			rg, _ := regexp.Compile(`[0-9,.]+[TGMK]B`)
			site.Info.Download = rg.FindAllString(s.Text(), 1)[0]
		}
	})
}
func (site *chineseMovies) ExtractShare(doc *goquery.Document) { //分享率
	doc.Find("table.mainouter td#outer table.main table td").Each(func(i int, s *goquery.Selection) {
		if 14 == i {
			rg, _ := regexp.Compile(`[0-9,.]+`)
			rp, _ := regexp.Compile(`[,]+`)
			r := rp.ReplaceAllString(rg.FindAllString(s.Text(), 1)[0], "")
			site.Info.Share, _ = strconv.ParseFloat(r, 64)
		}
	})
}
func (site *chineseMovies) ExtractMagic(doc *goquery.Document) { //魔力值
	doc.Find("table.mainouter td#outer table.main table td").Each(func(i int, s *goquery.Selection) {
		if 34 == i {
			rg, _ := regexp.Compile(`[0-9,.]+`)
			rp, _ := regexp.Compile(`[,]+`)
			r := rp.ReplaceAllString(rg.FindAllString(s.Text(), 1)[0], "")
			site.Info.Magic, _ = strconv.ParseFloat(r, 64)
		}
	})
}
func (site *chineseMovies) ExtractSeed(doc *goquery.Document) { //做种数
	doc.Find("table.mainouter td#outer table.main table td").Each(func(i int, s *goquery.Selection) {
		if 38 == i {
			rg, _ := regexp.Compile(`^[0-9,]+`)
			rp, _ := regexp.Compile(`[,]+`)
			r := rp.ReplaceAllString(rg.FindAllString(s.Text(), 1)[0], "")
			site.Info.Seed, _ = strconv.Atoi(r)
		}
	})
}
func (site *chineseMovies) ExtractBulk(doc *goquery.Document) { //体积
	doc.Find("table.mainouter td#outer table.main table td").Each(func(i int, s *goquery.Selection) {
		if 38 == i {
			rg, _ := regexp.Compile(`:[0-9,.]+[TGMK]B\]`)
			rp, _ := regexp.Compile(`[:\]]+`)
			site.Info.Bulk = rp.ReplaceAllString(rg.FindAllString(s.Text(), 1)[0], "")
		}
	})
}
func (site *chineseMovies) ExtractDay(doc *goquery.Document) { //天数
	doc.Find("table.mainouter td#outer table.main table td").Each(func(i int, s *goquery.Selection) {
		if 21 == i {
			rg, _ := regexp.Compile(`\d+`)
			rp, _ := regexp.Compile(`\d{2}:\d{2}:\d{2}`)
			r := rp.ReplaceAllString(s.Text(), "")
			site.Info.Day, _ = strconv.Atoi(rg.FindAllString(r, 1)[0])
		}
	})
}
func (site *chineseMovies) ExtractWeek(doc *goquery.Document) { //入站

}
func (site *chineseMovies) ExtractLevel(doc *goquery.Document) { //等级

}

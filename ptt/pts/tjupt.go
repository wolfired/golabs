package pts

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

func newTjupt(account *Account) Siter {
	return &tjupt{Account: account}
}

type tjupt struct {
	*Account
}

func (site *tjupt) Login() {
	var page string

	if 0 == site.Mode || 1 == site.Mode {
		page = "login.php"
		httpc.Get(site.Host + "/" + page)

		kvs := make(url.Values)
		kvs.Add("username", site.Name)
		kvs.Add("password", site.Pass)
		kvs.Add("logout", "forever")

		page = "takelogin.php"
		httpc.PostForm(site.Host+"/"+page, kvs)
	}

	var resp *http.Response
	var doc *goquery.Document

	page = "index.php"
	if 2 == site.Mode {
		f, _ := os.Open("./" + site.Site + "_" + page)
		defer f.Close()

		doc = Body2HTML(f)
	} else if 0 == site.Mode || 1 == site.Mode {
		resp, _ = httpc.Get(site.Host + "/" + page)

		if nil == resp {
			return
		}

		defer resp.Body.Close()

		bs, _ := ioutil.ReadAll(resp.Body)

		if 1 == site.Mode {
			ioutil.WriteFile("./"+site.Site+"_"+page, []byte(Body2String(bytes.NewReader(bs))), os.ModePerm)
		}

		doc = Body2HTML(bytes.NewReader(bs))
	}

	site.ExtractUID(doc)
}

func (site *tjupt) Update() {
	if 0 == site.Info.UID {
		return
	}

	var page string
	var resp *http.Response
	var doc *goquery.Document

	page = "userdetails.php"

	if 2 == site.Mode {
		f, _ := os.Open("./" + site.Site + "_" + page)
		defer f.Close()

		doc = Body2HTML(f)
	} else if 0 == site.Mode || 1 == site.Mode {
		resp, _ = httpc.Get(site.Host + "/" + page + "?id=" + strconv.Itoa(site.Info.UID))

		if nil == resp {
			return
		}

		defer resp.Body.Close()

		bs, _ := ioutil.ReadAll(resp.Body)

		if 1 == site.Mode {
			ioutil.WriteFile("./"+site.Site+"_"+page, []byte(Body2String(bytes.NewReader(bs))), os.ModePerm)
		}

		doc = Body2HTML(bytes.NewReader(bs))
	}

	if 0 == site.Mode || 2 == site.Mode {
		site.ExtractUpload(doc)
		site.ExtractDownload(doc)
		site.ExtractShare(doc)
		site.ExtractMagic(doc)
		site.ExtractSeed(doc)
		site.ExtractBulk(doc)
		site.ExtractDay(doc)
	}
}

func (site *tjupt) Print() {
	fmt.Println(site.Account.Info)
}

func (site *tjupt) ExtractUID(doc *goquery.Document) { //UID
	doc.Find("table#info_block tr td table tr td span span a").Each(func(i int, s *goquery.Selection) {
		rg, _ := regexp.Compile(`\d+`)
		if href, ok := s.Attr("href"); ok {
			site.Info.UID, _ = strconv.Atoi(rg.FindString(href))
		}
	})
}
func (site *tjupt) ExtractUpload(doc *goquery.Document) { //上传量
	doc.Find("table.mainouter td#outer table.main table td").Each(func(i int, s *goquery.Selection) {
		if 32 == i {
			rg, _ := regexp.Compile(`[0-9,.]+\s*[TGMK]B`)
			site.Info.Upload = rg.FindAllString(s.Text(), 1)[0]
		}
	})
}
func (site *tjupt) ExtractDownload(doc *goquery.Document) { //下载量
	doc.Find("table.mainouter td#outer table.main table td").Each(func(i int, s *goquery.Selection) {
		if 33 == i {
			rg, _ := regexp.Compile(`[0-9,.]+\s*[TGMK]B`)
			site.Info.Download = rg.FindAllString(s.Text(), 1)[0]
		}
	})
}
func (site *tjupt) ExtractShare(doc *goquery.Document) { //分享率
	doc.Find("table.mainouter td#outer table.main table td").Each(func(i int, s *goquery.Selection) {
		if 30 == i {
			rg, _ := regexp.Compile(`[0-9,.]+`)
			rp, _ := regexp.Compile(`[,]+`)
			r := rp.ReplaceAllString(rg.FindAllString(s.Text(), 1)[0], "")
			site.Info.Share, _ = strconv.ParseFloat(r, 64)
		}
	})
}
func (site *tjupt) ExtractMagic(doc *goquery.Document) { //魔力值
	doc.Find("table.mainouter td#outer table.main table td").Each(func(i int, s *goquery.Selection) {
		if 56 == i {
			rg, _ := regexp.Compile(`[0-9,.]+`)
			rp, _ := regexp.Compile(`[,]+`)
			r := rp.ReplaceAllString(rg.FindAllString(s.Text(), 1)[0], "")
			site.Info.Magic, _ = strconv.ParseFloat(r, 64)
		}
	})
}
func (site *tjupt) ExtractSeed(doc *goquery.Document) { //做种数
	// doc.Find("table.mainouter td#outer table.main table td").Each(func(i int, s *goquery.Selection) {
	// 	if 38 == i {
	// 		rg, _ := regexp.Compile(`^[0-9,]+`)
	// 		rp, _ := regexp.Compile(`[,]+`)
	// 		r := rp.ReplaceAllString(rg.FindAllString(s.Text(), 1)[0], "")
	// 		site.Info.Seed, _ = strconv.Atoi(r)
	// 	}
	// })
}
func (site *tjupt) ExtractBulk(doc *goquery.Document) { //体积
	// doc.Find("table.mainouter td#outer table.main table td").Each(func(i int, s *goquery.Selection) {
	// 	if 38 == i {
	// 		rg, _ := regexp.Compile(`:[0-9,.]+[TGMK]B\]`)
	// 		rp, _ := regexp.Compile(`[:\]]+`)
	// 		site.Info.Bulk = rp.ReplaceAllString(rg.FindAllString(s.Text(), 1)[0], "")
	// 	}
	// })
}
func (site *tjupt) ExtractDay(doc *goquery.Document) { //天数
	doc.Find("table.mainouter td#outer table.main table td").Each(func(i int, s *goquery.Selection) {
		if 43 == i {
			rg, _ := regexp.Compile(`\d+`)
			rp, _ := regexp.Compile(`\d{2}:\d{2}:\d{2}`)
			r := rp.ReplaceAllString(s.Text(), "")
			site.Info.Day, _ = strconv.Atoi(rg.FindAllString(r, 1)[0])
		}
	})
}
func (site *tjupt) ExtractWeek(doc *goquery.Document) { //入站

}
func (site *tjupt) ExtractLevel(doc *goquery.Document) { //等级

}

package pts

import (
	"bytes"
	"io"
	"net/http"
	"net/http/cookiejar"

	"github.com/PuerkitoBio/goquery"
)

const (
	//ModeIgnore 跳过
	ModeIgnore int = 1
	//ModeOnline 连网不保存
	ModeOnline int = 2
	//ModeOnlineAndSave 连网并保存
	ModeOnlineAndSave int = 3
	//ModeOffline 离线
	ModeOffline int = 4
)

var (
	cj    *cookiejar.Jar
	httpc *http.Client
	newfn map[string]func(account *Account) Siter
)

func init() {
	cj, _ = cookiejar.New(nil)
	httpc = &http.Client{Jar: cj}

	newfn = make(map[string]func(account *Account) Siter)

	newfn["chinesemovies"] = newChineseMovies
	newfn["hyperay"] = newHyperay
	newfn["tjupt"] = newTjupt
}

/*
Siter 网站
*/
type Siter interface {
	Login()
	Update()
	Print()

	ExtractUID(doc *goquery.Document)      //UID
	ExtractUpload(doc *goquery.Document)   //上传量
	ExtractDownload(doc *goquery.Document) //下载量
	ExtractShare(doc *goquery.Document)    //分享率
	ExtractMagic(doc *goquery.Document)    //魔力值
	ExtractSeed(doc *goquery.Document)     //做种数
	ExtractBulk(doc *goquery.Document)     //体积
	ExtractDay(doc *goquery.Document)      //天数
	ExtractWeek(doc *goquery.Document)     //入站
	ExtractLevel(doc *goquery.Document)    //等级
}

/*
JSONPtt 全部账号
*/
type JSONPtt struct {
	Accounts []*Account
}

/*
Account 账号
*/
type Account struct {
	Mode int
	Site string
	Host string
	Name string
	Pass string
	Info Info
}

/*
Info 信息
*/
type Info struct {
	UID      int     //UID
	Upload   string  //上传量
	Download string  //下载量
	Share    float64 //分享率
	Magic    float64 //魔力值
	Seed     int     //做种数
	Bulk     string  //体积
	Day      int     //天数
	Week     int     //入站
	Level    int     //等级
}

/*
Body2String 文本
*/
func Body2String(r io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r)
	return buf.String()
}

/*
Body2HTML 页面
*/
func Body2HTML(r io.Reader) *goquery.Document {
	doc, _ := goquery.NewDocumentFromReader(r)
	return doc
}

/*
New 新建
*/
func New(account *Account) Siter {
	return newfn[account.Site](account)
}

package ptt

import (
	"encoding/json"
	"io/ioutil"

	"github.com/wolfired/golabs/ptt/pts"
)

/*
Boot 启动
*/
func Boot() {
	parse()

	bs, _ := ioutil.ReadFile(flags.json)

	jp := &pts.JSONPtt{}
	json.Unmarshal(bs, jp)

	for _, p := range jp.Accounts {
		if pts.ModeIgnore == p.Mode {
			continue
		}
		site := pts.New(p)
		site.Login()
		site.Update()
		site.Print()
	}

}

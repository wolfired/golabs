package main

import (
	"fmt"
	"github.com/wolfired/golabs/bitorent"
	"io/ioutil"

)


func main() {
	bs, _ := ioutil.ReadFile("./bg.torrent")
	d := bitorent.Decode(bs)

	for k, v := range d {
		fmt.Println(k, ":", v)
	}

	// info := d["info"].(map[string]interface{})
	// fmt.Println(hex.EncodeToString([]byte(info["pieces"].(string))))

	// ioutil.WriteFile("./tmp.torrent", bitorent.Encode(d), 0777)

	m := bitorent.MetaInfo{}
	m.Parse(d)
	fmt.Println(m)
}

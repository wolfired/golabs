package main

import (
	"fmt"
	"github.com/wolfired/golabs/idiotDB"
	"github.com/wolfired/golabs/swfchef"
	"io/ioutil"
	"os"
)

func main() {
	testSwfchef()
}

func testIdiotDB() {
	// zz := idiotDB.CreateZipWrapper()
	// zz.AddZipItem("name.txt", []byte("LinkWu"))
	// zz.Close()

	// zw := idiotDB.CreateZipWrapper()
	// zw.AddZipItem("name.zip", zz.RawData())

	// zw.AddZipItems(map[string]string{"age": "12", "phone": "18601011241"})
	// zw.Close()

	// ioutil.WriteFile("C:\\Users\\zelda\\Desktop\\res\\zip_item.zip", zw.RawData(), os.ModePerm)

	table := idiotDB.MetaTable{}
	table.Name = "string"
	table.Fields = map[string]idiotDB.MetaField{}
	table.Fields["value"] = idiotDB.MetaField{"value", "string", ""}
	ioutil.WriteFile("C:\\Users\\zelda\\Desktop\\res\\"+table.Name+"\\.zip", table.RawData(), os.ModePerm)

	table = idiotDB.MetaTable{}
	table.Name = "item"
	table.Fields = map[string]idiotDB.MetaField{}
	table.Fields["name"] = idiotDB.MetaField{"name", "string", ""}
	table.Fields["price_id"] = idiotDB.MetaField{"price_id", "uint", "0"}
	table.Fields["price_count"] = idiotDB.MetaField{"price_count", "uint", "0"}
	ioutil.WriteFile("C:\\Users\\zelda\\Desktop\\res\\"+table.Name+"\\.zip", table.RawData(), os.ModePerm)
}

func testSwfchef() {
	s := swfchef.SwfRead("C:\\Users\\zelda\\Desktop\\res\\C.swf")

	fmt.Printf("%s\n", s.Signature())
	fmt.Printf("%d\n", s.Version())
	fmt.Printf("%d\n", s.Length())
	fmt.Println(s.FrameSize())
	fmt.Printf("%f\n", s.FrameRate())
	fmt.Println(s.FrameCount())

	// fmt.Println(swfchef.Raw2sbn([]byte{0x60, 0x00, 0x28, 0x00, 0x00, 0x28, 0x00}, 41, 12))
	// fmt.Println(swfchef.Raw2sbn([]byte{0xFF, 0xFE, 0xFD, 0xFC}, 8, 8))
}

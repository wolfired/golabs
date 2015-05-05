package main

import (
	"github.com/wolfired/golabs/idiotDB"
	"io/ioutil"
	"os"
)

func main() {
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

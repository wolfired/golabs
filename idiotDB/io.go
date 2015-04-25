package idiotDB

import (
	"archive/zip"
	"bytes"
	"log"
	"os"
)

func create_zip() {
	buffer := new(bytes.Buffer)

	ptr_zip_writer := zip.NewWriter(buffer)

	for k, v := range name_content {
		add_zip_item(ptr_zip_writer, k, []byte(v))
	}

	ptr_zip_writer.Close()
}

func create_zip_item(ptr_zip_writer *zip.Writer, item_name string, item_raw_data []byte) {
	item_writer, err := ptr_zip_writer.Create(item_name)

	if nil != err {
		log.Fatalln(err)
	}
	item_writer.Write(item_raw_data)
}

func save_zip_file(path string, buffer *bytes.Buffer) {
	ptr_zip_file, err := os.Create(path)
	defer ptr_zip_file.Close()

	if nil != err {
		log.Fatalln(err)
	}
	ptr_zip_file.Write(buffer.Bytes())
}

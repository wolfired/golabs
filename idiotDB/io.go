package idiotDB

import (
	"archive/zip"
	"bytes"
	"log"
	"os"
)

func create_zip(path string, name_content map[string]string) {
	buffer := bytes.Buffer{}

	ptr_zip_writer := zip.NewWriter(&buffer)

	for k, v := range name_content {
		add_zip_item(ptr_zip_writer, k, []byte(v))
	}

	ptr_zip_writer.Close()

	save_zip_file(path, &buffer)
}

func add_zip_item(ptr_zip_writer *zip.Writer, item_name string, raw_bytes []byte) {
	item_writer, err := ptr_zip_writer.Create(item_name)

	if nil != err {
		log.Fatalln(err)
	}
	item_writer.Write(raw_bytes)
}

func save_zip_file(path string, buffer *bytes.Buffer) {
	ptr_zip_file, err := os.Create(path)
	defer ptr_zip_file.Close()

	if nil != err {
		log.Fatalln(err)
	}
	ptr_zip_file.Write(buffer.Bytes())
}

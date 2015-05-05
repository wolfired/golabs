package idiotDB

import (
	"archive/zip"
	"bytes"
)

type ZipWrapper struct {
	zipBuffer *bytes.Buffer
	zipWriter *zip.Writer
}

func CreateZipWrapper() *ZipWrapper {
	z := new(ZipWrapper)
	z.zipBuffer = new(bytes.Buffer)
	z.zipWriter = zip.NewWriter(z.zipBuffer)
	return z
}

func (z *ZipWrapper) AddZipItem(itemKey string, itemData []byte) {
	zipItemWriter, _ := z.zipWriter.Create(itemKey)
	zipItemWriter.Write(itemData)
}

func (z *ZipWrapper) AddZipItems(itemKeyData map[string]string) {
	for itemKey, itemData := range itemKeyData {
		z.AddZipItem(itemKey, []byte(itemData))
	}
}

func (z *ZipWrapper) Size() int {
	return z.zipBuffer.Len()
}

func (z *ZipWrapper) Close() {
	z.zipWriter.Close()
}

func (z *ZipWrapper) RawData() []byte {
	return z.zipBuffer.Bytes()
}

package amfchef

type marker byte

const (
	undefinedMarker    marker = 0x00
	nullMarker         marker = 0x01
	falseMarker        marker = 0x02
	trueMarker         marker = 0x03
	integerMarker      marker = 0x04
	doubleMarker       marker = 0x05
	stringMarker       marker = 0x06
	xmlDocMarker       marker = 0x07
	dateMarker         marker = 0x08
	arrayMarker        marker = 0x09
	objectMarker       marker = 0x0a
	xmlMarker          marker = 0x0b
	byteArrayMarker    marker = 0x0c
	vectorIntMarker    marker = 0x0d
	vectorUintMarker   marker = 0x0e
	vectorDoubleMarker marker = 0x0f
	vectorObjectMarker marker = 0x10
	dictionaryMarker   marker = 0x11
)

package swfchef

import (
	"reflect"
)

var (
	tagMap [128]reflect.Type = [128]reflect.Type{}
)

func init() {
}

type tagHeader struct {
}

func (th *tagHeader) analyse() {

}

type tagEnd struct {
	tagHeader
}
type tagShowFrame struct {
	tagHeader
}
type tagDefineShape struct {
	tagHeader
}
type tagPlaceObject struct {
	tagHeader
}
type tagRemoveObject struct {
	tagHeader
}
type tagDefineBits struct {
	tagHeader
}
type tagDefineButton struct {
	tagHeader
}
type tagJPEGTables struct {
	tagHeader
}
type tagSetBackgroundColor struct {
	tagHeader
}
type tagDefineFont struct {
	tagHeader
}
type tagDefineText struct {
	tagHeader
}
type tagDoAction struct {
	tagHeader
}
type tagDefineFontInfo struct {
	tagHeader
}
type tagDefineSound struct {
	tagHeader
}
type tagStartSound struct {
	tagHeader
}
type tagDefineButtonSound struct {
	tagHeader
}
type tagSoundStreamHead struct {
	tagHeader
}
type tagSoundStreamBlock struct {
	tagHeader
}
type tagDefineBitsLossless struct {
	tagHeader
}
type tagDefineBitsJPEG2 struct {
	tagHeader
}
type tagDefineShape2 struct {
	tagHeader
}
type tagDefineButtonCxform struct {
	tagHeader
}
type tagProtect struct {
	tagHeader
}
type tagPlaceObject2 struct {
	tagHeader
}
type tagRemoveObject2 struct {
	tagHeader
}
type tagDefineShape3 struct {
	tagHeader
}
type tagDefineText2 struct {
	tagHeader
}
type tagDefineButton2 struct {
	tagHeader
}
type tagDefineBitsJPEG3 struct {
	tagHeader
}
type tagDefineBitsLossless2 struct {
	tagHeader
}
type tagDefineEditText struct {
	tagHeader
}
type tagDefineSprite struct {
	tagHeader
}
type tagFrameLabel struct {
	tagHeader
}
type tagSoundStreamHead2 struct {
	tagHeader
}
type tagDefineMorphShape struct {
	tagHeader
}
type tagDefineFont2 struct {
	tagHeader
}
type tagExportAssets struct {
	tagHeader
}
type tagImportAssets struct {
	tagHeader
}
type tagEnableDebugger struct {
	tagHeader
}
type tagDoInitAction struct {
	tagHeader
}
type tagDefineVideoStream struct {
	tagHeader
}
type tagVideoFrame struct {
	tagHeader
}
type tagDefineFontInfo2 struct {
	tagHeader
}
type tagEnableDebugger2 struct {
	tagHeader
}
type tagScriptLimits struct {
	tagHeader
}
type tagSetTabIndex struct {
	tagHeader
}
type tagFileAttributes struct {
	tagHeader
}
type tagPlaceObject3 struct {
	tagHeader
}
type tagImportAssets2 struct {
	tagHeader
}
type tagDefineFontAlignZones struct {
	tagHeader
}
type tagCSMTextSettings struct {
	tagHeader
}
type tagDefineFont3 struct {
	tagHeader
}
type tagSymbolClass struct {
	tagHeader
}
type tagMetadata struct {
	tagHeader
}
type tagDefineScalingGrid struct {
	tagHeader
}
type tagDoABC struct {
	tagHeader
}
type tagDefineShape4 struct {
	tagHeader
}
type tagDefineMorphShape2 struct {
	tagHeader
}
type tagDefineSceneAndFrameLabelData struct {
	tagHeader
}
type tagDefineBinaryData struct {
	tagHeader
}
type tagDefineFontName struct {
	tagHeader
}
type tagStartSound2 struct {
	tagHeader
}
type tagDefineBitsJPEG4 struct {
	tagHeader
}
type tagDefineFont4 struct {
	tagHeader
}
type tagEnableTelemetry struct {
	tagHeader
}

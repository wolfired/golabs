package swfchef

const (
	tagIndexEnd                          uint = 0
	tagIndexShowFrame                    uint = 1
	tagIndexDefineShape                  uint = 2
	tagIndexPlaceObject                  uint = 4
	tagIndexRemoveObject                 uint = 5
	tagIndexDefineBits                   uint = 6
	tagIndexDefineButton                 uint = 7
	tagIndexJPEGTables                   uint = 8
	tagIndexSetBackgroundColor           uint = 9
	tagIndexDefineFont                   uint = 10
	tagIndexDefineText                   uint = 11
	tagIndexDoAction                     uint = 12
	tagIndexDefineFontInfo               uint = 13
	tagIndexDefineSound                  uint = 14
	tagIndexStartSound                   uint = 15
	tagIndexDefineButtonSound            uint = 17
	tagIndexSoundStreamHead              uint = 18
	tagIndexSoundStreamBlock             uint = 19
	tagIndexDefineBitsLossless           uint = 20
	tagIndexDefineBitsJPEG2              uint = 21
	tagIndexDefineShape2                 uint = 22
	tagIndexDefineButtonCxform           uint = 23
	tagIndexProtect                      uint = 24
	tagIndexPlaceObject2                 uint = 26
	tagIndexRemoveObject2                uint = 28
	tagIndexDefineShape3                 uint = 32
	tagIndexDefineText2                  uint = 33
	tagIndexDefineButton2                uint = 34
	tagIndexDefineBitsJPEG3              uint = 35
	tagIndexDefineBitsLossless2          uint = 36
	tagIndexDefineEditText               uint = 37
	tagIndexDefineSprite                 uint = 39
	tagIndexFrameLabel                   uint = 43
	tagIndexSoundStreamHead2             uint = 45
	tagIndexDefineMorphShape             uint = 46
	tagIndexDefineFont2                  uint = 48
	tagIndexExportAssets                 uint = 56
	tagIndexImportAssets                 uint = 57
	tagIndexEnableDebugger               uint = 58
	tagIndexDoInitAction                 uint = 59
	tagIndexDefineVideoStream            uint = 60
	tagIndexVideoFrame                   uint = 61
	tagIndexDefineFontInfo2              uint = 62
	tagIndexEnableDebugger2              uint = 64
	tagIndexScriptLimits                 uint = 65
	tagIndexSetTabIndex                  uint = 66
	tagIndexFileAttributes               uint = 69
	tagIndexPlaceObject3                 uint = 70
	tagIndexImportAssets2                uint = 71
	tagIndexDefineFontAlignZones         uint = 73
	tagIndexCSMTextSettings              uint = 74
	tagIndexDefineFont3                  uint = 75
	tagIndexSymbolClass                  uint = 76
	tagIndexMetadata                     uint = 77
	tagIndexDefineScalingGrid            uint = 78
	tagIndexDoABC                        uint = 82
	tagIndexDefineShape4                 uint = 83
	tagIndexDefineMorphShape2            uint = 84
	tagIndexDefineSceneAndFrameLabelData uint = 86
	tagIndexDefineBinaryData             uint = 87
	tagIndexDefineFontName               uint = 88
	tagIndexStartSound2                  uint = 89
	tagIndexDefineBitsJPEG4              uint = 90
	tagIndexDefineFont4                  uint = 91
	tagIndexEnableTelemetry              uint = 93

	tagLabelEnd                          string = "End"
	tagLabelShowFrame                    string = "ShowFrame"
	tagLabelDefineShape                  string = "DefineShape"
	tagLabelPlaceObject                  string = "PlaceObject"
	tagLabelRemoveObject                 string = "RemoveObject"
	tagLabelDefineBits                   string = "DefineBits"
	tagLabelDefineButton                 string = "DefineButton"
	tagLabelJPEGTables                   string = "JPEGTables"
	tagLabelSetBackgroundColor           string = "SetBackgroundColor"
	tagLabelDefineFont                   string = "DefineFont"
	tagLabelDefineText                   string = "DefineText"
	tagLabelDoAction                     string = "DoAction"
	tagLabelDefineFontInfo               string = "DefineFontInfo"
	tagLabelDefineSound                  string = "DefineSound"
	tagLabelStartSound                   string = "StartSound"
	tagLabelDefineButtonSound            string = "DefineButtonSound"
	tagLabelSoundStreamHead              string = "SoundStreamHead"
	tagLabelSoundStreamBlock             string = "SoundStreamBlock"
	tagLabelDefineBitsLossless           string = "DefineBitsLossless"
	tagLabelDefineBitsJPEG2              string = "DefineBitsJPEG2"
	tagLabelDefineShape2                 string = "DefineShape2"
	tagLabelDefineButtonCxform           string = "DefineButtonCxform"
	tagLabelProtect                      string = "Protect"
	tagLabelPlaceObject2                 string = "PlaceObject2"
	tagLabelRemoveObject2                string = "RemoveObject2"
	tagLabelDefineShape3                 string = "DefineShape3"
	tagLabelDefineText2                  string = "DefineText2"
	tagLabelDefineButton2                string = "DefineButton2"
	tagLabelDefineBitsJPEG3              string = "DefineBitsJPEG3"
	tagLabelDefineBitsLossless2          string = "DefineBitsLossless2"
	tagLabelDefineEditText               string = "DefineEditText"
	tagLabelDefineSprite                 string = "DefineSprite"
	tagLabelFrameLabel                   string = "FrameLabel"
	tagLabelSoundStreamHead2             string = "SoundStreamHead2"
	tagLabelDefineMorphShape             string = "DefineMorphShape"
	tagLabelDefineFont2                  string = "DefineFont2"
	tagLabelExportAssets                 string = "ExportAssets"
	tagLabelImportAssets                 string = "ImportAssets"
	tagLabelEnableDebugger               string = "EnableDebugger"
	tagLabelDoInitAction                 string = "DoInitAction"
	tagLabelDefineVideoStream            string = "DefineVideoStream"
	tagLabelVideoFrame                   string = "VideoFrame"
	tagLabelDefineFontInfo2              string = "DefineFontInfo2"
	tagLabelEnableDebugger2              string = "EnableDebugger2"
	tagLabelScriptLimits                 string = "ScriptLimits"
	tagLabelSetTabIndex                  string = "SetTabIndex"
	tagLabelFileAttributes               string = "FileAttributes"
	tagLabelPlaceObject3                 string = "PlaceObject3"
	tagLabelImportAssets2                string = "ImportAssets2"
	tagLabelDefineFontAlignZones         string = "DefineFontAlignZones"
	tagLabelCSMTextSettings              string = "CSMTextSettings"
	tagLabelDefineFont3                  string = "DefineFont3"
	tagLabelSymbolClass                  string = "SymbolClass"
	tagLabelMetadata                     string = "Metadata"
	tagLabelDefineScalingGrid            string = "DefineScalingGrid"
	tagLabelDoABC                        string = "DoABC"
	tagLabelDefineShape4                 string = "DefineShape4"
	tagLabelDefineMorphShape2            string = "DefineMorphShape2"
	tagLabelDefineSceneAndFrameLabelData string = "DefineSceneAndFrameLabelData"
	tagLabelDefineBinaryData             string = "DefineBinaryData"
	tagLabelDefineFontName               string = "DefineFontName"
	tagLabelStartSound2                  string = "StartSound2"
	tagLabelDefineBitsJPEG4              string = "DefineBitsJPEG4"
	tagLabelDefineFont4                  string = "DefineFont4"
	tagLabelEnableTelemetry              string = "EnableTelemetry"
)

type tagHandler func([]byte) interface{}

var (
	tagLabelMap   [128]string     = [128]string{}
	tagHandlerMap [128]tagHandler = [128]tagHandler{}
)

func init() {
	tagLabelMap[tagIndexEnd] = tagLabelEnd
	tagLabelMap[tagIndexShowFrame] = tagLabelShowFrame
	tagLabelMap[tagIndexDefineShape] = tagLabelDefineShape
	tagLabelMap[tagIndexPlaceObject] = tagLabelPlaceObject
	tagLabelMap[tagIndexRemoveObject] = tagLabelRemoveObject
	tagLabelMap[tagIndexDefineBits] = tagLabelDefineBits
	tagLabelMap[tagIndexDefineButton] = tagLabelDefineButton
	tagLabelMap[tagIndexJPEGTables] = tagLabelJPEGTables
	tagLabelMap[tagIndexSetBackgroundColor] = tagLabelSetBackgroundColor
	tagLabelMap[tagIndexDefineFont] = tagLabelDefineFont
	tagLabelMap[tagIndexDefineText] = tagLabelDefineText
	tagLabelMap[tagIndexDoAction] = tagLabelDoAction
	tagLabelMap[tagIndexDefineFontInfo] = tagLabelDefineFontInfo
	tagLabelMap[tagIndexDefineSound] = tagLabelDefineSound
	tagLabelMap[tagIndexStartSound] = tagLabelStartSound
	tagLabelMap[tagIndexDefineButtonSound] = tagLabelDefineButtonSound
	tagLabelMap[tagIndexSoundStreamHead] = tagLabelSoundStreamHead
	tagLabelMap[tagIndexSoundStreamBlock] = tagLabelSoundStreamBlock
	tagLabelMap[tagIndexDefineBitsLossless] = tagLabelDefineBitsLossless
	tagLabelMap[tagIndexDefineBitsJPEG2] = tagLabelDefineBitsJPEG2
	tagLabelMap[tagIndexDefineShape2] = tagLabelDefineShape2
	tagLabelMap[tagIndexDefineButtonCxform] = tagLabelDefineButtonCxform
	tagLabelMap[tagIndexProtect] = tagLabelProtect
	tagLabelMap[tagIndexPlaceObject2] = tagLabelPlaceObject2
	tagLabelMap[tagIndexRemoveObject2] = tagLabelRemoveObject2
	tagLabelMap[tagIndexDefineShape3] = tagLabelDefineShape3
	tagLabelMap[tagIndexDefineText2] = tagLabelDefineText2
	tagLabelMap[tagIndexDefineButton2] = tagLabelDefineButton2
	tagLabelMap[tagIndexDefineBitsJPEG3] = tagLabelDefineBitsJPEG3
	tagLabelMap[tagIndexDefineBitsLossless2] = tagLabelDefineBitsLossless2
	tagLabelMap[tagIndexDefineEditText] = tagLabelDefineEditText
	tagLabelMap[tagIndexDefineSprite] = tagLabelDefineSprite
	tagLabelMap[tagIndexFrameLabel] = tagLabelFrameLabel
	tagLabelMap[tagIndexSoundStreamHead2] = tagLabelSoundStreamHead2
	tagLabelMap[tagIndexDefineMorphShape] = tagLabelDefineMorphShape
	tagLabelMap[tagIndexDefineFont2] = tagLabelDefineFont2
	tagLabelMap[tagIndexExportAssets] = tagLabelExportAssets
	tagLabelMap[tagIndexImportAssets] = tagLabelImportAssets
	tagLabelMap[tagIndexEnableDebugger] = tagLabelEnableDebugger
	tagLabelMap[tagIndexDoInitAction] = tagLabelDoInitAction
	tagLabelMap[tagIndexDefineVideoStream] = tagLabelDefineVideoStream
	tagLabelMap[tagIndexVideoFrame] = tagLabelVideoFrame
	tagLabelMap[tagIndexDefineFontInfo2] = tagLabelDefineFontInfo2
	tagLabelMap[tagIndexEnableDebugger2] = tagLabelEnableDebugger2
	tagLabelMap[tagIndexScriptLimits] = tagLabelScriptLimits
	tagLabelMap[tagIndexSetTabIndex] = tagLabelSetTabIndex
	tagLabelMap[tagIndexFileAttributes] = tagLabelFileAttributes
	tagLabelMap[tagIndexPlaceObject3] = tagLabelPlaceObject3
	tagLabelMap[tagIndexImportAssets2] = tagLabelImportAssets2
	tagLabelMap[tagIndexDefineFontAlignZones] = tagLabelDefineFontAlignZones
	tagLabelMap[tagIndexCSMTextSettings] = tagLabelCSMTextSettings
	tagLabelMap[tagIndexDefineFont3] = tagLabelDefineFont3
	tagLabelMap[tagIndexSymbolClass] = tagLabelSymbolClass
	tagLabelMap[tagIndexMetadata] = tagLabelMetadata
	tagLabelMap[tagIndexDefineScalingGrid] = tagLabelDefineScalingGrid
	tagLabelMap[tagIndexDoABC] = tagLabelDoABC
	tagLabelMap[tagIndexDefineShape4] = tagLabelDefineShape4
	tagLabelMap[tagIndexDefineMorphShape2] = tagLabelDefineMorphShape2
	tagLabelMap[tagIndexDefineSceneAndFrameLabelData] = tagLabelDefineSceneAndFrameLabelData
	tagLabelMap[tagIndexDefineBinaryData] = tagLabelDefineBinaryData
	tagLabelMap[tagIndexDefineFontName] = tagLabelDefineFontName
	tagLabelMap[tagIndexStartSound2] = tagLabelStartSound2
	tagLabelMap[tagIndexDefineBitsJPEG4] = tagLabelDefineBitsJPEG4
	tagLabelMap[tagIndexDefineFont4] = tagLabelDefineFont4
	tagLabelMap[tagIndexEnableTelemetry] = tagLabelEnableTelemetry

	tagHandlerMap[tagIndexEnd] = tagHandlerEnd
	tagHandlerMap[tagIndexShowFrame] = tagHandlerShowFrame
	tagHandlerMap[tagIndexDefineShape] = tagHandlerDefineShape
	tagHandlerMap[tagIndexPlaceObject] = tagHandlerPlaceObject
	tagHandlerMap[tagIndexRemoveObject] = tagHandlerRemoveObject
	tagHandlerMap[tagIndexDefineBits] = tagHandlerDefineBits
	tagHandlerMap[tagIndexDefineButton] = tagHandlerDefineButton
	tagHandlerMap[tagIndexJPEGTables] = tagHandlerJPEGTables
	tagHandlerMap[tagIndexSetBackgroundColor] = tagHandlerSetBackgroundColor
	tagHandlerMap[tagIndexDefineFont] = tagHandlerDefineFont
	tagHandlerMap[tagIndexDefineText] = tagHandlerDefineText
	tagHandlerMap[tagIndexDoAction] = tagHandlerDoAction
	tagHandlerMap[tagIndexDefineFontInfo] = tagHandlerDefineFontInfo
	tagHandlerMap[tagIndexDefineSound] = tagHandlerDefineSound
	tagHandlerMap[tagIndexStartSound] = tagHandlerStartSound
	tagHandlerMap[tagIndexDefineButtonSound] = tagHandlerDefineButtonSound
	tagHandlerMap[tagIndexSoundStreamHead] = tagHandlerSoundStreamHead
	tagHandlerMap[tagIndexSoundStreamBlock] = tagHandlerSoundStreamBlock
	tagHandlerMap[tagIndexDefineBitsLossless] = tagHandlerDefineBitsLossless
	tagHandlerMap[tagIndexDefineBitsJPEG2] = tagHandlerDefineBitsJPEG2
	tagHandlerMap[tagIndexDefineShape2] = tagHandlerDefineShape2
	tagHandlerMap[tagIndexDefineButtonCxform] = tagHandlerDefineButtonCxform
	tagHandlerMap[tagIndexProtect] = tagHandlerProtect
	tagHandlerMap[tagIndexPlaceObject2] = tagHandlerPlaceObject2
	tagHandlerMap[tagIndexRemoveObject2] = tagHandlerRemoveObject2
	tagHandlerMap[tagIndexDefineShape3] = tagHandlerDefineShape3
	tagHandlerMap[tagIndexDefineText2] = tagHandlerDefineText2
	tagHandlerMap[tagIndexDefineButton2] = tagHandlerDefineButton2
	tagHandlerMap[tagIndexDefineBitsJPEG3] = tagHandlerDefineBitsJPEG3
	tagHandlerMap[tagIndexDefineBitsLossless2] = tagHandlerDefineBitsLossless2
	tagHandlerMap[tagIndexDefineEditText] = tagHandlerDefineEditText
	tagHandlerMap[tagIndexDefineSprite] = tagHandlerDefineSprite
	tagHandlerMap[tagIndexFrameLabel] = tagHandlerFrameLabel
	tagHandlerMap[tagIndexSoundStreamHead2] = tagHandlerSoundStreamHead2
	tagHandlerMap[tagIndexDefineMorphShape] = tagHandlerDefineMorphShape
	tagHandlerMap[tagIndexDefineFont2] = tagHandlerDefineFont2
	tagHandlerMap[tagIndexExportAssets] = tagHandlerExportAssets
	tagHandlerMap[tagIndexImportAssets] = tagHandlerImportAssets
	tagHandlerMap[tagIndexEnableDebugger] = tagHandlerEnableDebugger
	tagHandlerMap[tagIndexDoInitAction] = tagHandlerDoInitAction
	tagHandlerMap[tagIndexDefineVideoStream] = tagHandlerDefineVideoStream
	tagHandlerMap[tagIndexVideoFrame] = tagHandlerVideoFrame
	tagHandlerMap[tagIndexDefineFontInfo2] = tagHandlerDefineFontInfo2
	tagHandlerMap[tagIndexEnableDebugger2] = tagHandlerEnableDebugger2
	tagHandlerMap[tagIndexScriptLimits] = tagHandlerScriptLimits
	tagHandlerMap[tagIndexSetTabIndex] = tagHandlerSetTabIndex
	tagHandlerMap[tagIndexFileAttributes] = tagHandlerFileAttributes
	tagHandlerMap[tagIndexPlaceObject3] = tagHandlerPlaceObject3
	tagHandlerMap[tagIndexImportAssets2] = tagHandlerImportAssets2
	tagHandlerMap[tagIndexDefineFontAlignZones] = tagHandlerDefineFontAlignZones
	tagHandlerMap[tagIndexCSMTextSettings] = tagHandlerCSMTextSettings
	tagHandlerMap[tagIndexDefineFont3] = tagHandlerDefineFont3
	tagHandlerMap[tagIndexSymbolClass] = tagHandlerSymbolClass
	tagHandlerMap[tagIndexMetadata] = tagHandlerMetadata
	tagHandlerMap[tagIndexDefineScalingGrid] = tagHandlerDefineScalingGrid
	tagHandlerMap[tagIndexDoABC] = tagHandlerDoABC
	tagHandlerMap[tagIndexDefineShape4] = tagHandlerDefineShape4
	tagHandlerMap[tagIndexDefineMorphShape2] = tagHandlerDefineMorphShape2
	tagHandlerMap[tagIndexDefineSceneAndFrameLabelData] = tagHandlerDefineSceneAndFrameLabelData
	tagHandlerMap[tagIndexDefineBinaryData] = tagHandlerDefineBinaryData
	tagHandlerMap[tagIndexDefineFontName] = tagHandlerDefineFontName
	tagHandlerMap[tagIndexStartSound2] = tagHandlerStartSound2
	tagHandlerMap[tagIndexDefineBitsJPEG4] = tagHandlerDefineBitsJPEG4
	tagHandlerMap[tagIndexDefineFont4] = tagHandlerDefineFont4
	tagHandlerMap[tagIndexEnableTelemetry] = tagHandlerEnableTelemetry

}

func tagHandlerEnd(raw []byte) interface{} {
	return tag{}
}
func tagHandlerShowFrame(raw []byte) interface{} {
	return tag{}
}
func tagHandlerDefineShape(raw []byte) interface{} {
	return tag{}
}
func tagHandlerPlaceObject(raw []byte) interface{} {
	return tag{}
}
func tagHandlerRemoveObject(raw []byte) interface{} {
	return tag{}
}
func tagHandlerDefineBits(raw []byte) interface{} {
	return tag{}
}
func tagHandlerDefineButton(raw []byte) interface{} {
	return tag{}
}
func tagHandlerJPEGTables(raw []byte) interface{} {
	return tag{}
}
func tagHandlerSetBackgroundColor(raw []byte) interface{} {
	return tag{}
}
func tagHandlerDefineFont(raw []byte) interface{} {
	return tag{}
}
func tagHandlerDefineText(raw []byte) interface{} {
	return tag{}
}
func tagHandlerDoAction(raw []byte) interface{} {
	return tag{}
}
func tagHandlerDefineFontInfo(raw []byte) interface{} {
	return tag{}
}
func tagHandlerDefineSound(raw []byte) interface{} {
	return tag{}
}
func tagHandlerStartSound(raw []byte) interface{} {
	return tag{}
}
func tagHandlerDefineButtonSound(raw []byte) interface{} {
	return tag{}
}
func tagHandlerSoundStreamHead(raw []byte) interface{} {
	return tag{}
}
func tagHandlerSoundStreamBlock(raw []byte) interface{} {
	return tag{}
}
func tagHandlerDefineBitsLossless(raw []byte) interface{} {
	return tag{}
}
func tagHandlerDefineBitsJPEG2(raw []byte) interface{} {
	return tag{}
}
func tagHandlerDefineShape2(raw []byte) interface{} {
	return tag{}
}
func tagHandlerDefineButtonCxform(raw []byte) interface{} {
	return tag{}
}
func tagHandlerProtect(raw []byte) interface{} {
	return tag{}
}
func tagHandlerPlaceObject2(raw []byte) interface{} {
	return tag{}
}
func tagHandlerRemoveObject2(raw []byte) interface{} {
	return tag{}
}
func tagHandlerDefineShape3(raw []byte) interface{} {
	return tag{}
}
func tagHandlerDefineText2(raw []byte) interface{} {
	return tag{}
}
func tagHandlerDefineButton2(raw []byte) interface{} {
	return tag{}
}
func tagHandlerDefineBitsJPEG3(raw []byte) interface{} {
	return tag{}
}
func tagHandlerDefineBitsLossless2(raw []byte) interface{} {
	return tag{}
}
func tagHandlerDefineEditText(raw []byte) interface{} {
	return tag{}
}
func tagHandlerDefineSprite(raw []byte) interface{} {
	return tag{}
}
func tagHandlerFrameLabel(raw []byte) interface{} {
	return tag{}
}
func tagHandlerSoundStreamHead2(raw []byte) interface{} {
	return tag{}
}
func tagHandlerDefineMorphShape(raw []byte) interface{} {
	return tag{}
}
func tagHandlerDefineFont2(raw []byte) interface{} {
	return tag{}
}
func tagHandlerExportAssets(raw []byte) interface{} {
	return tag{}
}
func tagHandlerImportAssets(raw []byte) interface{} {
	return tag{}
}
func tagHandlerEnableDebugger(raw []byte) interface{} {
	return tag{}
}
func tagHandlerDoInitAction(raw []byte) interface{} {
	return tag{}
}
func tagHandlerDefineVideoStream(raw []byte) interface{} {
	return tag{}
}
func tagHandlerVideoFrame(raw []byte) interface{} {
	return tag{}
}
func tagHandlerDefineFontInfo2(raw []byte) interface{} {
	return tag{}
}
func tagHandlerEnableDebugger2(raw []byte) interface{} {
	return tag{}
}
func tagHandlerScriptLimits(raw []byte) interface{} {
	return tag{}
}
func tagHandlerSetTabIndex(raw []byte) interface{} {
	return tag{}
}
func tagHandlerFileAttributes(raw []byte) interface{} {
	return tag{}
}
func tagHandlerPlaceObject3(raw []byte) interface{} {
	return tag{}
}
func tagHandlerImportAssets2(raw []byte) interface{} {
	return tag{}
}
func tagHandlerDefineFontAlignZones(raw []byte) interface{} {
	return tag{}
}
func tagHandlerCSMTextSettings(raw []byte) interface{} {
	return tag{}
}
func tagHandlerDefineFont3(raw []byte) interface{} {
	return tag{}
}
func tagHandlerSymbolClass(raw []byte) interface{} {
	return tag{}
}
func tagHandlerMetadata(raw []byte) interface{} {
	return tag{}
}
func tagHandlerDefineScalingGrid(raw []byte) interface{} {
	return tag{}
}
func tagHandlerDoABC(raw []byte) interface{} {
	return tag{}
}
func tagHandlerDefineShape4(raw []byte) interface{} {
	return tag{}
}
func tagHandlerDefineMorphShape2(raw []byte) interface{} {
	return tag{}
}
func tagHandlerDefineSceneAndFrameLabelData(raw []byte) interface{} {
	return tag{}
}
func tagHandlerDefineBinaryData(raw []byte) interface{} {
	return tag{}
}
func tagHandlerDefineFontName(raw []byte) interface{} {
	return tag{}
}
func tagHandlerStartSound2(raw []byte) interface{} {
	return tag{}
}
func tagHandlerDefineBitsJPEG4(raw []byte) interface{} {
	return tag{}
}
func tagHandlerDefineFont4(raw []byte) interface{} {
	return tag{}
}
func tagHandlerEnableTelemetry(raw []byte) interface{} {
	return tag{}
}

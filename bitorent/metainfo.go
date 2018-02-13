package bitorent

const (
	tagName string = "torrent"
)

/*MetaInfo 种子文件*/
type MetaInfo struct {
	Announce string   `torrent:"announce"`
	Info     InfoDict `torrent:"info"`
}

/*Parse 解析*/
func (m *MetaInfo)Parse(d map[string]interface{})  {
	m.Announce = d["announce"].(string)
	m.Info.Parse(d["info"].(map[string]interface{}))
}

/*InfoDict 信息字典*/
type InfoDict struct {
	Name        string `torrent:"name"`         //单文件: 文件名，多文件: 目录名
	PieceLength int `torrent:"piece length"` //
	Pieces      string `torrent:"pieces"`

	//单文件
	Length int `torrent:"length"`

	//多文件
	Files []FileDict `torrent:"files"`
}

/*Parse 解析*/
func (i *InfoDict)Parse(d map[string]interface{})  {
	i.Name = d["name"].(string)
	i.PieceLength = d["piece length"].(int)

	i.Length = d["length"].(int)
}

/*FileDict 文件字典*/
type FileDict struct {
	Length uint64   `torrent:"length"`
	Path   []string `torrent:"path"`
}

/*Parse 解析*/
func (f *FileDict)Parse(d map[string]interface{})  {

}

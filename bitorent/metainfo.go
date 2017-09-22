package bitorent

const (
	tagName string = "torrent"
)

/*MetaInfo 种子文件*/
type MetaInfo struct {
	Announce string   `torrent:"announce"`
	Info     InfoDict `torrent:"info"`
}

/*InfoDict 信息字典*/
type InfoDict struct {
	Name        string   `torrent:"name"`         //单文件: 文件名，多文件: 目录名
	PieceLength uint64   `torrent:"piece length"` //
	Pieces      []string `torrent:"pieces"`

	//单文件
	Length uint64 `torrent:"length"`

	//多文件
	Files []FileDict `torrent:"files"`
}

/*FileDict 文件字典*/
type FileDict struct {
	Length uint64   `torrent:"length"`
	Path   []string `torrent:"path"`
}

package server

import ()

type IProtocol interface {
	Serialize([]byte)
	Deserialize([]byte)
}

type Protocol struct {
	Id uint
}

func (p *Protocol) Serialize(raw []byte) {

}

func (p *Protocol) Deserialize(raw []byte) {

}

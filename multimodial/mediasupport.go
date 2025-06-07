package multimodial

import (
	_ "embed"
	"encoding/json"
	"log"
)

type mediaSupportModel struct {
	Media struct {
		Image []string `json:"image"`
		Audio []string `json:"audio"`
		Video []string `json:"video"`
	} `json:"media"`
}

//go:embed media_support.json
var byteMediaSupport []byte

var MediaSuport *mediaSupportModel

func init() {
	MediaSuport = new(mediaSupportModel)
	err := json.Unmarshal(byteMediaSupport, &MediaSuport)
	if err != nil {
		log.Fatal("error read file media_support.json")
	}
}

package geminimodels

import (
	"fmt"
	"log"

	"github.com/hend41234/gemini/geminiutils"
)

// you can choose the mode. the zero-shoot, streaming or multi-turn
//
//	"zero-shoot" = (you can empty params) use the "FL")
//
// response saved in geminimodels.Response
//
//	"stream" = use "FLS"
//
// if you use stream mode you can use 1 more value, this "context".
//
// when you want "contex" of chat saved alias history chat still use to next chat, so chat no out of contex.
//
//	you can input `s` = 'save', or 'n' = 'no save' .
//
// thats mean, if you save the history of context, you use 'flash-lite stream' + 'flash-lite multi-turn'
//
//	"multi-turn" = use "FLSM"
//
//	'streaming mode still does not support media upload'
func (reqConfig BaseRequestModel) Gemini2FL(mode ...string) bool {
	// byteConf, _ := json.Marshal(reqConfig)
	var m string
	if len(mode) > 0 {
		if mode[0] == "FLS" {
			save := true
			m = "2" + mode[0]
			if len(mode) > 1 {
				if mode[1] == "n" {
					save = false
				} else if mode[1] != "s" {
					log.Fatal("args mode params not valid value")
				}
				// m = "2FLS"
			}
			url := fmt.Sprintf("%v/%v%v",
				geminiutils.Utils.BaseURL,
				geminiutils.Utils.Endpoint[m],
				geminiutils.Utils.GeminiApiKey,
			)
			sendingStream(url, reqConfig, save)
		}
	} else {
		// fmt.Println("2FL")
		m = "2FL" // default model
		reqConfig.Contents[0].Role = nil
		url := fmt.Sprintf("%v/%v%v",
			geminiutils.Utils.BaseURL,
			geminiutils.Utils.Endpoint[m],
			geminiutils.Utils.GeminiApiKey,
		)
		send := sendRequest(url, reqConfig)
		return send
	}
	return false
}

// func (reqConfig BaseRequestModel) Gemini2FLStream() {
// byteConf, _ := json.Marshal(reqConfig)
// }

package geminimodels

import (
	"fmt"

	"github.com/hend41234/gemini/geminiutils"
)

func (reqConfig BaseRequestModel) Gemini2p5FL() {
	reqConfig.Contents[0].Role = nil
	// byteConf, _ := json.Marshal(reqConfig)

	url := fmt.Sprintf("%v/%v%v",
		geminiutils.Utils.BaseURL,
		geminiutils.Utils.Endpoint["2F"],
		geminiutils.Utils.GeminiApiKey,
	)

	send := sendRequest(url, reqConfig)
	fmt.Println(send)
}

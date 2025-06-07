package main

import (
	"github.com/hend41234/gemini/geminimodels"
)

func main() {
	geminimodels.QuickGenerateConfigRequest("bro, whats up!")
	geminimodels.ConfigRequest.Gemini2FL("FLS", "n")

	// fmt.Println(resModel.Candidates[0].Contents.Parts[0].Text)
	// fmt.Println(resModel.GetUsageMetadata())

}

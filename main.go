package main

import "github.com/hend41234/gemini/geminimodels"

func main() {
	// sample Streaming Multimodial
	geminimodels.QuickGenerateConfigRequest("bro, whats up!")
	geminimodels.ConfigRequest.Gemini2FL("FLS", "s")

}

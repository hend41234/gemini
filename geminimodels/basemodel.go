package geminimodels

type InlineData struct {
	MimeType string `json:"mime_type"`
	// directly use data from os.ReadFile(), and then encode to Base64
	Data string `json:"data"`
}
type Part struct {
	Text string `json:"text"`
	// (optional) using when input multimodial, so you can join input text, and media file.
	InlineDatas *InlineData `json:"inline_data,omitempty"`
}

type Content struct {
	Role  *string `json:"role,omitempty"`
	Parts []Part  `json:"parts"`
}

type BaseRequestModel struct {
	Contents []Content `json:"contents"`
	// (optional) its advance config, and some models maybe not support in some params
	//
	// make sure model that used is support with params that you used
	//
	// see more : https://ai.google.dev/api/generate-content?hl=id#v1beta.GenerationConfig
	GenerationConfig *GenerateContentConfig `json:"generationConfig,omitempty"`
}

type ResModels struct {
	Candidates []struct {
		Contents       Content `json:"content"`
		FinishedReason string  `json:"finishedReason"`
		AvgLogprobs    float64 `json:"avgLogprobs"`
	} `json:"candidates"`
	UsageMetadata struct {
		PromptTokenCount       int        `json:"promptTokenCount"`
		CandidatesTokenCount   int        `json:"candidatesTokenCount"`
		TotalTokenCount        int        `json:"totalTokenCount"`
		PromptTokensDetails    []Modality `json:"promtTokensDetails"`
		CandidatesTokenDetails []Modality `json:"candidatesTokenDetails"`
	} `json:"usageMetadata"`
	ModelVersion string `json:"modelVersion"`
	ResponseID   string `json:"responseId"`
}

type TokenDetails struct {
	Modality   string `json:"modality"`
	TokenCount int    `json:"tokenCount"`
}

type ResStreamingModels struct {
	Data ResModels `json:"data"`
}

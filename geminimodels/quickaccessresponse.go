package geminimodels

import "encoding/json"

func (res ResModels) GetContentText() string {
	return res.Candidates[0].Contents.Parts[0].Text
}

func (res ResModels) GetModelVersionUsed() string {
	return res.ModelVersion
}

func (res ResModels) GetAvgLogprobs() float64 {
	return res.Candidates[0].AvgLogprobs
}

func (res ResModels) GetUsageMetadata() string {
	metadata, _ := json.Marshal(res.UsageMetadata)
	return string(metadata)
}

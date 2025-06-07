package geminimodels

type GenerateContentConfig struct {
	Temperature                *float64         `json:"temperature,omitempty"`
	TopP                       *float64         `json:"topP,omitempty"`
	TopK                       *int             `json:"topK,omitempty"`
	MaxOutputTokens            *int             `json:"maxOutputTokens,omitempty"`
	StopSequences              []string         `json:"stopSequences,omitempty"`
	ResponseMimeType           *string          `json:"responseMimeType,omitempty"`
	ResponseSchema             *Schema          `json:"responseSchema,omitempty"`
	ResponseModalities         []Modality       `json:"responseModalities,omitempty"`
	CandidateCount             *int             `json:"candidateCount,omitempty"`
	Seed                       *int64           `json:"seed,omitempty"`
	PresencePenalty            *float64         `json:"presencePenalty,omitempty"`
	FrequencyPenalty           *float64         `json:"frequencyPenalty,omitempty"`
	ResponseLogprobs           *bool            `json:"responseLogprobs,omitempty"`
	Logprobs                   *int             `json:"logprobs,omitempty"`
	EnableEnhancedCivicAnswers *bool            `json:"enableEnhancedCivicAnswers,omitempty"`
	SpeechConfig               *SpeechConfig    `json:"speechConfig,omitempty"`
	ThinkingConfig             *ThinkingConfig  `json:"thinkingConfig,omitempty"`
	MediaResolution            *MediaResolution `json:"mediaResolution,omitempty"`
}

type Schema struct {
	Type             string             `json:"type,omitempty"`             // "object", "array", "string", "number", "boolean", "null"
	Properties       map[string]*Schema `json:"properties,omitempty"`       // Properti untuk tipe "object"
	Items            *Schema            `json:"items,omitempty"`            // Skema untuk elemen array
	Required         []string           `json:"required,omitempty"`         // Daftar properti yang wajib ada
	Enum             []string           `json:"enum,omitempty"`             // Nilai yang diperbolehkan
	Format           string             `json:"format,omitempty"`           // Format khusus, seperti "date-time", "email"
	Minimum          *float64           `json:"minimum,omitempty"`          // Nilai minimum untuk tipe "number"
	Maximum          *float64           `json:"maximum,omitempty"`          // Nilai maksimum untuk tipe "number"
	MinItems         *int               `json:"minItems,omitempty"`         // Jumlah minimum elemen dalam array
	MaxItems         *int               `json:"maxItems,omitempty"`         // Jumlah maksimum elemen dalam array
	Nullable         *bool              `json:"nullable,omitempty"`         // Apakah nilai bisa null
	Description      string             `json:"description,omitempty"`      // Deskripsi skema
	AnyOf            []*Schema          `json:"anyOf,omitempty"`            // Valid jika cocok dengan salah satu skema
	PropertyOrdering []string           `json:"propertyOrdering,omitempty"` // Urutan properti dalam output
}


type SpeechConfig struct {
	Voice           *string  `json:"voice,omitempty"`
	Pitch           *float64 `json:"pitch,omitempty"`
	SpeakingRate    *float64 `json:"speakingRate,omitempty"`
	VolumeGainDb    *float64 `json:"volumeGainDb,omitempty"`
	AudioEncoding   *string  `json:"audioEncoding,omitempty"` // Contoh: "LINEAR16", "MP3"
	SampleRateHertz *int     `json:"sampleRateHertz,omitempty"`
}

type ThinkingConfig struct {
	// Placeholder: belum ada detail resmi, isi sesuai fitur yang muncul.
	EnableStepByStepReasoning *bool `json:"enableStepByStepReasoning,omitempty"`
	DebugMode                 *bool `json:"debugMode,omitempty"`
}

type Modality string

const (
	ModalityText  Modality = "TEXT"
	ModalityImage Modality = "IMAGE"
	ModalityAudio Modality = "AUDIO"
	ModalityVideo Modality = "VIDEO"
)

type MediaResolution string

const (
	ResolutionLow    MediaResolution = "LOW"
	ResolutionMedium MediaResolution = "MEDIUM"
	ResolutionHigh   MediaResolution = "HIGH"
)

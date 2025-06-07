package geminimodels

import (
	"encoding/base64"
	"log"
	"os"
	"path/filepath"

	geminiutilsetc "github.com/hend41234/gemini/geminiutils/etc"
	"github.com/hend41234/gemini/multimodial"
)

var ConfigRequest *BaseRequestModel

// quick generate config, even though you can still edit the config, like add the params GenerationConfig or else.
//
// the config saved in :
//
//	geminimodels.ConfigRequest
//
// text params it will use to prompt
//
//	'if you want to run as Streaming, the text will be reset'
//
//	nameFileToUpload = ["name file which to upload"]
//
// the nameFileToUpload is mean when you will sending multimodial, or sending media, thats support image, audio, and video.
func QuickGenerateConfigRequest(text string, nameFileToUpload ...string) {
	ConfigRequest = new(BaseRequestModel)
	// var content []Content
	var parts []Part

	// Content, and Parts
	{
		part := Part{Text: text} // set text input / prompt
		var ctn Content
		if len(nameFileToUpload) > 0 { // if used inlineData
			file, err := os.ReadFile(nameFileToUpload[0])
			if err != nil {
				log.Fatal("file not found, please check params nameFile")
			}
			b64Data := base64.StdEncoding.EncodeToString(file)
			inline := InlineData{Data: b64Data}
			inline.MimeType = detectMimeType(nameFileToUpload[0])
			part.InlineDatas = &inline
		}

		parts = append(parts, part)

		ctn.Parts = parts
		rl := "user"
		ctn.Role = &rl

		ConfigRequest.Contents = append(ConfigRequest.Contents, ctn)
	}
}

func detectMimeType(nameFile string) string {
	ext := filepath.Ext(nameFile)
	if geminiutilsetc.Contains(ext, multimodial.MediaSuport.Media.Image) {
		return "image/" + ext
	}
	if geminiutilsetc.Contains(ext, multimodial.MediaSuport.Media.Audio) {
		return "audio/" + ext
	}
	if geminiutilsetc.Contains(ext, multimodial.MediaSuport.Media.Video) {
		return "video/" + ext
	}
	log.Fatal("file not supported")
	return ""
}

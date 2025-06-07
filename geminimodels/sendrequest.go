package geminimodels

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/fatih/color"
	geminiutilsetc "github.com/hend41234/gemini/geminiutils/etc"
)

func sendRequest(url string, body BaseRequestModel) bool {
	byteBody, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(byteBody))
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal("error: " + err.Error())
		return false
	}
	defer res.Body.Close()
	// fmt.Println(url)
	// fmt.Println(string(byteBody))
	if res.StatusCode != 200 {
		readAll, _ := io.ReadAll(res.Body)
		fmt.Println("response not 200")
		res.Body.Close()
		fmt.Println(string(readAll))
		return false
	}

	readResp, _ := io.ReadAll(res.Body)
	fmt.Println(string(readResp))
	return true
}

func sendingStream(url string, bodyConf BaseRequestModel, saveContext bool) {
	// printed := fmt.Sprintf("[ %v ]===>  %v", *bodyConf.Contents[0].Role, bodyConf.Contents[0].Parts[0].Text)
	// fmt.Println(printed)
	IDChat := geminiutilsetc.NewUUID()
	if ok := saveContexts(bodyConf, IDChat); !ok {
		log.Fatal("error")
	}
	input := bufio.NewScanner(os.Stdin)
	fmt.Print("[ user ]===> ")
	input.Scan()

	bodyConf.Contents[0].Parts[0].Text = input.Text()
	for {
		byteBody, _ := json.Marshal(bodyConf)
		newReq, _ := http.NewRequest("POST", url, bytes.NewBuffer(byteBody))
		newReq.Header.Set("Content-Type", "application/json")
		fmt.Println(string(byteBody))
		client := http.Client{}
		res, err := client.Do(newReq)
		if err != nil {
			log.Println("error request : " + err.Error())
			continue

		}
		if res.StatusCode != 200 {
			log.Println("error response : ")
			readRes, _ := io.ReadAll(res.Body)
			fmt.Println(string(readRes))
			// break
			continue
		}
		defer res.Body.Close()

		var responseContent string = ""
		var newModelContent Content
		green := color.New(color.FgGreen).SprintFunc()
		fmt.Print(green("[ model ]===> "))

		scanner := bufio.NewScanner(res.Body)
		for scanner.Scan() {
			var resModel ResModels
			var chunk string
			line := scanner.Text()
			if strings.HasPrefix(scanner.Text(), "data: ") {
				// fmt.Println(line[6:])
				chunk = line[6:]
			}
			byteRes := []byte(chunk)
			err := json.NewDecoder(bytes.NewBuffer(byteRes)).Decode(&resModel)
			if err != nil {
				continue
			}
			newModelContent = resModel.Candidates[0].Contents
			content := resModel.Candidates[0].Contents.Parts[0].Text
			responseContent += content
			// print content response
			fmt.Print(green(resModel.Candidates[0].Contents.Parts[0].Text))

		}
		newModelContent.Parts[0].Text = responseContent // this result Content from SSE / chunk

		//  printed model AI
		// printed = fmt.Sprintf("[ %v ]===> %v", *newModelContent.Role, responseContent)
		// fmt.Println(printed)
		bodyConf.Contents = append(bodyConf.Contents, newModelContent) // save response from model gemini

		var newContent Content
		fmt.Print("[ user ]===> ")
		if !input.Scan() {
			break
		}
		inputText := input.Text()
		if inputText == "exit" || inputText == "quit" {
			if ok := saveContexts(bodyConf, IDChat); !ok {
				log.Fatal("error save context")
			}
			break
		}
		rl := "user"
		newContent.Role = &rl
		newContent.Parts = append(newContent.Parts, Part{Text: inputText})
		if saveContext {
			bodyConf.Contents = append(bodyConf.Contents, newContent)
		} else {
			bodyConf.Contents = []Content{newContent}
		}

		if ok := saveContexts(bodyConf, IDChat); !ok {
			log.Fatal("error save context")
		}
		// fmt.Println("______________________________________________________________________________________________")
	}

}

func saveContexts(body BaseRequestModel, IDChat string) bool {
	return true
}

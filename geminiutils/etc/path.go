package geminiutilsetc

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func IsExist(namePath string) bool {
	_, err := os.Stat(namePath)
	return os.IsExist(err)
}

func CreatePath(namePath string) bool {
	if err := os.Mkdir(namePath, 0755); err != nil {
		return false
	}
	return true
}

func SaveFile(nameFile string, data []byte) bool {
	err := ioutil.WriteFile(nameFile, data, 0755)
	return err == nil
}

func ExtractPathFileInString(pathFile string) []string {
	dir := filepath.Dir(pathFile)
	file := filepath.Base(pathFile)
	return []string{dir, file}
}

func ContentOfDirectory(nameDir string) (listContent []string) {
	list, err := os.ReadDir(nameDir)
	if err != nil {
		log.Fatal("error, please check name of directory")
	}
	for _, content := range list {
		listContent = append(listContent, content.Name())
	}
	return
}

type DataHistory struct {
	ID      string `json:"id"`
	Context string `json:"context"`
}
type ListOfHistory struct {
	Data []DataHistory `json:"data"`
}

func AddListOfContext(IDChat string, context string) bool {
	filePath := "data/listhistory.json"
	var data []DataHistory
	data = append(data, DataHistory{ID: IDChat, Context: context})
	newList := ListOfHistory{Data: data}

	{
		// check dir data
		extracted := ExtractPathFileInString(filePath)

		_, err := os.Stat(filePath)
		if os.IsNotExist(err) {

			if err := os.Mkdir(extracted[0], 0755); err != nil {
				log.Println("error create directory ", extracted[0])
				return false
			}
		}
		// check existing file listhistory.json
		_, err = os.Stat(filePath)
		if os.IsNotExist(err) {
			byteData, err := json.Marshal(newList)
			if err != nil {
				log.Println("error encode")
				return false
			}
			if err := ioutil.WriteFile(filePath, byteData, 0755); err != nil {
				log.Println("error write file ", filePath)
				return false
			}
			return true
		}
	}

	{
		// if fixed file to save is exist
		fileJson, err := os.Open(filePath)
		var newData ListOfHistory
		if err != nil {
			log.Println("error open file listhistory.json")
			return false
		}
		defer fileJson.Close()

		err = json.NewDecoder(fileJson).Decode(&newData)
		if err != nil {
			log.Println("error decode listhistory.json")
			return false
		}

		newData.Data = append(newData.Data, newList.Data[0])
		newByteData, err := json.Marshal(newData)

		if err != nil {
			log.Println("error decode newByteData")
			return false
		}
		err = ioutil.WriteFile(filePath, newByteData, 0755)
		if err != nil {
			log.Println("error save " + filePath)
			return false
		}
	}
	return true

}

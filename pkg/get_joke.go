package pkg

import (
	"VKBotAPI/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

// не помню чем руководствовался пока писал эту функцию, но помню что сайт возвращает ответ
// в кодировке win1251 И С КОРЯВЫМ JSON и из-за этого были проблемы.
func GetJoke() (msg string, err error) {

	resp, err := http.Get("http://rzhunemogu.ru/RandJSON.aspx?CType=11")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	enc := charmap.Windows1251.NewDecoder()
	win, err := enc.Bytes(bodyBytes)
	if err != nil {
		return "", err
	}

	var joke models.Joke

	bodyString := string(win)

	bodyString = strings.Replace(bodyString, "\r", "\\r", -1)
	bodyString = strings.Replace(bodyString, "\n", "\\n", -1)
	bodyString = strings.Replace(bodyString, "\t", "\\t", -1)

	if ok := json.Unmarshal([]byte(bodyString), &joke); ok != nil {

		bodyString = strings.Replace(bodyString, "\\r", "\r", -1)
		bodyString = strings.Replace(bodyString, "\\n", "\n", -1)
		bodyString = strings.Replace(bodyString, "\\t", "\t", -1)
		joke.Content = "2\n" + bodyString[12:len(bodyString)-3]
	}

	msg = fmt.Sprintf("%s", joke.Content)
	return msg, nil
}

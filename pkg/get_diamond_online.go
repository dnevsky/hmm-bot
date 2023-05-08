package pkg

import (
	"errors"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
)

func GetOnlineDiamond() (players map[string]int, err error) {
	resp, err := http.Get("https://diamondrp.ru/")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	bodyString := string(bodyBytes)

	re, err := regexp.Compile(`<p>(?P<first>\w+)<br \/><small>(?P<second>\d+) \/ \d+<\/small><\/p>`)
	if err != nil {
		return nil, err
	}

	res := re.FindAllStringSubmatch(bodyString, -1)
	if len(res) == 0 {
		err = errors.New("getOnlineDPR: minimum 1 expected, 0 received in body request")
		return nil, err
	}

	result := make(map[string]int)

	for i := 0; i <= (len(res)/2)-1; i++ {
		ii, _ := strconv.Atoi(res[i][2])
		result[res[i][1]] = ii
	}

	return result, nil
}

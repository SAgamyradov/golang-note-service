package service

import (
	"encoding/json"
	"net/http"
	"net/url"
)

const YANDEX_SPELL_CHECK_URL = "https://speller.yandex.net/services/spellservice.json/checkText"

type SpellCheckResponse struct {
	Code        int      `json:"code"`
	Pos         int      `json:"pos"`
	Word        string   `json:"word"`
	Suggestions []string `json:"s"`
}

func CheckSpelling(text string) (bool, error) {
	resp, err := http.PostForm(YANDEX_SPELL_CHECK_URL, url.Values{"text": {text}})
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	var findings []SpellCheckResponse
	if err := json.NewDecoder(resp.Body).Decode(&findings); err != nil {
		return false, err
	}

	// Если найдены ошибки, вернем false
	return len(findings) == 0, nil
}

package tts

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/Yobubble/yona-bot/config"
	"github.com/Yobubble/yona-bot/internal/log"
)

type vve struct {
	baseUrl string
}

func (v *vve) audioQuery(text string, speakerId int) ([]byte, error) {
	url := fmt.Sprintf(v.baseUrl+"/audio_query?text=%s&speaker=%d", url.QueryEscape(text), speakerId)
	res, err := http.Post(url, "application/json", nil)
	if err != nil {
		log.Sugar.Warn("Error get audio query from preset")
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Sugar.Warn("Error reading audio query response body")
		return nil, err
	}

	return body, nil
}

func (v *vve) TextToSpeech(text string, outputPath string) error {
	query, err := v.audioQuery(text, 8)
	if err != nil {
		return err
	}

	url := fmt.Sprintf(v.baseUrl+"/synthesis?speaker=%d&enable_interrogative_upspeak=%t", 8, true)
	res, err := http.Post(url, "application/json", bytes.NewBuffer(query))
	if err != nil {
		log.Sugar.Warn("Error sending audio synthesis request")
		return err
	}
	defer res.Body.Close()

	out, err := os.Create(outputPath)
	if err != nil {
		log.Sugar.Warn("Error creating output file")
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, res.Body)
	if err != nil {
		log.Sugar.Warn("Error writing audio data to file")
		return err
	}

	return nil
}

func newVVE(cfg *config.Cfg) TTSModel {
	return &vve{
		baseUrl: cfg.GetVVEBaseUrl(),
	}
}

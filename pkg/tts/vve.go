package tts

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/Yobubble/yona-bot/config"
)

type vve struct {
	baseUrl string
}

func (v *vve) audioQuery(text string, speakerId int) ([]byte, error) {
	url := fmt.Sprintf(v.baseUrl+"/audio_query?text=%s&speaker=%d", url.QueryEscape(text), speakerId)
	res, err := http.Post(url, "application/json", nil)
	if err != nil {
		return nil, fmt.Errorf("text to speech: failed to get audio query: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("text to speech: failed to read audio query response: %w", err)
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
		return fmt.Errorf("text to speech: failed to synthesize audio: %w", err)
	}
	defer res.Body.Close()

	out, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("text to speech: failed to create output file: %w", err)
	}
	defer out.Close()

	_, err = io.Copy(out, res.Body)
	if err != nil {
		return fmt.Errorf("text to speech: failed to write audio data to file: %w", err)
	}

	return nil
}

func newVVE(cfg *config.Cfg) TTSModel {
	return &vve{
		baseUrl: cfg.GetVVEBaseUrl(),
	}
}

package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/Yobubble/yona-bot/config"
	"github.com/Yobubble/yona-bot/internal/log"
)

var Vve *VoicevoxEngineApi

type preset struct {
	ID                int      `json:"id"`
	Name              string   `json:"name"`
	SpeakerUUID       string   `json:"speaker_uuid"`
	StyleID           int      `json:"style_id"`
	SpeedScale        float64  `json:"speedScale"`
	PitchScale        float64  `json:"pitchScale"`
	IntonationScale   float64  `json:"intonationScale"`
	VolumeScale       float64  `json:"volumeScale"`
	PrePhonemeLength  float64  `json:"prePhonemeLength"`
	PostPhonemeLength float64  `json:"postPhonemeLength"`
	PauseLength       *float64 `json:"pauseLength"`
	PauseLengthScale  float64  `json:"pauseLengthScale"`
}

type VoicevoxEngineApi struct {
	baseUrl string
}

func (v *VoicevoxEngineApi) addPreset(preset preset) error {
	jsonBytes, err := json.Marshal(preset)
	if err != nil {
		log.Sugar.Warn("Error convert preset to json")
		return err
	}
	res, err := http.Post(v.baseUrl+"/add_preset", "application/json", bytes.NewBuffer(jsonBytes))
	if err != nil {
		log.Sugar.Warn("Error sending add preset request")
		return err
	}
	defer res.Body.Close()
	return nil
}

func (v *VoicevoxEngineApi) deletePreset(id int) error {
	url := fmt.Sprintf(v.baseUrl+"/delete_preset?id=%d", id)
	res, err := http.Post(url, "application/json", nil)
	if err != nil {
		log.Sugar.Warn("Error sending delete preset request")
		return err
	}
	defer res.Body.Close()
	return nil
}

func (v *VoicevoxEngineApi) getAudioQueryFromPreset(text string) ([]byte, error) {
	url := fmt.Sprintf(v.baseUrl+"/audio_query_from_preset?text=%s&preset_id=%d", url.QueryEscape(text), 1)
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

func (v *VoicevoxEngineApi) TextToAudio(text string, outputPath string) error {
	query, err := v.getAudioQueryFromPreset(text)
	if err != nil {
		return err
	}

	url := fmt.Sprintf(v.baseUrl+"/synthesis?speaker=%d&enable_interrogative_upspeak=%t", 29, true)
	res, err := http.Post(url, "application/json", bytes.NewBuffer(query))
	if err != nil {
		log.Sugar.Warn("Error sending audio synthesis request")
		return err
	}
	defer res.Body.Close()

	// Create output file
	out, err := os.Create(outputPath)
	if err != nil {
		log.Sugar.Warn("Error creating output file")
		return err
	}
	defer out.Close()

	// Copy audio data directly to file
	_, err = io.Copy(out, res.Body)
	if err != nil {
		log.Sugar.Warn("Error writing audio data to file")
		return err
	}

	return nil
}

func (v *VoicevoxEngineApi) CloseVoicevox() error {
	err := v.deletePreset(1)
	if err != nil {
		return err
	}
	return nil
}

func InitVoicevoxEngineApi(cfg *config.Cfg) {
	// Ref: http://localhost:50021/speakers
	preset := preset{
		ID:                1,
		Name:              "Yona",
		SpeakerUUID:       "35b2c544-660e-401e-b503-0e14c635303a",
		StyleID:           8,
		SpeedScale:        1,
		PitchScale:        0.0,
		IntonationScale:   1,
		VolumeScale:       1,
		PrePhonemeLength:  0.1,
		PostPhonemeLength: 0.1,
		PauseLength:       nil,
		PauseLengthScale:  1.0,
	}

	vveApi := &VoicevoxEngineApi{
		baseUrl: cfg.GetVoicevoxEngineBaseUrl(),
	}

	vveApi.addPreset(preset)
	Vve = vveApi
}

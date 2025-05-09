package main

import (
	"fmt"
	"os"
	"strings"
)

type s3Configs struct {
	S3Bucket           string
	AWSAccessKeyId     string
	AWSSecretAccessKey string
	AWSRegion          string
}

var (
	storageChoices = map[string]string{
		"L": "Local",
		"S": "S3",
	}
	langChoices = map[string]string{
		"JP": "Japanese",
	}
	lmChoices = map[string]string{
		"G": "GPT4o",
	}
)

func main() {
	var env strings.Builder

	var discordBotToken string
	var storage string
	var s3 s3Configs
	var lang string
	var lm string
	var openAiAPIKey string
	var vveUrl string

	fmt.Println("Please don't put any whitespace in the input.")

	// --- Discord Bot Token ---
	fmt.Print("Discord Bot Token: ")
	fmt.Scanln(&discordBotToken)

	// --- Storage ---
	var selectedStorage string

	for storageChoices[selectedStorage] == "" {
		fmt.Print("Select your storage (Type 'L' for Local or 'S' for S3): ")
		fmt.Scanln(&selectedStorage)
	}

	if storageChoices[selectedStorage] == "Local" {
		storage = "Local"
	} else {
		storage = "S3"
		fmt.Print("S3 Bucket name: ")
		fmt.Scanln(&s3.S3Bucket)
		fmt.Print("AWS Access Key ID: ")
		fmt.Scanln(&s3.AWSAccessKeyId)
		fmt.Print("AWS Secret Access Key ID: ")
		fmt.Scanln(&s3.AWSSecretAccessKey)
		fmt.Print("AWS Region (e.g., ap-southeast-7): ")
		fmt.Scanln(&s3.AWSRegion)
	}

	// --- Language ---
	var selectedLang string

	for langChoices[selectedLang] == "" {
		fmt.Print("Select Language (Type 'JP' for Japanese): ")
		fmt.Scanln(&selectedLang)
	}

	if langChoices[selectedLang] == "Japanese" {
		lang = "JP"
		vveUrl = "http://vve:50021"
	}

	// --- LM ---
	var selectedLM string
	for lmChoices[selectedLM] == "" {
		fmt.Print("Select LM (Type 'G' for GPT4o): ")
		fmt.Scanln(&selectedLM)
	}

	if lmChoices[selectedLM] == "GPT4o" {
		lm = "GPT4o"
	}

	// --- OpenAI API Key ---
	fmt.Print("OpenAI API Key: ")
	fmt.Scanln(&openAiAPIKey)

	// --- Digest data ---
	env.Write([]byte("DISCORD_BOT_TOKEN=" + discordBotToken + "\n"))
	env.Write([]byte("STORAGE=" + storage + "\n"))

	env.Write([]byte("S3_BUCKET=" + s3.S3Bucket + "\n"))
	env.Write([]byte("AWS_ACCESS_KEY_ID=" + s3.AWSAccessKeyId + "\n"))
	env.Write([]byte("AWS_SECRET_ACCESS_KEY=" + s3.AWSSecretAccessKey + "\n"))
	env.Write([]byte("AWS_REGION=" + s3.AWSRegion + "\n"))

	env.Write([]byte("LANGUAGE=" + lang + "\n"))
	env.Write([]byte("LM=" + lm + "\n"))
	env.Write([]byte("OPENAI_API_KEY=" + openAiAPIKey + "\n"))
	env.Write([]byte("VOICEVOX_ENGINE_URL=" + vveUrl + "\n"))

	// --- Print the finalized configuration ---
	fmt.Printf("Configuration: \n%v\n", env.String())

	// --- Save to .env file ---
	if err := os.WriteFile(".env", []byte(env.String()), 0644); err != nil {
		fmt.Printf("Error writing to .env file: %v\n", err)
	} else {
		fmt.Println(".env file created successfully.")
	}
}

package helper

import (
	"fmt"
	"os"
	"os/exec"
)

func ConvertOggToMp3(inputPath string, outputPath string) error {
	cmd := exec.Command("ffmpeg", "-y", "-i", inputPath, "-vn", "-ar", "48000", "-ac", "2", "-b:a", "192k", outputPath)
	_, err := cmd.Output()
	if err != nil {
		return err
	}

	return nil
}

// Ref: https://github.com/bwmarrin/discordgo/blob/master/examples/airhorn/README.md
func ConvertMp3ToDCA(inputPath string, outputPath string) error {
	cmdFFmpeg := exec.Command("ffmpeg", "-y", "-i", inputPath, "-f", "s16le", "-ar", "48000", "-ac", "2", "pipe:1")
	cmdDCA := exec.Command("dca")

	pipe, err := cmdFFmpeg.StdoutPipe()
	if err != nil {
		return fmt.Errorf("failed to create pipe: %v", err)
	}
	cmdDCA.Stdin = pipe

	output, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %v", err)
	}
	defer output.Close()

	cmdDCA.Stdout = output

	if err := cmdFFmpeg.Start(); err != nil {
		return fmt.Errorf("failed to start ffmpeg: %v", err)
	}
	if err := cmdDCA.Start(); err != nil {
		return fmt.Errorf("failed to start dca: %v", err)
	}

	if err := cmdFFmpeg.Wait(); err != nil {
		return fmt.Errorf("ffmpeg error: %v", err)
	}
	if err := cmdDCA.Wait(); err != nil {
		return fmt.Errorf("dca error: %v", err)
	}

	return nil
}

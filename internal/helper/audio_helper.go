package helper

import (
	"os"
	"os/exec"
)

type AudioHelper struct {
}

// Ref: https://github.com/bwmarrin/discordgo/blob/master/examples/airhorn/README.md
func (a *AudioHelper) ConvertToDCA(inputPath string, outputPath string) error {
	cmdFFmpeg := exec.Command("ffmpeg", "-y", "-i", inputPath, "-f", "s16le", "-ar", "48000", "-ac", "2", "pipe:1")
	cmdDCA := exec.Command("dca")

	pipe, err := cmdFFmpeg.StdoutPipe()
	if err != nil {
		return err
	}
	cmdDCA.Stdin = pipe

	output, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer output.Close()

	cmdDCA.Stdout = output

	if err := cmdFFmpeg.Start(); err != nil {
		return err
	}
	if err := cmdDCA.Start(); err != nil {
		return err
	}

	if err := cmdFFmpeg.Wait(); err != nil {
		return err
	}
	if err := cmdDCA.Wait(); err != nil {
		return err
	}

	return nil
}

func (a *AudioHelper) ConvertToMp3(inputPath string, outputPath string) error {
	cmd := exec.Command("ffmpeg", "-y", "-i", inputPath, "-vn", "-ar", "48000", "-ac", "2", "-b:a", "192k", outputPath)

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

package listening

import (
	"bufio"
	"log"
	"luna/src/utils"
	"os/exec"
	"strings"
)

// ListenForCommand listens for a command and returns it
func ListenForCommand() string {
	cmd := exec.Command("sh", "scripts/run_whisper.sh")
	stdout, err := cmd.StdoutPipe()
	message := ""
	if err != nil {
		log.Fatalf("Erro ao obter stdout: %v", err)
	}
	if err := cmd.Start(); err != nil {
		log.Fatalf("Erro ao iniciar o comando: %v", err)
	}

	scanner := bufio.NewScanner(stdout)
	for message == "" && scanner.Scan() {
		utils.Debug("[listening]", "Listening for command...")

		textSlices := strings.Split(scanner.Text(), string(rune(27)))
		text := textSlices[getLastIndex(len(textSlices))]
		matches := detectionRegex.FindAllStringSubmatch(strings.ReplaceAll(text, string(rune(27)), ""), 1)

		if len(matches) > 0 {
			cmd.Process.Release()
			speech := removeRunes(matches[0][2], speechRunes)[2:]
			utils.Debug("[speech]", speech)

			calledMatches := calledDetectionRegex.FindAllStringSubmatch(speech, 1)
			if len(calledMatches) > 0 {
				utils.Debug("[speech]", "[called]", speech)
				keyword := calledMatches[0][0]
				message = strings.ReplaceAll(speech, keyword, "")
			} else {
				utils.Debug("[speech]", "[not called]", speech)
			}
		} else {
			utils.Debug(text)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Erro ao ler a saída: %v", err)
	}

	return message
}

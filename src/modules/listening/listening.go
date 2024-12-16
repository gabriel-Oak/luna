package main

import (
	"bufio"
	"log"
	"os/exec"
)

func main() {
	// Cria o comando
	cmd := exec.Command("sh", "run_whisper.sh")

	// Obtém o stdout do comando
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalf("Erro ao obter stdout: %v", err)
	}

	// Inicia o comando
	if err := cmd.Start(); err != nil {
		log.Fatalf("Erro ao iniciar o comando: %v", err)
	}

	// Cria um scanner para ler a saída continuamente
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		log.Println(scanner.Text())
	}

	// Verifica se houve algum erro ao ler a saída
	if err := scanner.Err(); err != nil {
		log.Fatalf("Erro ao ler a saída: %v", err)
	}

	// Espera o comando terminar
	if err := cmd.Wait(); err != nil {
		log.Fatalf("Erro ao esperar o comando terminar: %v", err)
	}
}

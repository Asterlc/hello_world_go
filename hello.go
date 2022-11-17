package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const ARQUIVO_TXT = "sites.txt"

func main() {
	welcome()
	menu()
	for {
		comando()
	}
}

func welcome() {
	nome := "Lucas"
	fmt.Println("Bem vindo", nome)
}

func menu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("3 - Sair do programa")
}

func comando() int {
	var comando int
	fmt.Scan(&comando)

	switch comando {
	case 1:
		monitoramento()
	case 2:
		exibirLogs()
	case 3:
		sairApp()
	default:
		fmt.Println("Comando desconhecido - Fechando App")
		os.Exit(0)
	}

	return comando

}

func monitoramento() {
	fmt.Println("Iniciando monitoramento...")
	sites, error := monitorados()
	if error != nil {

		return
	}
	testeStatus(sites)
}

func exibirLogs() {
	fmt.Println("Exibindo logs")
	menu()
}

func sairApp() {
	fmt.Println("Saindo do programa...")
	os.Exit(0)
}

func monitorados() ([]string, error) {
	sites, error := lerTXT()
	return sites, error
}

func testeStatus(sites []string) {
	for _, site := range sites {
		response, error := http.Get(site)

		if response.StatusCode > 400 || error != nil {
			fmt.Println("Site:", site, "Falha carregamento:", error)
		} else {
			fmt.Println("Site:", site, "carregado com sucesso!", response.StatusCode)
			registraLog(site, response.StatusCode)
		}
	}
	fmt.Println()
	menu()
}

func lerTXT() ([]string, error) {
	fmt.Println("Checando arquivos...")
	// var sites []string
	var sites = make([]string, 0)
	// var sites = []string{}

	arquivos, error := os.Open(ARQUIVO_TXT)
	if error != nil {
		fmt.Println("Erro de leitura no arquivo TXT", error)
		return sites, error
	}
	leitor := bufio.NewReader(arquivos)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)
		if err == io.EOF {
			break
		}
	}
	fmt.Println("Arquivos checados...")
	arquivos.Close()
	return sites, error
}

func registraLog(site string, status int) {
	file, err := os.OpenFile("healthlog.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Erro para registrar log", err)
	} else {
		fmt.Println("Registrando log:", site)
		date := time.Now().UTC().String()
		fmt.Println(date)
		formatStatus := strconv.Itoa(status)
		file.WriteString(date + " " + site + " " + " " + "status" + " " + formatStatus + "\n")
		file.Close()
	}
}

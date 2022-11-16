package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	lerTXT()
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
		fmt.Println("Comando desconhecido")
		os.Exit(0)
	}

	return comando

}

func monitoramento() {
	fmt.Println("Iniciando monitoramento...")
	sites := monitorados()
	testeStatus(sites)
}

func exibirLogs() {
	fmt.Println("Exibindo logs")
}

func sairApp() {
	fmt.Println("Saindo do programa...")
	os.Exit(0)
}

func monitorados() []string {
	sites := lerTXT()
	return sites
}

func testeStatus(sites []string) {
	for _, site := range sites {
		response, error := http.Get(site)

		if response.StatusCode > 400 || error != nil {
			fmt.Println("Site:", site, "Falha carregamento:", error)
		} else {
			fmt.Println("Site:", site, "carregado com sucesso!", response.StatusCode)
		}
	}
}

func lerTXT() []string {
	var sites []string
	arquivos, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
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

	return sites
}

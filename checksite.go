package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const SITES_TXT = "sites.txt"
const LOG_TXT = "log.txt"

// const UNIX_HELPER = 1000000 formatçaão em unix time

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
		fmt.Println(error)
		return
	}
	testeStatus(sites)
}

func exibirLogs() {
	fmt.Println("Exibindo logs")
	file, err := ioutil.ReadFile(LOG_TXT)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(string(file))
	}
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
		}
		registraLog(site, response.StatusCode)
	}
	fmt.Println()
	menu()
}

func lerTXT() ([]string, error) {
	fmt.Println("Checando arquivos...")
	// var sites []string
	var sites = make([]string, 0)
	// var sites = []string{}

	files, error := os.Open(SITES_TXT)
	if error != nil {
		fmt.Println("Erro de leitura no arquivo TXT", error)
		return sites, error
	}
	leitor := bufio.NewReader(files)

	for {
		line, err := leitor.ReadString('\n')
		line = strings.TrimSpace(line)
		sites = append(sites, line)
		if err == io.EOF {
			break
		}
	}
	fmt.Println("Arquivos checados...")
	files.Close()
	return sites, error
}

func registraLog(site string, status int) {
	file, err := os.OpenFile(LOG_TXT, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Erro para registrar log", err)
		file.Close()
	} else {
		formatStatus := strconv.Itoa(status)
		// dateFormat := strconv.Itoa(int(time.Now().UnixNano() / UNIX_HELPER))
		dateFormat := time.Now().Format("02/01/2006 15:04:05")

		file.WriteString(dateFormat + " - " + site + " - " + "status" + " - " + formatStatus + "\n")
		file.Close()
	}
}

package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
)

func main() {
	lerTXT()
	welcome()
	menu()
	for {
		decisao(comando())
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

func decisao(comando int) {
	println("Comando", comando)
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
	sites := []string{
		"http://random-status-code.herokuapp.com/",
		"https://app.fluigidentity.net/ui/login",
	}

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
	// arquivos, err := ioutil.ReadFile("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
	linha, err := bufio.NewReader(arquivos).ReadString('\n')
	if err != nil {
		fmt.Println("Erro bufio:", err)
	}
	fmt.Println("linha:", linha)
	// fmt.Println(string(arquivos))

	return sites
}

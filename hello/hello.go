package main

import (
	"bufio" //Pacote para ler programas de diferentes formar ex: linha a linhas
	"fmt"   //Mostrar coisas na tela
	"io"
	"io/ioutil" //Para ler arquivos e mostrar na tela
	"net/http"  //usar a internet
	"os"        //Pacote do sistema operacional
	"strconv"   //Conversor de qualquer coisa para string
	"strings"
	"time" //Pacote de tempo
)

const monitoramentos = 3
const delay = 5

func main() {
	exibeIntroducao()
	for {

		exibeMenu()
		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			imprimeLogs()
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}
	}

}

func exibeIntroducao() {
	nome := "Douglas"
	versao := 1.5

	fmt.Println("Olá se.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)

	fmt.Println("O comando escolhido foi", comandoLido)
	fmt.Println("")
	return comandoLido
}

func devolveNomeEIdade() (string, int) {
	nome := "Erick"
	idade := 21
	return nome, idade
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	// var sites = [3]string
	// sites [0] = "https://www.cocamar.com.br"
	// sites [1] = "https://alura.com"
	// sites [2] = "https://google.com"
	// fmt.Println(sites)

	//sites := []string{"https://www.cocamar.com.br", "https://alura.com", "https://google.com"}

	sites := leSitesDoArquivo()

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Testando", i, ":", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
	fmt.Println("")
}

func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("Site:", site, "esta com problemas. Status Code:", resp.StatusCode)
		registraLog(site, false)
	}
}

func leSitesDoArquivo() []string {

	sites := []string{}
	arquivo, err := os.Open("sites.txt") //Para abrir um arquivo

	//arquivo, err := ioutil.ReadFile("sites.txt")

	leitor := bufio.NewReader(arquivo)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)

	}

	for {
		linha, err := leitor.ReadString('\n') //Lendo até o final da linha
		linha = strings.TrimSpace(linha)
		fmt.Println(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break
		}
	}
	arquivo.Close()
	return sites
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " " + site + " - online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLogs() {

	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))
}

package main

import (
	"strings"
	"time"
	"fmt" // interaçao com strings
	"os" // interaçao com o sistema operativo
	"net/http" // interaçao com a web
	"bufio"
	"io"
	"strconv"
	"io/ioutil"
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
			fmt.Println("Exibindo Logs..")
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
	versao := 1.1
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa esta na versão", versao)
}

func exibeMenu()  {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do Programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O Comando escolhifo foi", comandoLido)

	return comandoLido
}


func iniciarMonitoramento()  {
	fmt.Println("Monitorando...")
	
	sites := leSitesDoArquivo()
	//  sites := []string{"https://www.alura.com.br",
	// 	"https://caelum.com.br", "https://random-status-code.herokuapp.com/"}
	
	// for i := 0; i < len(sites); i++ {
	// 	   fmt.Println(sites[i])
	// }
	for i:= 0; i < monitoramentos; i++ {
		for i, site := range(sites) {
			fmt.Println("Posiçao", i, "Testando site:", site)
			testaSite(site)
		}
		time.Sleep( delay * time.Second)
		fmt.Println()
	}
	fmt.Println()
}

func testaSite(site string)  {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}
	
	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		registraLog(site, true)
    } else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
		registraLog(site, false)
    }
}

func leSitesDoArquivo() []string {
	var sites []string
	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro", err)

	}
	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)

		if err == io.EOF{
			break
		}
	}

	arquivo.Close()
	return sites
}

func registraLog(site string, status bool)  {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	arquivo.WriteString( time.Now().Format("02/01/2006 15:04:05") + " - " + site + "  -- online: " + strconv.FormatBool(status) + "\n")
	arquivo.Close()
}

func imprimeLogs()  {
	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(arquivo))
	
}
package main

import (
	"time"
	"fmt" // interaçao com strings
	"os" // interaçao com o sistema operativo
	"net/http" // interaçao com a web
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
	
	sites := []string{"https://www.alura.com.br",
		"https://caelum.com.br", "https://random-status-code.herokuapp.com/"}
	
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
	resp, _ := http.Get(site)
	
	if resp.StatusCode == 200 {
        fmt.Println("Site:", site, "foi carregado com sucesso!")
    } else {
        fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
    }
}
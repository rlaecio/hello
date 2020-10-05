package main

import "fmt" // interaçao com strings
import "os" // interaçao com o sistema operativo
import "net/http" // interaçao com a web


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
	site := "https://www.alura.com.br"
	
	resp, _ := http.Get(site)
	fmt.Println(resp)

	if resp.StatusCode == 200 {
        fmt.Println("Site:", site, "foi carregado com sucesso!")
    } else {
        fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
    }

}
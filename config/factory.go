package config

import (
	"flag"
	"fmt"
)

func BuildCmdConfigParams() *ConfigParams {

	var (
		url         string
		requests    int
		concurrency int
	)

	flag.StringVar(&url, "url", "", "URL do serviço a ser testado")
	flag.IntVar(&requests, "requests", 0, "Número total de requests")
	flag.IntVar(&concurrency, "concurrency", 0, "Número de chamadas simultâneas")
	flag.Parse()

	fmt.Println("Configurado com os parâmetros:", url, requests, concurrency)

	return &ConfigParams{
		Url:         url,
		Requests:    requests,
		Concurrency: concurrency,
	}
}

package main

import (
	"fmt"
	"stress-test/config"
	"stress-test/report"
	"stress-test/runner"
)

func main() {

	cfg := config.BuildCmdConfigParams()
	config.ValidateConfig(cfg)

	results := runner.RunLoadTest(cfg)
	report.GenerateReport(results)

	fmt.Println("Teste de carga concluído com sucesso!")
}

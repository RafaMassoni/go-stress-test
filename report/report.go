package report

import (
	"fmt"
	"stress-test/runner"
)

func GenerateReport(results runner.TestResults) {
	// Implementar a geração de relatório aqui
	fmt.Println("\n################# Relatório de Teste de Carga: #################\n")
	fmt.Printf("- Tempo total de execução: %v\n", results.Duration.Seconds())
	fmt.Printf("- Quantidade total de requests realizados: %d\n", results.TotalRequests)
	fmt.Printf("- Quantidade de requests com status HTTP 200: %d\n", results.StatusCodesCount[200])

	// Mostrar distribuição de outros códigos de status
	fmt.Println("\nDistribuição de códigos de status HTTP:")
	for code, count := range results.StatusCodesCount {
		if code != 200 {
			fmt.Printf("-Status %d: %d\n", code, count)
		}
	}
	fmt.Println("\n--------------------------------------------------------")
}

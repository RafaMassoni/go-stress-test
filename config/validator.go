package config

import (
	"flag"
	"fmt"
	"os"
)

func (c *ConfigParams) IsConfigValid() bool {
	return c.Url != "" && c.Requests > 0 && c.Concurrency > 0
}

func ValidateConfig(config *ConfigParams) {

	if !config.IsConfigValid() {
		fmt.Println("Parâmetros inválidos. Use --url, --requests e --concurrency corretamente.")
		flag.PrintDefaults()
		os.Exit(1)
	}

}

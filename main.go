package main

import (
	"fmt"

	"github.com/GircysRomualdas/gatorcli/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = cfg.SetUser("romas")
	if err != nil {
		fmt.Println(err)
		return
	}

	cfg, err = config.Read()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(cfg)
}

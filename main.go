package main

import (
	"fmt"
	"net/http"

	"github.com/hashicorp/vault/api"
)

func getVaultConfig() map[string]interface{} {
	client, _ := api.NewClient(&api.Config{Address: "http://vault:8200"})
	client.SetToken("s.1234567890abcdef")
	secret, _ := client.Logical().Read("secret/data/myapp")
	return secret.Data["data"].(map[string]interface{})
}

func main() {
	cfg := getVaultConfig()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get(cfg["secretHeader"].(string)) == cfg["secretPass"].(string) {
			fmt.Fprintf(w, "adminToken: %s", cfg["adminToken"].(string))
		} else {
			fmt.Fprintf(w, "Привет, мир from go!")
		}
	})

	port := "8080"
	if cfg["port"] != nil {
		port = cfg["port"].(string)
	}
	http.ListenAndServe(":"+port, nil)
}

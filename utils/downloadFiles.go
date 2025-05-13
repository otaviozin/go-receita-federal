package utils

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
)

func downloadFile(dir, filename string) {
	url := fmt.Sprintf("https://arquivos.receitafederal.gov.br/dados/cnpj/dados_abertos_cnpj/2025-04/%s", filename)
	destPath := fmt.Sprintf("%s/%s", dir, filename)

	out, err := os.Create(destPath)
	if err != nil {
		log.Printf("Erro ao criar %s: %s\n", filename, err)
		return
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Erro ao baixar %s: %s\n", filename, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Erro HTTP ao baixar %s: status %d\n", filename, resp.StatusCode)
		return
	}

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Printf("Erro ao salvar %s: %s\n", filename, err)
		return
	}

	log.Printf("Download de %s concluído.\n", filename)
}

func DownloadEmpresas() {
	var wg sync.WaitGroup

	dir := "downloads"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.Mkdir(dir, 0755); err != nil {
			log.Fatalf("Erro ao criar diretório: %s", err)
		}
	}

	for i := 0; i < 10; i += 2 {
		wg.Add(2)

		i0 := i
		i1 := i + 1

		go func() {
			defer wg.Done()
			downloadFile(dir, fmt.Sprintf("Empresas%d.zip", i0))
		}()

		go func() {
			defer wg.Done()
			downloadFile(dir, fmt.Sprintf("Empresas%d.zip", i1))
		}()

		wg.Wait()
	}

	log.Println("Todos os downloads foram concluídos.")
}

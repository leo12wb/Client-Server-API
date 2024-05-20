package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type Quote struct {
    Bid float64 `json:"bid"`
}

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
    defer cancel()

    req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
    if err != nil {
        log.Fatal(err)
    }

    client := http.Client{
        Timeout: 300 * time.Millisecond,
    }
    resp, err := client.Do(req)
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        log.Fatalf("Erro ao obter cotação: %s", resp.Status)
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }

    var quote Quote
    if err := json.Unmarshal(body, &quote); err != nil {
        log.Fatal(err)
    }

    // Salva a cotação em um arquivo
    err = saveQuoteToFile(quote)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Cotação do dólar: %.2f\n", quote.Bid)
}

func saveQuoteToFile(quote Quote) error {
    file, err := os.Create("cotacao.txt")
    if err != nil {
        return err
   }
    defer file.Close()

    _, err = fmt.Fprintf(file, "Dólar: %.2f\n", quote.Bid)
    if err != nil {
        return err
   }

    return nil
}

package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type Quote struct {
    Bid float64 `json:"bid"`
}

func main() {
    db, err := sqlx.Open("sqlite3", "./quotes.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    http.HandleFunc("/cotacao", func(w http.ResponseWriter, r *http.Request) {
        ctx, cancel := context.WithTimeout(r.Context(), 200*time.Millisecond)
        defer cancel()

        select {
        case <-ctx.Done():
            log.Println("Timeout ao buscar cotação do dólar")
            http.Error(w, "Timeout ao buscar cotação do dólar", http.StatusInternalServerError)
            return
        default:
            quote, err := fetchQuote(ctx)
            if err != nil {
                log.Println("Erro ao buscar cotação do dólar:", err)
                http.Error(w, "Erro ao buscar cotação do dólar", http.StatusInternalServerError)
                return
            }
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(quote)

            // Salva a cotação no banco de dados
            if err := saveQuote(ctx, db, quote); err != nil {
                log.Println("Erro ao salvar cotação no banco de dados:", err)
                return
            }
        }
    })

    log.Fatal(http.ListenAndServe(":8080", nil))
}

func fetchQuote(ctx context.Context) (*Quote, error) {
    client := http.Client{
        Timeout: 200 * time.Millisecond,
    }
    req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
    if err != nil {
        return nil, err
    }

    resp, err := client.Do(req)
    if err != nil {
        return nil, err
   }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    var data map[string]Quote
    if err := json.Unmarshal(body, &data); err != nil {
        return nil, err
    }

    return &data["USDBRL"], nil
}

func saveQuote(ctx context.Context, db *sqlx.DB, quote *Quote) error {
    _, err := db.ExecContext(ctx, "INSERT INTO quotes (bid) VALUES (?)", quote.Bid)
    if err != nil {
        return err
    }
    return nil
}

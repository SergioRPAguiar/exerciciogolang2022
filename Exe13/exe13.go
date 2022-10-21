package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"strconv"
)

func main() {
	brlUsd()
}

type Conversao struct {
	Code        string `json:"code"`
	Codein      string `json:"codein"`
	Name        string `json:"name"`
	High        string `json:"high"`
	Low         string `json:"low"`
	VarBid      string `json:"varBid"`
	PctChange   string `json:"pctChange"`
	Bid         string `json:"bid"`
	Ask         string `json:"ask"`
	Timestamp   string `json:"timestamp"`
	Create_date string `json:"create_date"`
}

func brlUsd() {
	url := "https://economia.awesomeapi.com.br/json/BRL-USD/1"

	res, errReq := http.Get(url)
	if errReq != nil {
		fmt.Println("Erro na requisição")
		return
	}

	data, errBody := ioutil.ReadAll(res.Body)
	if errBody != nil {
		fmt.Println("Erro na resposta")
		return
	}

	conversoes := []Conversao{}
	errConversoes := json.Unmarshal(data, &conversoes)
	if errConversoes != nil {
		fmt.Println("Erro na conversão", errConversoes)
		return
	}
	fmt.Println("Conversor de Real para Dolar")
	fmt.Println("Digite quantos R$ deseja converter para $: ")
	var real float64
	fmt.Scan(&real)
	dolar := conversoes[0].Bid
	dolar1, errConver := strconv.ParseFloat(dolar, 2)
	if errConver != nil {
		fmt.Println("Erro na conversão para dolar", errConver)
		return
	}
	resultado := dolar1 * real
	resultado = math.Round(resultado*100) / 100
	fmt.Println("R$", real, "= $", resultado)
}

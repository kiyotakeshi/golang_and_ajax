package main

import (
	"encoding/json"
	"log"
	"os"
)

type Person struct {
	ID      int    `json:"id"`
	Name    string `json:"name"` // nameというキーで格納
	Email   string `json:"-"`    // JSONに格納しない
	Age     int    `json:"age"`
	Address string `json:"address,omitempty"` // 値が空なら出力しない
	memo    string
}

func main() {
	person := &Person{
		ID:      1,
		Name:    "Gopher",
		Email:   "gopher@example.org",
		Age:     10,
		Address: "",
		memo:    "golang lover",
	}

	// ファイルを開く
	file, err := os.Create("./create_by_go.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// エンコーダの取得
	encoder := json.NewEncoder(file)

	// JSONエンコードしたデータの書き込み
	err = encoder.Encode(person)
	if err != nil {
		log.Fatal(err)
	}

}

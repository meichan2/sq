package main

import (
    "net/http"
    "io/ioutil"
    "fmt"
)


func main() {
    url := "http://localhost:8080"  // アクセスするURLだよ！

    resp, err := http.Get(url)      // GETリクエストでアクセスするよ！
    if err != nil {                 // err ってのはエラーの時にエラーの内容が入ってくるよ！
        panic(err)                  // panicは処理を中断してエラーの中身を出力するよ！
    }
    defer resp.Body.Close()         // 関数が終了するとなんかクローズするよ！（おまじない的な）

    byteArray, err := ioutil.ReadAll(resp.Body) // 帰ってきたレスポンスの中身を取り出すよ！
    if err != nil {
        panic(err)
    }
    fmt.Println(string(byteArray))  // 取り出したリクエスト結果をバイナリ配列からstring型に変換して出力するよ！
}

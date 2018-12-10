package main

import (
    "bytes"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    // "os"
    "math/rand"
    "time"
    "strconv"
)

func timeTrack(start time.Time, name string) {
    elapsed := time.Since(start)
    log.Printf("%s took %s", name, elapsed)
}

func main() {


    defer timeTrack(time.Now(), "main")

    // response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")
    rand.Seed(time.Now().UTC().UnixNano())

    var x int
    var ids [50]int
    for index := range ids {
        ids[index] = rand.Intn(1000)
    }

    fmt.Println("ids:",ids)

    url := "https://goapi.apps.pcfone.io/people/"
    // url := "http://localhost:8000/people/"
    fmt.Println("URL:>", url)

    var jsonStr = []byte(`{"firstname":"Mary","lastname":"Mo","address":{"city":"City Mo","state":"State Mo"}}`)

    for index := range ids {
//         ids[index] = rand.Intn(1000)
        req, err := http.NewRequest("POST", url+strconv.Itoa(ids[index]), bytes.NewBuffer(jsonStr))
        req.Header.Set("X-Custom-Header", "myvalue")
        req.Header.Set("Content-Type", "application/json")

        client := &http.Client{}
        resp, err := client.Do(req)
        if err != nil {
            panic(err)
        }
        defer resp.Body.Close()

        // fmt.Println("response Status:", resp.Status)
        // fmt.Println("response Headers:", resp.Header)
        body, _ := ioutil.ReadAll(resp.Body)
        x = x + len(body)
        // fmt.Println("response Body:", string(body))
        // fmt.Println(len(body))
    }


    fmt.Println("Total:",x)
    // response, err := http.Get(url)

    // if err != nil {
    //     fmt.Print(err.Error())
    //     os.Exit(1)
    // }
    
    // responseData, err := ioutil.ReadAll(response.Body)
    // if err != nil {
    //     log.Fatal(err)
    // }
    // fmt.Println(string(responseData))

}

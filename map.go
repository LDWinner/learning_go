package main

import "fmt"

func main() {
    var countryCapitalMap map[string]string
    countryCapitalMap = make(map[string]string)

    countryCapitalMap["France"] = "Paris"
    countryCapitalMap["Italy"] = "Roma"
    countryCapitalMap["Japan"] = "Tokyo"
    countryCapitalMap["India"] = "new"

    for country := range countryCapitalMap {
        fmt.Println(country,"captial is",countryCapitalMap[country])
    }

    captial,ok := countryCapitalMap["usa"]
    
    if(ok) {
        fmt.Println("The usa's captial is",captial)
    } else {
        fmt.Println("don't exist")
    }
}

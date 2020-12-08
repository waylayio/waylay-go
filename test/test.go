package main

import (
	"encoding/json"
	"fmt"
	"waylay"
)

type Data struct {
	Number int32 `json:"number"`
	Result string `json:"result"`
}

func main () {
	data := Data{
		Number: 4,
		Result: "nothing",
	}

	resp := waylay.PlugResponse{
		Message: "Here we go",
		ObservedState: "Done",
		RawData: data,
	}

	st, err := json.Marshal(resp)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(st))

	st2, err2 := json.Marshal(data)

	if err2 != nil {
		panic(err)
	}

	fmt.Println(string(st2))
}
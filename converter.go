package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// type Meta struct {
// 	country     string
// 	countryCode string
// 	currency    string
// 	code        string
// }

func main() {

	args := os.Args
	err := checkArgs(args)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	amount := args[1]
	from := args[2]
	to := args[3]

	resp, err := http.Get("https://www.google.com/finance/converter?a=" + amount + "&from=" + from + "&to=" + to)
	defer resp.Body.Close()
	body := io.Reader(resp.Body)
	doc, _ := goquery.NewDocumentFromReader(body)
	result := doc.Find("span.bld").Text()
	fmt.Println(result)
}

func checkArgs(args []string) error {
	args = strings.Fields(args[1])
	amount := args[0]

	if _, err := strconv.Atoi(amount); err != nil {
		return err
	}

	return nil
}

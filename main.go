package main

import (
	"golang.org/x/net/html"
	"strings"
	"fmt"
	"net/http"
	"log"
)

const breakpoint = "rnews:articleBody"

func main() {

	resp, err := http.Get("https://www.diolinux.com.br/2018/05/microsoft-bloqueia-atualizacao-do-windows-10-em-ssds-da-intel.html")

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body := html.NewTokenizer(resp.Body)

	//words := findWordsInArticle(body)
	words := findAllWords(body)
	words = cleanWords(words)
	
	fmt.Println(words)
}

func findWordsInArticle(body *html.Tokenizer) []string {
	var words = []string{}

	for {
		payload := body.Next()

		if payload == html.StartTagToken {
			content := body.Token()

			if strings.Contains(content.String(), breakpoint) {

				aux := body.Next()
				for !(aux == html.EndTagToken && strings.Contains(body.Token().String(), "</div>")) {
					//fmt.Print(aux) -> Mosta os Tokens : StartTag, Text, EndTag
					//fmt.Print(body.Token().String()) -> Mostra o Conteudo do Token

					if aux == html.TextToken {
						//fmt.Println(body.Token().String())

						temp := strings.Fields(body.Token().String())
						words = append(words, temp...)
					}
					aux = body.Next()
				}

				break
			}
		}
	}

	return words
}

func findAllWords(body *html.Tokenizer) []string {
	words := make([]string, 5)

	for !strings.Contains(body.Token().String(), "</html") {
		payload := body.Next()

		if strings.Contains(body.Token().String(), "<body") {
			for !strings.Contains(body.Token().String(), "</body") {
				payload = body.Next()

				if !strings.Contains(body.Token().String(), "<script") &&
						!strings.Contains(body.Token().String(), "<style") {
					payload = body.Next()
					if payload == html.TextToken {
						temp := strings.Fields(body.Token().String())
						words = append(words, temp...)

					}
				}
			}
			break
		}
	}

	return words
}

func countRepetitionWords(words []string) {
	repetition := map[string]int{}

	for _, word := range words {
		repetition[word] = repetition[word] + 1
	}

	fmt.Print(repetition)
}

func cleanWords(words []string) []string {

	for key, _ := range words {

		words[key] = strings.Replace(words[key], ".", "", -1)
		words[key] = strings.Replace(words[key], "\"", "", -1)
		words[key] = strings.Replace(words[key], "'", "", -1)
		words[key] = strings.Replace(words[key], "(", "", -1)
		words[key] = strings.Replace(words[key], ")", "", -1)

		if strings.Contains(words[key], "@*.com") {
			words[key] = ""
		}
	}

	return words
}

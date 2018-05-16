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

	var words = []string{}

	resp, err := http.Get("http://www.ifpb.edu.br/noticias/2018/05/campus-sousa-comunidade-escolhe-diretor-geral")

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body := html.NewTokenizer(resp.Body)

	words = findWordsInArticle(body)

	//fmt.Print(words)
	countRepetitionWords(words)

}

func findWordsInArticle(body *html.Tokenizer) []string  {
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

func findAllWords(body *html.Tokenizer)  {
	for {
		payload := body.Next()

		if payload == html.TextToken {

			fmt.Print(body.Token())

		}
	}
}

func countRepetitionWords(words []string) {
	repetition := map[string]int{}

	for _, word := range words {
		repetition[word] = repetition[word] + 1
	}

	fmt.Print(repetition)
}
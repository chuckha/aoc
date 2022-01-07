package main

import (
	"bufio"
	"bytes"
	"context"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/chuckha/aoc/app"
	"github.com/chuckha/aoc/internal/core/services/advent"
	"github.com/mitchellh/go-wordwrap"
	"jaytaylor.com/html2text"
)

func main() {
	year := flag.Int("year", time.Now().Year(), "the year")
	defaultDay := 1
	if time.Now().Month() == 12 {
		defaultDay = time.Now().Day()
	}
	day := flag.Int("day", defaultDay, "the day")
	level := flag.Int("level", 1, "the first or second level")
	desc := flag.Bool("desc", false, "get the description (always overrides)")
	submit := flag.Bool("submit", false, "submit an answer read from stdin")
	flag.Parse()

	ctx := context.Background()

	aoc, err := app.NewApplication()
	if err != nil {
		panic(fmt.Sprintf("%+v", err))
	}

	if *submit {
		scanner := bufio.NewScanner(os.Stdin)
		answer := ""
		for scanner.Scan() {
			answer = strings.TrimSpace(scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}

		resp, err := aoc.Service.SubmitAnswer(ctx, *year, *day, *level, answer)
		if err != nil {
			panic(fmt.Sprintf("%+v", err))
		}
		if resp == nil {
			fmt.Println("Correct!")
			os.Exit(0)
		}
		r, _ := html2text.FromString(string(resp))
		fmt.Println(wordwrap.WrapString(r, 80))
		os.Exit(1)
	}

	if *desc {
		desc, err := aoc.Service.GetDescription(ctx, &advent.GetDescriptionInput{Year: *year, Day: *day, Force: true})
		if err != nil {
			panic(fmt.Sprintf("%+v", err))
		}
		s, err := html2text.FromReader(bytes.NewReader(desc), html2text.Options{
			OmitLinks:    true,
			PrettyTables: true,
		})
		if err != nil {
			panic(err)
		}
		s = wordwrap.WrapString(s, 80)
		fmt.Println(s)
		os.Exit(0)
	}

	input, err := aoc.GetInput(ctx, *year, *day)
	if err != nil {
		panic(fmt.Sprintf("%+v", err))
	}
	fmt.Println(string(input))
}

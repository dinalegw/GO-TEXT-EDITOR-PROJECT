package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(string(s[0])) + strings.ToLower(s[1:])
}

func joinCommands(words []string) []string {
	var result []string
	for i := 0; i < len(words); i++ {

		if strings.HasPrefix(words[i], "(") && !strings.Contains(words[i], ")") {
			cmd := words[i]
			j := i + 1

			for j < len(words) && !strings.Contains(words[j], ")") {
				cmd += " " + words[j]
				j++
			}

			if j < len(words) {
				cmd += " " + words[j]
				i = j
			}

			result = append(result, cmd)
		} else {
			result = append(result, words[i])
		}
	}
	return result
}

func main() {

	var inF, outF string

	fmt.Print("Input file: ")
	fmt.Scan(&inF)

	fmt.Print("Output file: ")
	fmt.Scan(&outF)

	in, err := os.Open(inF)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer in.Close()

	out, err := os.Create(outF)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer out.Close()

	sc := bufio.NewScanner(in)
	wr := bufio.NewWriter(out)
	defer wr.Flush()

	for sc.Scan() {

		words := strings.Fields(sc.Text())

		words = joinCommands(words)

		for i := 0; i < len(words); i++ {

			if i > 0 && (words[i] == "(hex)" || words[i] == "(bin)") {

				base := 16
				if words[i] == "(bin)" {
					base = 2
				}

				v, err := strconv.ParseInt(words[i-1], base, 64)
				if err == nil {
					words[i-1] = strconv.FormatInt(v, 10)
				}

				words = append(words[:i], words[i+1:]...)
				i--
				continue
			}

			if strings.HasPrefix(words[i], "(") && strings.HasSuffix(words[i], ")") {

				cmd := strings.Trim(words[i], "()")

				mode := cmd
				n := 1

				if strings.Contains(cmd, ",") {

					parts := strings.Split(cmd, ",")

					mode = strings.TrimSpace(parts[0])

					num, err := strconv.Atoi(strings.TrimSpace(parts[1]))
					if err == nil {
						n = num
					}
				}

				for j := i - n; j < i; j++ {

					if j >= 0 {

						if mode == "up" {
							words[j] = strings.ToUpper(words[j])
						}

						if mode == "low" {
							words[j] = strings.ToLower(words[j])
						}

						if mode == "cap" {
							words[j] = capitalize(words[j])
						}
					}
				}

				words = append(words[:i], words[i+1:]...)
				i--
			}
		}

		// fix a -> an
		for i := 0; i < len(words)-1; i++ {

			next := strings.Trim(words[i+1], ".,!?:;'\"")

			if strings.ToLower(words[i]) == "a" &&
				len(next) > 0 &&
				strings.ContainsRune("aeiouh", unicode.ToLower(rune(next[0]))) {

				words[i] = "an"
			}
		}

		s := strings.Join(words, " ")

		// remove spaces before punctuation
		for _, p := range []string{".", ",", "!", "?", ":", ";"} {
			s = strings.ReplaceAll(s, " "+p, p)
		}

		// fix quotes spacing
		s = strings.ReplaceAll(s, "' ", "'")
		s = strings.ReplaceAll(s, " '", "'")

		// fix punctuation combos
		s = strings.ReplaceAll(s, " !!", "!!")
		s = strings.ReplaceAll(s, " !?", "!?")
		s = strings.ReplaceAll(s, " ...", "...")

		wr.WriteString(s + "\n")
	}
}

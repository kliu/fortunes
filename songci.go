package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

const titlePrefix  = "詞牌："
const authorPrefix  = "作者："
const contentPrefix  = "詞文："

func main() {
	if _, err := exec.Command("/bin/sh", "-c", "rm -f songci-fortunes*").CombinedOutput(); err != nil {
		log.Fatalf("rm run failed with %s\n", err)
	}

	src, err := os.Open("songci-src.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer src.Close()

	dst, err := os.Create("songci-fortunes")
	if err != nil {
		log.Fatal(err)
	}
	defer dst.Close()

	var title = ""
	var author = ""
	var content = ""

	scanner := bufio.NewScanner(src)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "□") {
			continue
		}
		if strings.HasPrefix(line, titlePrefix) {
			title = line[len(titlePrefix):]
		} else if strings.HasPrefix(line, authorPrefix) {
			author = line[len(authorPrefix):]
		} else if strings.HasPrefix(line, contentPrefix) {
			content = line[len(contentPrefix):]
			content = strings.Replace(content, "。 ", "。 \n", -1)
		}
		if title != "" && author != "" && content != "" {
			//fmt.Println(title + ": " + author + ": " + content)
			w := bufio.NewWriter(dst)
			fmt.Fprintln(w, content)
			fmt.Fprintln(w, "\n- " + author + " · " + title)
			fmt.Fprintln(w, "%")
			w.Flush()
			title = ""
			author = ""
			content = ""
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if _, err := exec.Command("strfile", "songci-fortunes", "songci-fortunes.dat").CombinedOutput(); err != nil {
		log.Fatalf("strfile run failed with %s\n", err)
	}

	if _, err := exec.Command("/bin/sh", "-c", "cp songci-fortunes* /usr/local/share/games/fortunes/").CombinedOutput(); err != nil {
		log.Fatalf("cp run failed with %s\n", err)
	}

	fmt.Println("Done.")
}

package main

import (
    "bufio"
    "bytes"
    "container/vector"
    "json"
    "os"
    "regexp"
    "strings"
    "web"
)

var words = map[string]int{}
var prefixes = map[string]int{}

type Search map[string]int

func (c Search) perms(prefix string, rest string) {
    if _, ok := prefixes[prefix]; !ok {
        return
    }

    if _, ok := c[prefix]; ok {
        return
    }

    if _, ok := words[prefix]; ok {
        c[prefix] = 1
    }

    for i := 0; i < len(rest); i++ {
        np := prefix + string(rest[i])
        nr := rest[0:i] + rest[i+1:]
        c.perms(np, nr)
    }
}

func search(ctx *web.Context) string {

    var search = Search(map[string]int{})
    s := ctx.Params["letters"]
    r := ""
    index := strings.Index(s, "[")
    var reg *regexp.Regexp
    var err os.Error
    if index > 0 {
        rindex := strings.Index(s, "]")
        if rindex > 0 {
            r = s[index+1 : len(s)-1]
            reg, err = regexp.Compile(r)
            if err != nil {
                println("Error creating regular expression", err.String())
            }
        }
        s = s[0:index]
    }

    search.perms("", s)

    var results vector.StringVector

    //filter results
    for k, _ := range search {
        if reg != nil && !reg.MatchString(k) {
            continue
        }
        results.Push(k)
    }

    if ctx.Headers.Get("X-Requested-With") == "XMLHttpRequest" {
        ctx.SetHeader("Content-Type", "application/json", true)
        js, _ := json.Marshal(results)
        return string(js)
    }

    var buf bytes.Buffer
    for _, k := range results {
        buf.WriteString(k)
        buf.WriteString("\n")
    }
    return buf.String()
}

func main() {
    f, err := os.Open("twl.txt", os.O_RDONLY, 0666)
    if err != nil {
        println(err.String())
    }
    reader := bufio.NewReader(f)
    if err != nil {
        println(err.String())
    }

    for {
        line, err := reader.ReadString('\n')
        if err != nil {
            if err != os.EOF {
                println(err.String())
            }
            break
        }
        word := strings.TrimSpace(line)
        words[word] = 1
        for i := 0; i <= len(word); i++ {
            prefixes[word[0:i]] = 1
        }
    }
    
    f.Close()
    web.Get("/search", search)
    web.Run("0.0.0.0:8080")
}

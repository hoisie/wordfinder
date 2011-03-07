package main

import (
    "bytes"
    "json"
    "os"
    "bufio"
    "strings"
    "web"
)

var words = map[string]int{}
var prefixes = map[string]int{}

type CachedSearch map[string]int

func (c CachedSearch) perms (prefix string, rest string) {
    if _,ok := prefixes[prefix]; !ok {
        return
    }
    
    if _,ok := c[prefix]; ok {
        return
    }
    
    if _,ok := words[prefix]; ok {
        c[prefix] = 1
    }
    
    if len(rest) >= 3 && rest[0] == '[' && rest[2] == ']' {
        c.perms(prefix+string(rest[1]), rest[2:])
        return
    }
    
    for i := 0; i < len(rest); i++ {
        if rest[i] == '[' {
            i+=2
            continue
        }
        np := prefix + string(rest[i])
        nr := rest[0:i]+rest[i+1:]
        c.perms(np, nr)
    }
}

func search(ctx *web.Context) string {

    var search = CachedSearch(map[string]int{})
    s := ctx.Params["letters"]
    
    search.perms("", s)
    
    if ctx.Headers.Get("X-Requested-With") == "XMLHttpRequest" {
        ctx.SetHeader("Content-Type", "application/json", true)
        js, _ := json.Marshal(search)
        return string(js)
    }
    
    var buf bytes.Buffer
    for k,_ := range search {
        buf.WriteString(k)
        buf.WriteString("\n")
    }    
    return buf.String()
}

func main() {    
    f,err := os.Open("twl.txt", os.O_RDONLY, 0666)
    if err != nil {
        println(err.String())
    }
    reader := bufio.NewReader(f)
    if err != nil {
        println(err.String())
    }
    
    for {
        line,err := reader.ReadString('\n')
        if err != nil {
            if err != os.EOF {
                println(err.String())
            }
            break
        }
        word := strings.TrimSpace(line)
        words[word] = 1
        for i := 0; i < len(word); i++ {
            prefixes[word[0:i]] = 1
        }
    }
    
    web.Get("/search", search)
    web.Run("0.0.0.0:8080")
    
}
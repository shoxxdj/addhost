package main

import (
    "log"
    "fmt"
    "os"
     "regexp"
    "github.com/ttacon/chalk"
)

func main() {
    if(len(os.Args)<3){
        fmt.Println(chalk.Red,"Syntax : addhost IP DOMAIN",chalk.Reset)
	os.Exit(1)
    }

    ip := os.Args[1]
    ipB := []byte(ip)
    domain := os.Args[2]

    var re = regexp.MustCompile(`(?m)^((25[0-5]|(2[0-4]|1\d|[1-9]|)\d)\.?\b){4}$`)


    if(re.Match(ipB)){
        f, err := os.OpenFile("/etc/hosts", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

        if err != nil {
            log.Fatal(err)
        }
        if _, err := f.Write([]byte(ip+" "+domain+"\n")); err != nil {
            log.Fatal(err)
        }
        if err := f.Close(); err != nil {
            log.Fatal(err)
        }
    }else{
        fmt.Println(chalk.Red,"Invalid IP",chalk.Reset)
    }


}

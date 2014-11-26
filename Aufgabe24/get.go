package main

import(
	"fmt"
	"os"
	"io/ioutil"
	"net/http"
)

func check(e error){
	if e != nil{
		panic(e)
	}
}

func main(){
	client := &http.Client{}
	if len(os.Args) < 2{
		fmt.Printf("usage: get <Portnumber>\n")
	}else{
		fmt.Printf("# Baue Verbindung Server auf\n")
		req, err := http.NewRequest("GET","http://localhost:"+string(os.Args[1]), nil)
		check(err)
		req.Header.Add("User-Agent", "My SysProg Client")
		fmt.Printf("# Connected \n# Schicke GET request an Webserver\n")
		resp, err := client.Do(req)
		check(err)
		res, err := ioutil.ReadAll(resp.Body)
		check(err)
		resp.Body.Close()
		err = ioutil.WriteFile("response.html", res, 0644)
		check(err)
		fmt.Printf("%s %s\nDate: %s\nServer: %s\nX-Powered-By: %s\nLast-Modified: %s\nExpires: %s\nETag: %s\nPragma: %s\nCache-Control: %s\nVary: %s\nContent-Type: %s\n", resp.Proto, resp.Status, resp.Header.Get("Date"), resp.Header.Get("Server"), resp.Header.Get("X-Powered-By"), resp.Header.Get("Last-Modified"), resp.Header.Get("Expires"), resp.Header.Get("Etag"), resp.Header.Get("Pragma"), resp.Header.Get("Cache-Control"), resp.Header.Get("Vary"), resp.Header.Get("Content-Type"))
	}
}
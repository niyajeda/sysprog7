package main

import(
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func check(e error){
	if e != nil{
		panic(e)
	}
}

func main(){
	client := &http.Client{}
	fmt.Printf("# Baue Verbindung zu Webserver net.cs.uni-bonn.de auf\n")
	req, err := http.NewRequest("GET", "http://net.cs.uni-bonn.de/de/wg/cs/lehre/ws-201415/sysprog/", nil)
	check(err)
	fmt.Printf("# Connected \n# Schicke GET request an Webserver\n");
	resp, err := client.Do(req)
	check(err)
	res, err := ioutil.ReadAll(resp.Body)
	check(err)
	resp.Body.Close()
	err = ioutil.WriteFile("zettel.html", res, 0644)
	check(err)
	fmt.Printf("%s %s\nDate: %s\nServer: %s\nX-Powered-By: %s\nLast-Modified: %s\nExpires: %s\nETag: %s\nPragma: %s\nCache-Control: %s\nVary: %s\nContent-Type: %s\n", resp.Proto, resp.Status, resp.Header.Get("Date"), resp.Header.Get("Server"), resp.Header.Get("X-Powered-By"), resp.Header.Get("Last-Modified"), resp.Header.Get("Expires"), resp.Header.Get("Etag"), resp.Header.Get("Pragma"), resp.Header.Get("Cache-Control"), resp.Header.Get("Vary"), resp.Header.Get("Content-Type"))
	//Durchsuche Quelltext nach Übungszetteln
	fmt.Printf("# Die folgenden Zettel sind aktuellen verfügbar:\n\n")
	s := string(res[:])
	var myRegExp = regexp.MustCompile("Übungsblatt\\s[0-9][0-9]")
	var blaetter []string = myRegExp.FindAllString(s, -1)
	for i:=0; i < len(blaetter); i++{
		fmt.Printf("%s\n", blaetter[i])	
	}
}
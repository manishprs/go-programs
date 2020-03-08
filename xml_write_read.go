package main
 
import (
	"encoding/xml"
	"io/ioutil"
	"fmt"
)
 
type contents struct {
	To      string `xml:"to"`
	From    string `xml:"from"`
	Heading string `xml:"heading"`
	Body    string `xml:"body"`
}

func writeToXml(){
	content := &contents{To: "Receiver",
		From:    "Sender",
		Heading: "Heading Section",
		Body:    "Body Section",
	}
 
	file, _ := xml.MarshalIndent(content, "", " ")
 
	_ = ioutil.WriteFile("myXML.xml", file, 0644)
}

func readFromXml(){
	data, _ := ioutil.ReadFile("myXML.xml")
 
	content := &contents{}
 
	_ = xml.Unmarshal([]byte(data), &content)
 
	fmt.Println(content.To)
	fmt.Println(content.From)
	fmt.Println(content.Heading)
	fmt.Println(content.Body)
}
 
func main() {
	
	writeToXml()
	readFromXml()
}
package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type Post2 struct {
	XMLName xml.Name `xml:"post"`
	Id string `xml:"id,attr"`
	Content string `xml:"content"`
	Author Author `xml:"author"`
}

type Author struct {
	Id string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

func main(){
	post := Post2 {
		Id: "1",
		Content: "hello lcp0578",
		Author: Author {
			Id: "2",
			Name: "go web",
		},
	}
	//output, err := xml.Marshal(&post)
	output, err := xml.MarshalIndent(&post, "", "\t")
	if err != nil {
		fmt.Println("error marshalling to XML:", err)
		return
	}
	//err = ioutil.WriteFile("post2.xml", output, 0644)
	err = ioutil.WriteFile("post2.xml", []byte(xml.Header + string(output)), 0644)
	if err != nil {
		fmt.Println("Error writing XML to file:", err)
		return
	}
}



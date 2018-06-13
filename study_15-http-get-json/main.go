package main

import (
	//"fmt"
	"net/http"
	//"io"
	//"os"
	//"net/url"
	//"net/url"
	//"io"
	//"os"
	"fmt"
	"encoding/json"
	//"io/ioutil"
	"io/ioutil"
)

func main() {


	//resp, err := http.Get("http://example.com/")
	resp, err := http.Get("https://free-api.heweather.com/v5/now?lang=en&city=yanta&key=9e721082765d49029271c2ec2e967c86")
	if err != nil {
		fmt.Println(err)
	}

	defer func() {
		println("done", resp.StatusCode)
		resp.Body.Close()

	}()


	//io.Copy(os.Stdout, resp.Body) //output Body

	body, err := ioutil.ReadAll(resp.Body) //Get the array of body

	fmt.Println(string(body)) //output as string


	fmt.Println("\n------------------\n")




	s_body := body[:]// assign to a slice

	fmt.Println(s_body)


	var sf interface{}
	err = json.Unmarshal(s_body, &sf)

	fmt.Println(err)

	fmt.Println(sf,"\r\n") //decode the json map

	fmt.Println("\n------------------\n")

	m := sf.(map[string]interface{})

	for k,v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "string", vv)
		case int:
			fmt.Println(k, "int", vv)
		case []interface{}:
			fmt.Println(k, "array")
			for i, u := range vv {
				fmt.Println(i,":", u,"\n")
			}
		}
	}



	fmt.Println("\ntest------------------\n")
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	var f interface{}
	err = json.Unmarshal(b, &f)

	fmt.Println(err, f)



	/*
	resp, err := http.PostForm("http://example.com/posts", url.Values{"title":{"article title"},"content":{"article body"}})
	defer func() {
		fmt.Println("Done", resp.StatusCode)
		resp.Body.Close()
	}()


	if err != nil {
		fmt.Println(err)
	}
	*/

	/*
	resp, err := http.Head("http://example.com")
	defer func() {
		println(resp.StatusCode)
		resp.Body.Close()
	}()

	if err != nil {
		fmt.Println(err)
	}

	io.Copy(os.Stdout, resp.Body) //Donot display anything due to no body will be replied

	println(resp.Header.Get("123"))
*/


}

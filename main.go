package main

import (
	"flag"
	"fmt"
	"github.com/shurcooL/graphql"
	//"github.com/urfave/cli"
	"io/ioutil"
	"net/http"
	"strings"
)

//import "github.com/machinebox/graphql"

//type ResponseStruct struct {
//	data: Json
//}
//}

var query struct {
	Me struct {
		Name graphql.String
	}
}

var q struct {
	Human struct {
		Name   graphql.String
		Height graphql.Float `graphql:"height(unit: METER)"`
	} `graphql:"human(id: \"1000\")"`
}

func queryServer(query string) []byte {
	url := "https://swapi-graphql.netlify.com/.netlify/functions/index"
	method := "POST"

	//q = query != nil ? query : "{\"query\":\"{\\n  allPeople {\\n    pageInfo {\\n      startCursor\\n      hasNextPage\\n      endCursor\\n    }\\n    people {\\n      name\\n    }\\n  }\\n}\",\"variables\":{}}"
	payload := strings.NewReader(query)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	return body
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	wordPtr := flag.String("query", "foo", "a string")
	queryFilePath := flag.String("query-file", "query.graphql", "filepath for your json")
	josnFilePath := flag.String("json-file", "data.json", "filepath for your json")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")
	flag.Parse()
	fmt.Println("word:", *wordPtr)
	fmt.Println("josnFilePath:", *josnFilePath)
	fmt.Println("queryFilePath:", *queryFilePath)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
	fmt.Println("hello")

	// READ FILES

	dat, err := ioutil.ReadFile(*queryFilePath)
	check(err)
	fmt.Print(string(dat))

	//app := &cli.App{
	//	Name: "greet",
	//	Usage: "say a greeting",
	//	Action: func(c *cli.Context) error {
	//		fmt.Println("Greetings")
	//		return nil
	//	},
	//}

	//app.Run(os.Args)

	fmt.Println(string(queryServer(string(dat))))

}

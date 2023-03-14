package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	_ "time"
)

type PreUrl struct {
	Url    string      `json:"datos"` //Defines that url field will be of the same type as the field "datos" (in returned json)
	Client http.Client //for testing purposes, we store here the type of call made
}

// Develops temporary main to test if data are correctly received
func main() {
	result, _ := GetPreUrl() //Invokes a fn to obtain the url to be used for the request of weather data
	fmt.Println(result)      //Prints result Ex. https://opendata.aemet.es/opendata/sh/4aa1d5d1
}

func GetPreUrl() (string, error) {
	//Defines variable url with the endpoint that must return the forecast of a particular town (here: Abrera).
	url := "https://opendata.aemet.es/opendata/api/prediccion/especifica/municipio/diaria/08001/?api_key=eyJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJjcmlzY29jb2Rlc0BnbWFpbC5jb20iLCJqdGkiOiJhOGEzODcwNS04NmM0LTQzOWItYmViZS0yYTEzYzQ1Njk5NjEiLCJpc3MiOiJBRU1FVCIsImlhdCI6MTY3ODcyNzU1MiwidXNlcklkIjoiYThhMzg3MDUtODZjNC00MzliLWJlYmUtMmExM2M0NTY5OTYxIiwicm9sZSI6IiJ9.4WOelPnL1dBSp-Js7mZBfIoNSA6I0U6hCZF3D_V8Dts"

	//Prepares GET request using http package
	req, _ := http.NewRequest("GET", url, nil) //Sets the GET method, the request's url, 3rd param is nil as we won't use it.

	//Adds control headers so that our request is not cached
	req.Header.Add("cache-control", "no-cache")

	//makes request using Do() method and pass variable req as a parameter -> contains request
	res, err := http.DefaultClient.Do(req)

	// controls if an error occurs
	if err != nil {
		log.Println("error contactant amb aemet.es", err)
		return "", err //returns error message
	}

	defer res.Body.Close()                //defers response
	body, err := ioutil.ReadAll(res.Body) //reads res.body
	if err != nil {
		log.Println("error decoding/reading json", err)
		return "", err //returns error message
	}

	preUrl := PreUrl{}                  //Creates empty struct
	err = json.Unmarshal(body, &preUrl) //stores body value in preUrl struct
	if err != nil {
		log.Println("error unmarshalling", err)
		return "", err
	}

	return preUrl.Url, err //returns data value

}

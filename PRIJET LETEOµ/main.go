package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

type People struct {
	Total int `json:"total_records"`
	Data  []struct {
		Uid  string `json:"uid"`
		Name string `json:"name"`
	} `json:"results"`
}

type Profil struct {
	Message string `json:"message"`
	Result  struct {
		Properties struct {
			Height    string `json:"height"`
			Mass      string `json:"mass"`
			HairColor string `json:"hair_color"`
			EyeColor  string `json:"eye_color"`
			Name      string `json:"name"`
			Homeworld string `json:"homeworld"`
		} `json:"properties"`
		Description string `json:"description"`
		ID          string `json:"_id"`
		UID         string `json:"uid"`
		V           int    `json:"__v"`
	} `json:"result"`
}

type Movies struct {
	Message string `json:"message"`
	Result  []struct {
		Properties struct {
			Characters   []string  `json:"characters"`
			Planets      []string  `json:"planets"`
			Starships    []string  `json:"starships"`
			Vehicles     []string  `json:"vehicles"`
			Species      []string  `json:"species"`
			Created      time.Time `json:"created"`
			Edited       time.Time `json:"edited"`
			Producer     string    `json:"producer"`
			Title        string    `json:"title"`
			EpisodeID    int       `json:"episode_id"`
			Director     string    `json:"director"`
			ReleaseDate  string    `json:"release_date"`
			OpeningCrawl string    `json:"opening_crawl"`
			URL          string    `json:"url"`
		} `json:"properties"`
		Description string `json:"description"`
		ID          string `json:"_id"`
		UID         string `json:"uid"`
		V           int    `json:"__v"`
	} `json:"result"`
}

type MovieDetail struct {
	Request struct {
		Type     string `json:"type"`
		Query    string `json:"query"`
		Language string `json:"language"`
		Unit     string `json:"unit"`
	} `json:"request"`
	Location struct {
		Name           string `json:"name"`
		Country        string `json:"country"`
		Region         string `json:"region"`
		Lat            string `json:"lat"`
		Lon            string `json:"lon"`
		TimezoneID     string `json:"timezone_id"`
		Localtime      string `json:"localtime"`
		LocaltimeEpoch int    `json:"localtime_epoch"`
		UtcOffset      string `json:"utc_offset"`
	} `json:"location"`
	Current struct {
		ObservationTime     string   `json:"observation_time"`
		Temperature         int      `json:"temperature"`
		WeatherCode         int      `json:"weather_code"`
		WeatherIcons        []string `json:"weather_icons"`
		WeatherDescriptions []string `json:"weather_descriptions"`
		WindSpeed           int      `json:"wind_speed"`
		WindDegree          int      `json:"wind_degree"`
		WindDir             string   `json:"wind_dir"`
		Pressure            int      `json:"pressure"`
		Precip              int      `json:"precip"`
		Humidity            int      `json:"humidity"`
		Cloudcover          int      `json:"cloudcover"`
		Feelslike           int      `json:"feelslike"`
		UvIndex             int      `json:"uv_index"`
		Visibility          int      `json:"visibility"`
		IsDay               string   `json:"is_day"`
	} `json:"current"`
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		url := "http://api.weatherstack.com/current?access_key=aaffa68db028d29f11cc04ca79a2b9f3&query=Canada"

		httpClient := http.Client{
			Timeout: time.Second * 8, // define timeout
		}

		//create request
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			log.Fatal(err)
		}

		//make api call
		res, getErr := httpClient.Do(req)
		if getErr != nil {
			log.Fatal(getErr)
		}

		req.Header.Set("User-Agent", "seb go tuto v2")

		if res.Body != nil {
			defer res.Body.Close()
		}

		//parse response
		body, readErr := ioutil.ReadAll(res.Body)
		if readErr != nil {
			log.Fatal(readErr)
		}

		response := People{}
		jsonErr := json.Unmarshal(body, &response)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.Execute(w, response)
	})

	http.HandleFunc("/movies", func(w http.ResponseWriter, r *http.Request) {

		url := "http://api.weatherstack.com/current?access_key=aaffa68db028d29f11cc04ca79a2b9f3&query=Canada"

		httpClient := http.Client{
			Timeout: time.Second * 6, // define timeout
		}

		//create request
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			log.Fatal(err)
		}

		//make api call
		res, getErr := httpClient.Do(req)
		if getErr != nil {
			log.Fatal(getErr)
		}

		req.Header.Set("User-Agent", "seb go tuto v2")

		if res.Body != nil {
			defer res.Body.Close()
		}

		//parse response
		body, readErr := ioutil.ReadAll(res.Body)
		if readErr != nil {
			log.Fatal(readErr)
		}

		response := Movies{}
		jsonErr := json.Unmarshal(body, &response)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}
		tmpl := template.Must(template.ParseFiles("movies.html"))
		tmpl.Execute(w, response)
	})

	http.HandleFunc("/profil/", func(w http.ResponseWriter, r *http.Request) {
		id := strings.ReplaceAll(r.URL.Path, "/query/", "")
		url := "http://api.weatherstack.com/current?access_key=aaffa68db028d29f11cc04ca79a2b9f3&query=" + id

		httpClient := http.Client{
			Timeout: time.Second * 6, // define timeout
		}

		//create request
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			log.Fatal(err)
		}

		//make api call
		res, getErr := httpClient.Do(req)
		if getErr != nil {
			log.Fatal(getErr)
		}

		req.Header.Set("User-Agent", "seb go tuto v2")

		if res.Body != nil {
			defer res.Body.Close()
		}

		//parse response
		body, readErr := ioutil.ReadAll(res.Body)
		if readErr != nil {
			log.Fatal(readErr)
		}

		response := Profil{}
		jsonErr := json.Unmarshal(body, &response)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}
		tmpl := template.Must(template.ParseFiles("profil.html"))
		tmpl.Execute(w, response)
	})

	http.HandleFunc("/query/", func(w http.ResponseWriter, r *http.Request) {
		id := strings.ReplaceAll(r.URL.Path, "/query/", "")
		url := "http://api.weatherstack.com/current?access_key=aaffa68db028d29f11cc04ca79a2b9f3&query=" + id

		httpClient := http.Client{
			Timeout: time.Second * 6, // define timeout
		}

		//create request
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			log.Fatal(err)
		}

		//make api call
		res, getErr := httpClient.Do(req)
		if getErr != nil {
			log.Fatal(getErr)
		}

		req.Header.Set("User-Agent", "seb go tuto v3")

		if res.Body != nil {
			defer res.Body.Close()
		}

		//parse response
		body, readErr := ioutil.ReadAll(res.Body)
		if readErr != nil {
			log.Fatal(readErr)
		}

		response := MovieDetail{}
		jsonErr := json.Unmarshal(body, &response)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}
		tmpl := template.Must(template.ParseFiles("movieDetail.html"))
		tmpl.Execute(w, response)
	})

	fs := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))

	img := http.FileServer(http.Dir("img"))
	http.Handle("/img/", http.StripPrefix("/img/", img))

	//serve http://localhost:8000
	http.ListenAndServe(":8000", nil)
}

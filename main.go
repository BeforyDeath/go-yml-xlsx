package main

import (
	"github.com/beforydeath/go-yml-xlsx/core"
	"net/http"
	"fmt"
	"log"
	"path/filepath"
	"io/ioutil"
	"html/template"
	"os"
)

func IndexHandle(w http.ResponseWriter, req *http.Request) {
	result := make(map[string]string)

	cwd, _ := os.Getwd()
	//cwd := "/home/byd/www/dev.beforydeath.ru/seo-yml"
	//fmt.Println(cwd)
	//fmt.Println(filepath.Join(cwd, "./template/default/index.html"))

	t, err := template.ParseFiles(filepath.Join(cwd, "./template/default/index.html"))
	if err != nil {
		fmt.Println(err)
	}
	url := req.URL.Query().Get("url");
	result["url"] = url
	if url == "" {
		t.ExecuteTemplate(w, "index", result)
		return
	}

	// http://dev.beforydeath.ru/yml/_export.yml

	response, err := http.Get(url)
	if err != nil {
		result["error"] = err.Error()
		t.ExecuteTemplate(w, "index", result)
		return
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		result["error"] = err.Error()
		t.ExecuteTemplate(w, "index", result)
		return
	}

	yml, err := core.GetYML(data)
	if err != nil {
		result["error"] = err.Error()
		t.ExecuteTemplate(w, "index", result)
		return
	}

	core.SetXlsx(yml)

	result["link"] = "tmp/xxxxxxxxxxx.xlsx"
	t.ExecuteTemplate(w, "index", result)

}

func main() {
	defer core.LogClose()
	core.Config.Get()

	http.Handle("/tmp/", http.StripPrefix("/tmp/", http.FileServer(http.Dir("./tmp/"))))
	http.HandleFunc("/", IndexHandle)

	fmt.Println("Server started ...")
	defer fmt.Println("Server stoped ...")
	log.Fatal(http.ListenAndServe(":8085", nil))
}
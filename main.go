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
	core.Config.Get()

	core.LogInf.Println("Handle:", req.RemoteAddr)

	result := make(map[string]string)

	cwd, _ := os.Getwd()
	if core.Config.BasePath != "" {
		cwd = core.Config.BasePath
	}

	t, err := template.ParseFiles(filepath.Join(cwd, "./template/default/index.html"))
	if err != nil {
		core.LogErr.Println(err.Error())
		return
	}
	url := req.URL.Query().Get("url");
	result["url"] = url
	if url == "" {
		t.ExecuteTemplate(w, "index", result)
		return
	}
	core.LogInf.Println("Get file:", url)

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

	name, err := core.SetXlsx(yml, url)
	if err != nil {
		core.LogErr.Println(err)
	}
	result["link"] = name
	t.ExecuteTemplate(w, "index", result)
	core.LogInf.Println("File result:", name)

}

func main() {

	http.Handle("/tmp/", http.StripPrefix("/tmp/", http.FileServer(http.Dir("./tmp/"))))
	http.HandleFunc("/", IndexHandle)

	stoped := func() {
		core.LogClose()
		fmt.Println("Server stoped ...")
		core.LogInf.Println("Server stoped ...")
	}

	defer stoped()

	core.LogInf.Println("Server started ...")
	fmt.Println("Server started ...")
	log.Fatal(http.ListenAndServe(":8085", nil))
}
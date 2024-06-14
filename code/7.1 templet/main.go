package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type stu struct {
	Name string
	Age  int16
	Sex  bool
}

func AddAge(age interface{}) int16 {
	var add int16
	if age, ok := age.(int16); ok {
		add = age + 1
	}
	return add
}

func pading(w http.ResponseWriter, r *http.Request) {
	// 1.定义模板 Details see dir templet
	// 2.解析模板
	t, err := template.ParseFiles("./templet/base.html")
	if err != nil {
		log.Fatal("文件解析错误: ", err)
	}
	// 3.渲染模板
	student := stu{
		Name: "zhu",
		Age:  26,
		Sex:  true,
	}
	sli := []string{"篮球", "足球", "跑步"}
	temp := map[string]interface{}{
		"stu":   student,
		"slice": sli,
	}
	t.Execute(w, temp)
}
func ageAdd(w http.ResponseWriter, r *http.Request) {
	// 1.定义模板 Details see dir templet
	htmlByte, err := os.ReadFile("./templet/ageAdd.html")
	if err != nil {
		fmt.Println("read html failed, err:", err)
		return
	}
	// 2.解析模板

	t, err := template.New("AgeAdd").Funcs(template.FuncMap{"AgeAdd": AddAge}).Parse(string(htmlByte))
	if err != nil {
		log.Fatal("解析模板错误: ", err)
	}

	// 渲染模板
	student := stu{
		Name: "zhu",
		Age:  26,
		Sex:  true,
	}
	t.Execute(w, student)
}
func main() {
	http.HandleFunc("/base", pading)
	http.HandleFunc("/age", ageAdd)

	err := http.ListenAndServe(":9000", nil)
	fmt.Println("Server Start.")
	if err != nil {
		log.Fatal("Fail to start service: ", err)
	}
}

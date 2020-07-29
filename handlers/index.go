package handlers

import (
	"github.com/jaykof/chitchat/models"
	"net/http"
)

// 论坛首页路由处理器方法
func Index(w http.ResponseWriter, r *http.Request) {
	//files := []string{"D:\\study\\go\\src\\github.com\\jaykof\\chitchat\\views/layout.html", "D:\\study\\go\\src\\github.com\\jaykof\\chitchat\\views/navbar.html", "D:\\study\\go\\src\\github.com\\jaykof\\chitchat\\views/index.html"}
	// 注意以下模板文件的路径，从不同的路径启动，相对路径也不一样，下面的一定要从当前main.go的目录下运行程序，否则模板路径会有问题
	//files := []string{"views/layout.html", "views/navbar.html", "views/index.html"}
	//templates := template.Must(template.ParseFiles(files...))
	////pathstr, _ := os.Getwd()
	////fmt.Println(pathstr)
	//threads, err := models.Threads()
	//if err == nil {
	//	templates.ExecuteTemplate(w, "layout", threads)
	//}

	threads, err := models.Threads()
	if err == nil {
		_, err := session(w, r)
		if err != nil {
			generateHTML(w, threads, "layout", "navbar", "index")
		} else {
			generateHTML(w, threads, "layout", "auth.navbar", "index")
		}
	}
}

func Err(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	_, err := session(writer, request)
	if err != nil {
		generateHTML(writer, vals.Get("msg"), "layout", "navbar", "error")
	} else {
		generateHTML(writer, vals.Get("msg"), "layout", "auth.navbar", "error")
	}
}

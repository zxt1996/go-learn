package myHttp

import (
	"log"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request)  {
	if r.URL.Path != "/" {
		//就调用http包中的方法NotFound，并将响应和请求传递给它。这将向客户端返回一个404响应。
		http.NotFound(w, r)
	} else {
		//通过设置Content-Type报头，服务器可告诉客户端，发送的是JSON数据
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		w.Write([]byte(`{"hello": "world"}`))
	}
}

func MyHttp()  {
	//HandleFunc用于注册对URL地址映射进行响应的函数
	http.HandleFunc("/", helloWorld)
	err := http.ListenAndServe(":9093", nil)

	if err != nil {
		log.Fatal(err)
	}
}

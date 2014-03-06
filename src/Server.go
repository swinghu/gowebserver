
package main
import (
    "fmt"
    "html/template"
    "log"
    "net/http"
    "strings"
    "time"
    "crypto/md5"
    "io"
    "strconv"
    "os"
)
func sayhelloName(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()       //解析url传递的参数，对于POST则解析响应包的主体（request body）
    //注意:如果没有调用ParseForm方法，下面无法获取表单的数据
    fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
    fmt.Println("path", r.URL.Path)
    fmt.Println("scheme", r.URL.Scheme)
    fmt.Println(r.Form["url_long"])
    for k, v := range r.Form {
        fmt.Println("key:", k)
        fmt.Println("val:", strings.Join(v, ""))
    }
    fmt.Fprintf(w, "Hello Golang User !") //这个写入到w的是输出到客户端的
}

func login(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method) //获取请求的方法
    if r.Method == "GET" {
    	crutime := time.Now().Unix()
    	h := md5.New()
    	io.WriteString(h,strconv.FormatInt(crutime,10))
    	token := fmt.Sprintf("%x",h.Sum(nil))
    	
        t, _ := template.ParseFiles("login.gtpl")
        w.Header().Set("Content-Type", "text/html; charset=utf-8")// fix the the return html file
        t.Execute(w, token)//增加 token：将nil 改为token
        
   		} else {
        //请求的是登陆数据，那么执行登陆的逻辑判断
        r.ParseForm()//默认handler 不对form进行解析   
        token := r.Form.Get("token")
        if token != "" {
        	fmt.Println("token is :", template.HTMLEscapeString(r.Form.Get("token")))
        } else {
        	fmt.Println("token  errer")
        }
        
        
        fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username")))
        fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
        fmt.Println("fruit:",template.HTMLEscapeString(r.Form.Get("fruit")))
        
        if len(r.Form.Get("password")) == 0 {
        	//template.HTMLEscape(w,[]byte("password is null"))
        	t, _ := template.ParseFiles("login.gtpl")
       		w.Header().Set("Content-Type", "text/html; charset=utf-8")
       		t.Execute(w, nil)
        }else {
        	t, _ := template.ParseFiles("upload.gtpl")
        	w.Header().Set("Content-Type", "text/html; charset=utf-8")
       		t.Execute(w, nil)
        	template.HTMLEscape(w,[]byte("Hello "+r.Form.Get("username")+",you have loginning,the web token is "+template.HTMLEscapeString(r.Form.Get("token"))))
        }
    }
}


func upload (w http.ResponseWriter, r *http.Request) {//上传文件处理
	fmt.Println("method",r.Method)
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h,strconv.FormatInt(crutime,10))
		token := fmt.Sprintf("%x",h.Sum(nil))
		t, _ := template.ParseFiles("upload.gtpl")
		t.Execute(w,token)
	}else {
		r.ParseMultipartForm(32<<20)
		file, handler ,err := r.FormFile("uploadfile")		
		if err != nil {
			fmt.Println(err)
			return
		}
		
		defer file.Close()
		fmt.Fprintf(w,"%v",handler.Header)
		
		f, err := os.OpenFile("./test/"+handler.Filename,os.O_WRONLY|os.O_CREATE,0666)
		if err != nil {
			fmt.Println(err)
			return
		} else {
			fmt.Println("File:"+handler.Filename+" upload successfully")
		}
		defer f.Close()
		io.Copy(f,file)
	}
}

func main() {
    http.HandleFunc("/", sayhelloName)       //设置访问的路由
    http.HandleFunc("/login", login)         //设置访问的路由
    http.HandleFunc("/upload",upload)
    err := http.ListenAndServe(":9090", nil) //设置监听的端口
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
    
    
}
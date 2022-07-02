//package main
//
//import (
//	"fmt"
//	"github.com/linvon/cuckoo-filter"
//)
//
//var strs = []string{"abc", "bcd", "cde", "abc"}
//
//func main() {
//	cf := cuckoo.NewFilter(4, 9, 3900, cuckoo.TableTypePacked)
//	fmt.Println(cf.Info())
//	fmt.Println(cf.FalsePositiveRate())
//
//	for _, str := range strs {
//
//		a := []byte(str)
//		//cf.Add(a) //添加
//		b, _ := cf.Encode()
//		ncf, _ := cuckoo.Decode(b)
//		contain := ncf.Contain(a)
//
//		if !contain {
//			cf.Add(a)
//		}
//		//cf.Delete(a)
//		//fmt.Println(cf.Size())
//
//	}
//	//a := []byte("A")
//	//cf.Add(a)
//	//fmt.Println(cf.Contain(a))
//	fmt.Println(cf.Size())
//
//	//b, _ := cf.Encode()
//	//ncf, _ := cuckoo.Decode(b)
//	//fmt.Println(ncf.Contain(a))
//
//	//cf.Delete(a)
//	//fmt.Println(cf.Size())
//}

package main

import (
	"fmt"
	"github.com/markbates/pkger"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(pkger.Dir("/fe/fe-app/dist/"))
	mux.Handle("/", fs)

	srv := &http.Server{
		Addr:    "127.0.0.1:4022",
		Handler: mux,
	}

	fmt.Println("serving at", srv.Addr)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

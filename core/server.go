package core

import "net/http"
import "fmt"
import "strconv"

func RunServer(brd *Board) {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		q := req.URL.Query()
		if len(q["r"]) == 0 || len(q["g"]) == 0 || len(q["b"]) == 0 || len(q["x"]) == 0 || len(q["y"]) == 0 {
			fmt.Fprint(w, "Missing values")
			return
		}
		r, _ := strconv.ParseInt(q["r"][0], 10, 0)
		g, _ := strconv.ParseInt(q["g"][0], 10, 0)
		b, _ := strconv.ParseInt(q["b"][0], 10, 0)
		x, _ := strconv.ParseInt(q["x"][0], 10, 0)
		y, _ := strconv.ParseInt(q["y"][0], 10, 0)
		brd.DrawPixel(int(x), int(y), MakeColor(int(r), int(g), int(b)))
		brd.Save()
	})
	http.HandleFunc("/clear", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		brd.DrawAll(MakeColor(0, 0, 0))
		brd.Save()
	})
	http.ListenAndServe(":8070", nil)
}

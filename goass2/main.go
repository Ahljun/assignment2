package mycalc

import (
        "html/template"
        "net/http"
        "strconv"

        "appengine"
        "appengine/datastore"
)

type RESULT struct {
        X           int64
        Y           int64		
        Sum     	int64
		Diff        int64
		Ave         float64
}

func init() {
        http.HandleFunc("/", root)
        http.HandleFunc("/sign", sign)
}


func resultKey(c appengine.Context) *datastore.Key {
        return datastore.NewKey(c, "MyResult", "my_result", 0, nil)
}

func root(w http.ResponseWriter, r *http.Request) {
        c := appengine.NewContext(r)
        q := datastore.NewQuery("RESULT").Ancestor(resultKey(c)).Limit(10)
        results := make([]RESULT, 0, 10)
        if _, err := q.GetAll(c, &results); err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
        }

        if err := mainpageTemplate.Execute(w, results); err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
        }
}

var mainpageTemplate = template.Must(template.New("mainpage").Parse(`
<html>
  <head>
    <title>Assignment 1</title>
    <style>
        table, th, td {
            border: 1px solid black;
            border-collapse: collapse;
        }
    </style>
  </head>
  <body>
    <table style="width:300px;margin:10px;">
    <tr>
        <th>X</th>
        <th>Y</th> 		
        <th>Sum</th>
        <th>Diff</th>
        <th>Ave</th>        
    </tr>
    {{range .}}
    <tr>
        <td>{{.X}}</td>
        <td>{{.Y}}</td> 
		<td>{{.Sum}}</td>
		<td>{{.Diff}}</td>
        <td>{{.Ave | printf "%.2f"}}</td>
    </tr>
    {{end}}
    </table>
    <form action="/sign" method="post">
      <div>X<input name="Xval" type="number"></input></div>
      <div>Y<input name="Yval" type="number"></input></div>
      <div><input type="submit" value="Submit"></div>
    </form>
  </body>
</html>
`))

func sign(w http.ResponseWriter, r *http.Request) {
        c := appengine.NewContext(r)
        Xcon, _ := strconv.ParseInt(r.FormValue("Xval"), 10, 64)
        Ycon, _ := strconv.ParseInt(r.FormValue("Yval"), 10, 64)		
        sum := Xcon + Ycon
        diff := Xcon - Ycon
        ave := float64(Xcon) + float64(Ycon) / 2		
        g := RESULT{
                X:          Xcon,
                Y:          Ycon,                
                Sum:        sum,
                Diff:       diff,
                Ave:        ave,         
        }
        key := datastore.NewIncompleteKey(c, "RESULT", resultKey(c))
        _, err := datastore.Put(c, key, &g)
        if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
        }
        http.Redirect(w, r, "/", http.StatusFound)
}
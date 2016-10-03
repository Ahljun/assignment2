package mycalc

import (
        "html/template"
        "net/http"
        "strconv"
        "time"
        "appengine"
        "appengine/datastore"
)

type RESULT struct {
		Height      int64
		Weight      int64
		BMI         float64
		Date        time.Time
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
	    <th>Date</h>
        <th>Weight</th>
		<th>Height</th>
		<th>BMI</th>
    </tr>
    {{range .}}
    <tr>
	    <td>{{.Date}}</td>
		<td>{{.Weight}}</td>
		<td>{{.Height}}</td>
		<td>{{.BMI | printf "%.2f"}}</td>
    </tr>
    {{end}}
    </table>
    <form action="/sign" method="post">
	  <div>Height: <input name="Height" type="number"></input></div>
	  <div>Weight: <input name="Weight" type="number"></input></div>
      <div><input type="submit" value="Submit"></div>
    </form>
  </body>
</html>
`))

func sign(w http.ResponseWriter, r *http.Request) {
        c := appengine.NewContext(r)
		Height, _ := strconv.ParseInt(r.FormValue("Height"), 10, 64)
		Weight, _ := strconv.ParseInt(r.FormValue("Weight"), 10, 64)
		BMI := (float64(Weight) / (float64(Height) * float64(Height))) * float64(10000) 
		
        g := RESULT{
                Date:       time.Now(),		
				Weight:     Weight,
				Height:     Height,
				BMI:        BMI,
        }
        key := datastore.NewIncompleteKey(c, "RESULT", resultKey(c))
        _, err := datastore.Put(c, key, &g)
        if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
        }
        http.Redirect(w, r, "/", http.StatusFound)
}
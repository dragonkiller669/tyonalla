package main
import "fmt"
import "time"
import "sort"


const (
	RFC3339 = "2006-01-02T15:04:05Z07:00"
)

func main() {
	t := time.Now()
//	t0, e := time.Parse(
//		RFC3339,
//		"2006-01-02T15:04:05Z")
	t1 := t.Format("24 Feb")
	t2 := t.Format("2 Jan")
	t3 := t.Format("2 Dec")
	t4 := t.Format("28 Jan")
	t5 := t.Format("26 Feb")
//	fmt.Println(t0,e)

        tavarat := []struct {
                nimi string
                pvm string
        }{
                {"kalja", t5},
                {"maito", t4},
		{"kananmuna", t3},
		{"juusto", t2},
		{"tomaatti", t1},

        }
//	_, e = time.Parse(time.ANSIC, "8:41PM")

sort.Slice(tavarat, func(i, j int) bool { 
	return tavarat[i].pvm < tavarat[j].pvm })

//	o := sort.Sort(len(tavarat.pvm))
//	fmt.Println(o)
        for _, tavara := range tavarat {
                fmt.Println(tavara)
	}
}




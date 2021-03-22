package cerebro

import (
	"github.com/Sebalvarez97/mutants-go/tools/matrix"
	"regexp"
	"sync"
)

var (
	wg    sync.WaitGroup
	mutex sync.Mutex
	re    = regexp.MustCompile(`(A{4}|T{4}|G{4}|C{4})`)
)

func IsMutant(dna [][]byte) (bool, int) {
	var found int

	t := matrix.Transpose(dna)
	d := matrix.Diagonals(dna)
	dt := matrix.Diagonals(t)

	r := matrix.Reverse(dna)
	dr := matrix.Diagonals(r)
	tr := matrix.Transpose(r)
	dtr := matrix.Diagonals(tr)

	for _, row := range t {
		dna = append(dna, row)
	}
	for _, row := range d {
		dna = append(dna, row)
	}
	for i := 1; i < len(dt); i++ {
		dna = append(dna, dt[i])
	}
	for _, row := range dr {
		dna = append(dna, row)
	}
	for i := 1; i < len(dtr); i++ {
		dna = append(dna, dtr[i])
	}

	rows := make(chan []byte)

	go func() {
		for _, row := range dna {
			rows <- row
		}
		close(rows)
	}()

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for row := range rows {
				mutex.Lock()
				if found <= 1 {
					if ir := re.FindAll(row, -1); ir != nil {
						found = found + len(ir)
					}
				}
				mutex.Unlock()
			}
		}()
	}
	wg.Wait()
	return found > 1, found
}

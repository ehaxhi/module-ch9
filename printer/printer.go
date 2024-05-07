package printer

import (
	"fmt"
	"github.com/ehaxhi/module-ch9/models"
	"os"
	"text/tabwriter"
)

func Format(num int) string {
	return fmt.Sprintf("The number is: %d", num)
}

type Printer struct {
	w *tabwriter.Writer
}

func New() *Printer {
	w := tabwriter.NewWriter(os.Stdout, 3, 0, 3, ' ', tabwriter.TabIndent)

	return &Printer{w: w}
}

func (p *Printer) CityHeader() {
	fmt.Fprintln(p.w, "Name\tTempC\tTempF\tBeach vacation ready?\tSki vacation ready?")
}

func (p *Printer) CityDetails(c models.CityTemp, q models.CityQuery) {
	fmt.Fprintf(p.w, "%v\t%v\t%v\t%v\t%v\n", c.Name(), c.TempC(q), c.TempF(q), c.BeachVacationReady(q), c.SkiVacationReady(q))
}

func (p *Printer) Cleanup() {
	p.w.Flush()
}

package window

import (
	"fmt"
	"olx/data"
	"olx/register"
	"os"
	"text/tabwriter"
)

func App(id string) {
	choice := ""
	for {
		var products, err = data.Get(id)
		if err != nil {
			fmt.Print(err)
			return
		}
		Writer(products, id, "Gold")
		Writer(products, id, "Silver")
		Writer(products, id, "Bronze")
		fmt.Print(`
new-n  ||  buy-b  ||  signup-s  ||  login-l
Choose >>> `)
		if _, err := fmt.Scan(&choice); err != nil {
			return
		}
		switch choice {
		case "n":
			var id = register.New()
			products, err = data.Get(id)
			if err != nil {
				fmt.Print(err)
				return
			}
		case "b":
		case "s":
		case "l":

		}
	}
}

func Writer(products []data.Product, id string, ty string) {
	var w = tabwriter.NewWriter(os.Stdout, 5, 1, 1, ' ', tabwriter.AlignRight|tabwriter.Debug)
	fmt.Println("\n", ty, "▼")
	fmt.Fprintf(w, "\tSeller\tName\tPrice\t\n")
	for i := range products {
		if id == products[i].TraderId && products[i].Type == ty {
			fmt.Fprintf(w, "\t%s\t%s\t%d\t\n", products[i].TraderName, products[i].Name, products[i].Price)
		}
	}
	for i := range products {
		if id != products[i].TraderId && products[i].Type == ty {
			fmt.Fprintf(w, "\t%s\t%s\t%d\t\n", products[i].TraderName, products[i].Name, products[i].Price)
		}
	}
	w.Flush()
}

# eztable

## What is eztable?

**eztable** is a simple, but effective package for creating tables inside of your terminal. It allows for both unicode characters such as emojis inside of the table contents, and escape codes too.

```
╔═════╤══════════════╤═════════════════════════════════════╗
║ #   │         Name │             Description             ║
╟━━━━━┼━━━━━━━━━━━━━━┼━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━╢
║ 1   │    hello     │ Contain items inside of your table! ║
║  50 │ Goodbye mate │  This is an 🌐 example! 😲😲😲😲  ║
║ 240 │        Woosh │             Align with padding too. ║
╚═════╧══════════════╧═════════════════════════════════════╝
```

## Key features

- **Custom Styling** for tables, you can use the pre-made styles for the tables OR you can create a custom style yourself to suit your needs.

- **Escape code support + unicode support** meaning that your table will not be manipulated by including unicode characters, emojis or ANSI escape codes!

- **Padding** space each cell of your table how you want, you can use; Centre, Left or Right spacing to ensure that you have your table in a layout that you desire.

- **Effective speeds** meaning that you will not be waiting for your table to generate for a long time, it is generated in literal milliseconds, perfectly everytime.

## Example

```go
package main

import (
	"eztable"
	"fmt"
	"strconv"
	"testing"
)

type Value struct {
	Name, Description string
	Ammount           int
}

var Values []*Value = []*Value{
	{
		Name:        "item1",
		Description: "This is item1",
		Ammount:     7,
	},
	{
		Name:        "item2",
		Description: "This is item2, watch how the table evolves.",
		Ammount:     100,
	},
	{
		Name:        "item3",
		Description: "n/a",
		Ammount:     0,
	},
}

func main() {
	var Cells = []*eztable.Cell{
		{
			Title: "Name",
			Alignment: &eztable.Align{
				Header: eztable.Centre,
				Body:   eztable.Centre,
			},
			Bodys: make([]string, 0),
		},
		{
			Title: "Description",
			Alignment: &eztable.Align{
				Header: eztable.Left,
				Body:   eztable.Centre,
			},
			Bodys: make([]string, 0),
		},
		{
			Title: "Amount",
			Alignment: &eztable.Align{
				Header: eztable.Right,
				Body:   eztable.Right,
			},
			Bodys: make([]string, 0),
		},
	}

	for _, v := range Values {
		Cells[0].Bodys = append(Cells[0].Bodys, v.Name)
		Cells[1].Bodys = append(Cells[1].Bodys, v.Description)
		Cells[2].Bodys = append(Cells[2].Bodys, strconv.Itoa(v.Ammount))
	}

	table := eztable.NewTable()
	table.Cells = Cells
	
	if err := table.SetStyle(eztable.Unicode); err != nil { panic(err) }
	fmt.Println(table.String())
}
```

Result:
```
╔═══════╤═════════════════════════════════════════════╤════════╗
║ Name  │ Description                                 │ Amount ║
╟━━━━━━━┼━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┼━━━━━━━━╢
║ item1 │                This is item1                │      7 ║
║ item2 │ This is item2, watch how the table evolves. │    100 ║
║ item3 │                     n/a                     │      0 ║
╚═══════╧═════════════════════════════════════════════╧════════╝
```
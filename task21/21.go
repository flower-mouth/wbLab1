//Реализовать паттерн «адаптер» на любом примере.

package task21

import "fmt"

type OldPrinter interface {
	Print(s string) string
}

type MyOldPrinter struct{}

func (p *MyOldPrinter) Print(s string) string {
	return fmt.Sprintf("Старая печать: %s", s)
}

type NewPrinter interface {
	PrintStored() string
}

type MyNewPrinter struct {
	Msg string
}

func (p *MyNewPrinter) PrintStored() string {
	return fmt.Sprintf("Новая печать: %s", p.Msg)
}

type PrinterAdapter struct {
	OldPrinter OldPrinter
	Msg        string
}

func (p *PrinterAdapter) PrintStored() string {
	if p.OldPrinter != nil {
		newMsg := fmt.Sprintf("Адаптер: %s", p.Msg)
		return p.OldPrinter.Print(newMsg)
	}
	return p.Msg
}

func Task21() {
	newPrinter := &MyNewPrinter{"Привет, мир!"}
	fmt.Println(newPrinter.PrintStored())

	oldPrinter := &MyOldPrinter{}
	adapter := &PrinterAdapter{
		OldPrinter: oldPrinter,
		Msg:        "Привет, мир!",
	}
	fmt.Println(adapter.PrintStored())
}

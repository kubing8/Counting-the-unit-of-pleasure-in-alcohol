package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type AlcName string

const unknownAlc AlcName = "Неизвестно"

type alcohol struct {
	name AlcName
	deg  float32
}

var alcohols = []alcohol{
	{
		name: "Пиво",
		deg:  4.5,
	},
	{
		name: "Игристое вино",
		deg:  8.0,
	},
	{
		name: "Вино",
		deg:  12.0,
	},
	{
		name: "Водка",
		deg:  40.0,
	},
}

func main() {
	fmt.Println("Программа для нахождения кайфовного кол-ва миллилитров алкоголя")
	fmt.Print("Введите кол-во миллилитров: ")
	var vol int
	if _, err := fmt.Scan(&vol); err != nil {
		panic(fmt.Sprintf("Error volume scan: %v", err))
	}

	fmt.Print("Введите кол-во процентов: ")
	var al float32
	if _, err := fmt.Scan(&al); err != nil {
		panic(fmt.Sprintf("Error percent scan: %v", err))
	}

	outputData(calc(vol, al))

	_, _ = bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func outputData(clearAlk float32, name AlcName) {
	for _, b := range alcohols {
		if name == b.name {
			continue
		}
		ml := ((100 - b.deg) * clearAlk / b.deg) + clearAlk
		fmt.Printf("Для %q ваша единица кайфа: %v мл, %2.1f%%\n", b.name, ml, b.deg)
	}

}

func calc(vol int, al float32) (float32, AlcName) {
	clearAlk := float32(vol) * al / 100.0

	name := getAlcName(al)

	fmt.Printf("Скорее всего Вы пьете: %q, %v мл %2.1f%%\n\n", name, vol, al)
	return clearAlk, name
}

func getAlcName(al float32) AlcName {
	sort.Slice(alcohols, func(i, j int) bool {
		return alcohols[i].deg < alcohols[j].deg
	})

	res := unknownAlc // "Неизвестно"
	for _, a := range alcohols {
		if al < a.deg {
			break
		}
		res = a.name
	}
	return res
}

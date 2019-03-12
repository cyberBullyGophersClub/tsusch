package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	getSchedule()
}

func getSchedule() {
	doc, err := goquery.NewDocument("http://schedule.tsu.tula.ru/?group=620251-%D0%9F%D0%91")
	if err != nil {
		fmt.Println(err.Error())
	}

	allSchedule := doc.Find("#results > table:nth-child(3) > tbody").Contents().Text()
	response := strings.Join(strings.Fields(allSchedule), " ")
	ef, err := os.Create("schedule.txt")
	if err != nil {
		fmt.Println("Problems with allSchedule", err.Error())
	}
	ef.WriteString(response)

	weekday := time.Now().Weekday()
	day := (int(weekday))
	fmt.Println(day)

	switch day {
	case 1:
		getDay("Понедельник", "Вторник")
	case 2:
		getDay("Вторник", "Среда")

	case 3:
		getDay("Среда", "Четверг")
	}
}

func getDay(day1, day2 string) {
	txt, err := ioutil.ReadFile("schedule.txt")
	if err != nil {
		fmt.Println(err.Error())
	}
	i := strings.Index(string(txt), day1)
	i2 := strings.Index(string(txt), day2)
	part := (txt[i:i2])
	text := (string(part))
	fmt.Print(text)
}

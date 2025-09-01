package main

import (
	"fmt"
	"time"
)

func main() {
	var startTime, name, surname, middleName  string 
	var sumPayments1, sumPayments2, sumPayments3 float64
	
	fmt.Scanln(&startTime)
	fmt.Scanln(&name)
	fmt.Scanln(&surname)
	fmt.Scanln(&middleName)
	fmt.Scanln(&sumPayments1)
	fmt.Scanln(&sumPayments2)
	fmt.Scanln(&sumPayments3)
	
	dateSupscribe, err := time.Parse("02.01.2006", startTime)

	if err != nil {
		fmt.Println("eror: ", err)
	}

	dateSupscribeResult := dateSupscribe.Add(15 * 24 * time.Hour)

	fullPayments := sumPayments1 + sumPayments2 + sumPayments3

	n := int(fullPayments)
	kopeyki := fullPayments - float64(n)
	resultKopeek := int(kopeyki * 100)

	fmt.Printf(`Уважаемый, %s %s %s, доводим до вашего сведения, что бухгалтерия сформировала документы по факту выполненной вами работы. 
Дата подписания договора: %s. Просим вас подойти в офис в любое удобное для вас время в этот день. 
Общая сумма выплат составит %d руб. %d коп.

С уважением, 
Гл. бух. Иванов А.Е.
`, surname, name, middleName, dateSupscribeResult.Format("02.01.2006"), n, resultKopeek)
}

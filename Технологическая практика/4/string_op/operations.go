package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	str := `Our Father, who art in heaven, hallowed be thy name; thy kingdom come; ` +
		`thy will be done; on earth as it is in heaven. Give us this day our daily bread. ` +
		`And forgive us our trespasses, as we forgive those who trespass against us. ` +
		`And lead us not into temptation; but deliver us from evil. For thine is the kingdom, ` +
		`the power and the glory, for ever and ever. Amen.`

	charC := utf8.RuneCountInString(str)
	fmt.Printf("Количество символов: %d\n", charC)

	substring := "heaven"
	if strings.Contains(str, substring) {
		pos := strings.Index(str, substring)
		fmt.Printf(substring, pos)
	} else {
		fmt.Printf(substring)
	}

	upper := strings.ToUpper(str)
	lower := strings.ToLower(str)
	fmt.Println("Верхний регистр:", upper)
	fmt.Println("Нижний регистр:", lower)

	replaced := strings.ReplaceAll(str, "kingdom", "Kingdom")
	fmt.Println(replaced)
}

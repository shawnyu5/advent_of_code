package day3

import (
	"io/ioutil"
	"log"
)

func Day3() {
	input := loadData()
	// alphabet.NewAlphabet()
	// fmt.Println(DNA.AllValid([]Letter("acgatcgatatagctatnagcatgc")))

}

func loadData() string {
	f, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return string(f)
}

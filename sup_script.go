package main

import (
	"bufio"
	"os"
)

// ReadFileLines читает строки из файла по переданному пути и возвращает их в виде среза строк.
func ReadFileLines(filePath string) ([]string, error) {
	var lines []string

	// Открытие файла
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Создание сканера для чтения файла построчно
	scanner := bufio.NewScanner(file)

	// Итерация по строкам файла
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

// package main

// import (
// 	"fmt"
// )
// type animal interface {
// 	makeSound()
// }

// type cat struct{}
// type dog struct{}


// func (c *cat) makeSound() {
// 	fmt.Println("meow")

// }

// func (d *dog) makeSound() {
// 	fmt.Println("gav")

// }


// func main() {
// 	var c,d animal = &cat{}, &dog{}
// 	c.makeSound()
// 	d.makeSound()
// }





// import (
// 	"fmt"
// 	"math"
// )

// type Vertex struct {
// 	X, Y float64
// }

// func Abs(v Vertex) float64 {
// 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }

// func Scale(v Vertex, f float64) {
// 	v.X = v.X * f
// 	v.Y = v.Y * f
// }

// func main() {
// 	v := Vertex{3, 4}
// 	Scale(&v, 10)
// 	fmt.Println(Abs(v))
// }


// import "fmt"

// // fibonacci is a function that returns
// // a function that returns an int.
// func fibonacci() func() int {

// 	a,b := 0,1

// 	return func() int {
// 		res := a

// 		a,b = b, a+b

// 		return res
// 	}

// }

// func main() {
// 	f := fibonacci()
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(f())
// 	}
// }

// import (
// 	"strings"
// 	// "golang.org/x/tour/wc"
// 	"fmt"
// )

// func WordCount(s string) map[string]int {
// 	// Создаем карту для подсчета количества вхождений каждого слова
// 	wordCounts := make(map[string]int)

// 	// Разбиваем строку на слова
// 	words := strings.Fields(s)

// 	// Подсчитываем количество вхождений каждого слова
// 	for _, word := range words {
// 		wordCounts[word]++
// 	}

// 	return wordCounts
// }

// func main() {
// 	// wc.Test(WordCount)
// 	counter := WordCount("i love i live")
// 	fmt.Sprintf("%T", counter)
// }

// import (
// 	"golang.org/x/tour/pic"
// )

// func Pic(dx, dy int) [][]uint8 {
// 	pic := make([][]uint8,dy)

// 	for y:=0; y < dy; y++ {
// 		row := make([]uint8, dx)

// 		for x:=0; x < dx; x++ {
// 			row[x] = uint8(x*y)
// 		}
// 		pic[y] = row
// 	}
// 	return pic
// }

// func main() {
// 	pic.Show(Pic)
// }


// func main() {
// 	pow := make([]int, 10)
// 	for i := range pow {
// 		pow[i] = 1 << uint(i) // == 2**i
// 	}
// 	for _, value := range pow {
// 		fmt.Printf("%d\n", value)
// 	}
// }

// func main() {
// 	i, j := 42, 2701

// 	p := &i         // point to i
// 	fmt.Println(*p) // read i through the pointer
// 	*p = 21         // set i through the pointer
// 	fmt.Println(i)  // see the new value of i

// 	p = &j         // point to j
// 	*p = *p / 37   // divide j through the pointer
// 	fmt.Println(j) // see the new value of j
// }

// func main() {
// 	fmt.Println("When's Saturday?")
// 	today := time.Now().Weekday()

// 	fmt.Printf("%s.\n", today)
	
// 	switch time.Saturday {
// 	case today + 0:
// 		fmt.Println("Today.")
// 	case today + 1:
// 		fmt.Println("Tomorrow.")
// 	case today + 2:
// 		fmt.Println("In two days.")
// 	default:
// 		fmt.Println("Too far away.")
// 	}
// }

// func pow(x, n, lim float64) float64 {
// 	if v := math.Pow(x, n); v < lim {
// 		return v
// 	} else {
// 		fmt.Printf("%g >= %g\n", v, lim)
// 	}
// 	// can't use v here, though
// 	return lim
// }

// func main() {
// 	fmt.Println(
// 		pow(3, 2, 10),
// 		pow(3, 3, 20),
// 		5,
// 	)
// }


// import (
// 	"fmt"
// 	"time"
// )

// // greeting returns a pleasant, semi-useful greeting.
// func greeting() string {
// 	return "Hello world, the time is: " + time.Now().String()
// }

// func main() {
// 	fmt.Println(greeting())
// 	fmt.Println(greeting())
// 	fmt.Println(greeting())
// }
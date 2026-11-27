package main
import "fmt"

func main() {

	//  1. Traditional for loop
	for i:=0;i<5;i++{
		fmt.Println("Traditional for loop:",i)
	}

	// 2. While style loop
	j:=0
	for j<5{
		fmt.Println("While style loop:",j)
		j++
	}
	// 3. Infinite loop:runs forever until break statement is used
	count:=0
	for{
		fmt.Println(count)
		count++
		if count==5{
			break
		}
	}
	// 4. Range based for loop:used to iterate over elements of an array, slice, map, or string.
	numbers :=[]int{10,20,30,40}
	for index,value := range numbers{
		fmt.Println(index,value)
	}

	// 5. Using Blank Identifier:If you don't need the index or value, you can use the blank identifier (_) to ignore them.
	for _,value := range numbers{
		fmt.Println(value)
	}
}
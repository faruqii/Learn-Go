package main

import "fmt"

// implements a array and record, record minimal has 3 different fields of type

type Record struct {
	Id int
	Name string
	Salary float64
}

type Array struct {
	data []Record
}

// insert a record into array
func (a *Array) Insert(id int, name string, salary float64) {
	a.data = append(a.data, Record{id, name, salary})
}

// delete a record from array
func (a *Array) Delete(id int) {
	for i, v := range a.data {
		if v.Id == id {
			a.data = append(a.data[:i], a.data[i+1:]...)
			break
		}
	}
}

// show all records in array
func (a *Array) Show() {
	for _, v := range a.data {
		fmt.Println(v)
	}
}

// sequential search
func (a *Array) SequentialSearch(id int) (Record, bool) {
	for _, v := range a.data {
		if v.Id == id {
			return v, true
		}
	}
	return Record{}, false
}

// binary search
func (a *Array) BinarySearch(id int) (Record, bool) {
	low := 0
	high := len(a.data) - 1
	for low <= high {
		mid := (low + high) / 2
		if a.data[mid].Id == id {
			return a.data[mid], true
		} else if a.data[mid].Id > id {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return Record{}, false
}

// bubble sort
func (a *Array) BubbleSort() {
	for i := 0; i < len(a.data)-1; i++ {
		for j := 0; j < len(a.data)-i-1; j++ {
			if a.data[j].Id > a.data[j+1].Id {
				a.data[j], a.data[j+1] = a.data[j+1], a.data[j]
			}
		}
	}
}

// selection sort
func (a *Array) SelectionSort() {
	for i := 0; i < len(a.data)-1; i++ {
		min := i
		for j := i + 1; j < len(a.data); j++ {
			if a.data[j].Id < a.data[min].Id {
				min = j
			}
		}
		a.data[i], a.data[min] = a.data[min], a.data[i]
	}
}

// insertion sort
func (a *Array) InsertionSort() {
	for i := 1; i < len(a.data); i++ {
		for j := i; j > 0; j-- {
			if a.data[j].Id < a.data[j-1].Id {
				a.data[j], a.data[j-1] = a.data[j-1], a.data[j]
			}
		}
	}
}

// main function
func main() {
	// create menu
	fmt.Println("1. Insert")
	fmt.Println("2. Delete")
	fmt.Println("3. Show")
	fmt.Println("4. Sequential Search")
	fmt.Println("5. Binary Search")
	fmt.Println("6. Bubble Sort")
	fmt.Println("7. Selection Sort")
	fmt.Println("8. Insertion Sort")
	fmt.Println("9. Exit")
	fmt.Println("Please input your choice:")
	var choice int
	fmt.Scanln(&choice)
	for {
		if choice == 1 {
			var id int
			var name string
			var salary float64
			fmt.Println("Please input id:")
			fmt.Scanln(&id)
			fmt.Println("Please input name:")
			fmt.Scanln(&name)
			fmt.Println("Please input salary:")
			fmt.Scanln(&salary)
			a := Array{}
			a.Insert(id, name, salary)
			fmt.Println("Insert successfully!")
		} else if choice == 2 {
			var id int
			fmt.Println("Please input id:")
			fmt.Scanln(&id)
			a := Array{}
			a.Delete(id)
	
		} else if choice == 3 {
			a := Array{}
			a.Show()
	
		} else if choice == 4 {
			var id int
			fmt.Println("Please input id:")
			fmt.Scanln(&id)
			a := Array{}
			r, ok := a.SequentialSearch(id)
			if ok {
				fmt.Println("Found:", r)
			} else {
				fmt.Println("Not found")
			}
	
		} else if choice == 5 {
			var id int
			fmt.Println("Please input id:")
			fmt.Scanln(&id)
			a := Array{}
			r, ok := a.BinarySearch(id)
			if ok {
				fmt.Println("Found:", r)
			} else {
				fmt.Println("Not found")
			}
	
		} else if choice == 6 {
			a := Array{}
			a.BubbleSort()
			a.Show()
	
		} else if choice == 7 {
			a := Array{}
			a.SelectionSort()
			a.Show()
	
		} else if choice == 8 {
			a := Array{}
			a.InsertionSort()
			a.Show()
	
		} else if choice == 9 {
			fmt.Println("Bye")
			break
		} else {
			fmt.Println("Invalid choice")
		}
	}
	
}




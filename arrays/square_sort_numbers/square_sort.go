package square_sort_numbers

import "fmt"

func squareSort(nums []int) []int {
	sq := make([]int, len(nums))
	for i, v := range nums {
		sq[i] = v*v
	}
	fmt.Println("before", sq)
	bubbleSort(sq)
	fmt.Println("after", sq)
	return sq
}

func bubbleSort(items []int) {  
	L:=len(items)  
	for i:=0;i<L;i++{  
		for j:=0;j< (L-1-i);j++{  
			if items[j] > items[j+1]{  
				items[j], items[j+1] = items[j+1], items[j]  
			}  
		}  
	}  
} 
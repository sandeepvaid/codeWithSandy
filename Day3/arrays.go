package main
import "fmt"

func main() {
	fmt.Println("Hey I am array")
	var arr1 = [3]int{1,2,3};
	//Change element in aray
	arr1[0]= 4;

	//Initialize specific element in the array , Here we initialize 1st index of the array
	arr2 := [3]int{1:2};
	fmt.Println(arr1);
	fmt.Println(arr2);

	//Length properies of the array
	arr3 := [3]string{};
	fmt.Println(len(arr3));

	//Slicing in an array
	fmt.Println(arr1[1:]);
	
}
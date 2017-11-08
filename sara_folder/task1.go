package sara_folder

import "fmt"

func main() {
	i := 1
	//initial print
	fmt.Print("Enter a number between 1-100: ")
	//user input
	var userNum int
	fmt.Scanf("%d", &userNum)
	
	for i <=userNum{
		if i%2 == 0{
			fmt.Println(i, "even")
		}else{
			fmt.Println(i, "odd")
		}//end else
		i = i+1
	}//end for
}//end main
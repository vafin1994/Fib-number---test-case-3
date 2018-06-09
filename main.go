package main

import (
	"fmt"
	"encoding/json"

)

func main(){

	var num int;
	var mistakes int = 0;
	var rightAnswersCounter int = 0;

	fmt.Println("Please insert Fibonacci sequence (starting from 0) ");


	for  i := 0; ; i++{
		mapVar1 := map[int]int{i: fibnumbers(i)}
		mapVar2, _ := json.Marshal(mapVar1)


		if(mistakes == 3){
		fmt.Println("3 mistakes, YOU LOOSE");
		break;
		}

		fmt.Scan(&num);
		if(num != fibnumbers(i)) {
			fmt.Println(string(mapVar2));
			mistakes++
			rightAnswersCounter = 0;

		}else if(num == fibnumbers(i)){

			rightAnswersCounter++
			if(rightAnswersCounter == 10){
			fmt.Println("10 right answers in a row, YOU WIN");
			}
		}


	}


}

func fibnumbers(n int)(int){

	if (n ==0){
		return 0;
	}
	if (n==1){
		return 1;
	}
	return(fibnumbers(n - 2) + fibnumbers(n - 1))

}




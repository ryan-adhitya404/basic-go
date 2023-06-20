package main
import(
	"fmt"
	"time"
)

func main(){
	i:=2
	fmt.Println("Write",i,"as")
	switch i{
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday(){
	case time.Saturday, time.Sunday:
		fmt.Println("its the weekend")
	default:
		fmt.Println("its a weekday")
	}

	whatAmI:=func(i interface{}){
		switch t:=i.(type){
		case bool:
			fmt.Println("im a boolean")
		case int:
			fmt.Println("im an integer")
		default:
			fmt.Println("dont know type %T\n",t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
package main
import "fmt"

func populatePhoneBook(phone *map[string]string)  {
	var T int
	fmt.Scan(&T)
	for i:= 0 ; i < T; i++{
		name, num := readPhoneInput()
		(*phone)[name]= num
	}
}

func printEntriesFrom(phone *map[string]string)  {
	for {
		var e string
		_ ,err := fmt.Scanf("%s\n", &e)
		if ( err != nil){
			break
		}
		if val, ok := (*phone)[e]; ok {
			fmt.Printf("%s=%s\n", e,val)
		}else{
			fmt.Println("Not found")
		}
	}
}
func main() {
	phoneBook:= make(map[string]string)
	populatePhoneBook(&phoneBook)
	printEntriesFrom(&phoneBook)
}
func readPhoneInput() (string, string) {
	var f, s string
	fmt.Scanf("%s %s\n", &f, &s)
	return f,s
}

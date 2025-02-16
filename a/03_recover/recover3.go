package main

import (
	"fmt"
)

func catch (f func(any)) {
	if er := recover(); er != nil {
		f(er)
	}
}

func main3() {
	
	for i := 0; i < 10; i++ {
		func() {
			defer catch(func(er any) { 
				fmt.Println(er) 
			})
			if i % 2 == 1 { panic(fmt.Sprintf("odd number %d", i)) }
			fmt.Printf("%d\n", i)
		}()
	}

	for i := 0; i < 10; i++ {
		{
			defer catch(func(er any) { 
				fmt.Println(er) 
			})
			if i % 2 == 1 { panic(fmt.Sprintf("odd number %d", i)) }
			fmt.Printf("%d\n", i)		
		}
	}	
}

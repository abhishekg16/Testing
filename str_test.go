package str

import "testing"
import "log"
import "fmt"


func TestStr(t *testing.T) {
	var test = [] struct {
		s , want string
	}{
		{"Hi","hi"},
		{"AbhiSHek","abhishek"},
	}
	for _,c := range test {
		log.Printf("Testing : %q",c.s)
		got := String(c.s) 
		if got != c.want {
			t.Errorf("String(%q) == %q , want = %q", c.s, got, c.want)
		}
	}
	
}

func BenchmarkString(b * testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf(String("Abhishek"))
	}   
}

/*
func ExampleString() {
	fmt.Println(String("Hi"))
	//Output:
	//hi
}
*/

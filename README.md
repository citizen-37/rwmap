### Usage

```go
import "github.com/citizen-37/rwmap"

func main() {
	example := rwmap.New[int, int]()
	
	example.Set(1, 1)
	
	fmt.Println(example.Get(1))
	fmt.Println(example.List())
}
```
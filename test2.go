package main
import (
	"fmt"
	"log"
)
func main() {
	n := 0
	fmt.Print("Введите целое число: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Вы ввели число: %d\n", n)
}
git init
git add .
git commit -m "first commit, version 0.1"
1.
git checkout -b develop
git branch feature/get_all_types
git checkout feature/get_all_types
package main
import (
"fmt"
"log"
)
func main() {
	var n string
	fmt.Print("Введите целое число: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Вы ввели число: %v\n", n)
}
git add .
git commit -m "new feature"
git checkout develop
git merge feature/get_all_types
2.
git checkout -b release/1.0
git tag -a v1.0 -m "Новый релиз"
git checkout master
git merge release/1.0
3.
git checkout -b hotfix/1.1
package main
import (
"fmt"
"log"
)
func main() {
	var n string
	fmt.Print("Введите данные: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Вы ввели следующие данные: %v\n", n)
}
git add .
git commit -m "hotfix v1.1"
git checkout master
git merge hotfix/1.1
git checkout develop
git merge hotfix/1.1


//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//func main() {
//
//	var m = [...]int{1, 2, 3}
//
//	for i, v := range m {
//		go func() {
//
//			fmt.Println(i, v)
//		}()
//		//fmt.Println("aaaaaaaaaa")
//		time.Sleep(time.Second)
//	}
//
//	time.Sleep(time.Second * 3)
//}

//package main
//
//import "fmt"
//
//type Math struct {
//	x, y int
//}
//
//var m = map[string]*Math{
//	"foo": &Math{2, 3},
//}
//
//func main() {
//	m["foo"].x = 4
//	fmt.Println(m["foo"].x)
//}

//package main
//
//import "fmt"
//
//func f(n int) (r int) {
//	defer func() {
//		r += n
//		recover()
//	}()
//
//	var f func()
//
//	defer f()
//	f = func() {
//		r += 2
//	}
//	return n + 1
//}
//
//func main() {
//	fmt.Println(f(3))
//}

//package main
//
//func main() {
//	if a := 1; false {
//	} else if b := 2; false {
//	} else {
//		println(a, b)
//	}
//}

//package main
//
//import "fmt"
//
//var p *int
//
//func foo() (*int, error) {
//	var i int = 5
//	return &i, nil
//}
//
//func bar() {
//	//use p
//	fmt.Println(*p)
//}
//
//func main() {
//	p, err := foo()
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	bar()
//	fmt.Println(*p)
//}

//package main
//
//import "fmt"
//
//func main(){
//	for i := 0; i < 10; i++{
//		fmt.Println(i)
//	}
//}

//package main
//
//import "fmt"
//
//func main() {
//
//	var m = map[string]int{
//		"A": 21,
//		"B": 22,
//		"C": 23,
//	}
//	counter := 0
//	for k, v := range m {
//		if counter == 0 {
//			fmt.Println("yyyyyyyyyyyy")
//			delete(m, "A")
//		}
//		counter++
//		fmt.Println(k, v)
//	}
//	fmt.Println("counter is ", counter)
//}
//
//

//package main
//
//import (
//	"database/sql"
//	"encoding/json"
//	"net/http"
//
//	_ "github.com/mattn/go-sqlite3"
//	"go.uber.org/dig"
//)
//
//type Config struct {
//	Enabled      bool
//	DatabasePath string
//	Port         string
//}
//
//// NewConfig config的构造器
//func NewConfig() *Config {
//	return &Config{
//		Enabled:      true,				//Enabled控制应用是否应该返回真实数据
//		DatabasePath: "./example.db",	//DatabasePath表示数据库位置
//		Port:         "8000",			//Port表示服务器运行时监听的端口。
//	}
//}
//
//type Person struct {
//	Id   int    `json:"id"`
//	Name string `json:"name"`
//	Age  int    `json:"age"`
//}
//
//
//func (repository *PersonRepository) FindAll() []*Person {
//	rows, _ := repository.database.Query(`SELECT id, name, age FROM people;`)
//	defer rows.Close()
//
//	people := []*Person{}
//
//	for rows.Next() {
//		var (
//			id   int
//			name string
//			age  int
//		)
//
//		rows.Scan(&id, &name, &age)
//
//		people = append(people, &Person{
//			Id:   id,
//			Name: name,
//			Age:  age,
//		})
//	}
//
//	return people
//}
//
//
//
//// NewPersonRepository 这个是personRepository的构造器
//func NewPersonRepository(database *sql.DB) *PersonRepository {
//	return &PersonRepository{database: database}
//}
//
//type PersonRepository struct {
//	database *sql.DB
//}
//
//type PersonService struct {
//	config     *Config
//	repository *PersonRepository
//}
//
//func (service *PersonService) FindAll() []*Person {
//	if service.config.Enabled {
//		return service.repository.FindAll()
//	}
//
//	return []*Person{}
//}
//
//func NewPersonService(config *Config, repository *PersonRepository) *PersonService {
//	return &PersonService{config: config, repository: repository}
//}
//
//type Server struct {
//	config        *Config
//	personService *PersonService
//}
//
//func (server *Server) Handler() http.Handler {
//	mux := http.NewServeMux()
//
//	mux.HandleFunc("/people", server.findPeople)
//
//	return mux
//}
//
//func (server *Server) Run() {
//	httpServer := &http.Server{
//		Addr:    ":" + server.config.Port,
//		Handler: server.Handler(),
//	}
//
//	httpServer.ListenAndServe()
//}
//
//func (server *Server) findPeople(writer http.ResponseWriter, request *http.Request) {
//	people := server.personService.FindAll()
//	bytes, _ := json.Marshal(people)
//
//	writer.Header().Set("Content-Type", "application/json")
//	writer.WriteHeader(http.StatusOK)
//	writer.Write(bytes)
//}
//
//func NewServer(config *Config, personService *PersonService) *Server {
//	return &Server{
//		config:        config,
//		personService: personService,
//	}
//}
//
//func ConnectDatabase(config *Config) (*sql.DB, error) {
//	return sql.Open("sqlite3", config.DatabasePath)
//}
//
//func BuildContainer() *dig.Container {
//	container := dig.New()
//
//	container.Provide(NewConfig)
//	container.Provide(ConnectDatabase)
//	container.Provide(NewPersonRepository)
//	container.Provide(NewPersonService)
//	container.Provide(NewServer)
//
//	return container
//}
//
//func main() {
//	container := BuildContainer()
//
//	err := container.Invoke(func(server *Server) {
//		server.Run()
//	})
//
//	if err != nil {
//		panic(err)
//	}
//}

// The manual way
//
// func main() {
// 	config := NewConfig()
//
// 	db, err := ConnectDatabase(config)
//
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	personRepository := NewPersonRepository(db)
//
// 	personService := NewPersonService(config, personRepository)
//
// 	server := NewServer(config, personService)
//
// 	server.Run()
// }
//package main
//
//import "fmt"
//
//type ConfigOne struct {
//	Daemon string
//}
//
//func (c *ConfigOne) String1() string {
//	return fmt.Sprintf("print: %v", c)
//}
//
//func main() {
//	c := &ConfigOne{}
//	c.String1()
//}

package main

func main() {

	println(DeferTest1(1))
	println(DeferTest2(1))
}

func DeferTest1(i int) (r int) {
	r = i
	defer func() {
		r += 3
	}()
	return r
}

func DeferTest2(i int) (r int) {
	defer func() {
		r += i
	}()
	return 2
}

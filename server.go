package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"r_d/handlers"
	"r_d/repository"
)

type Server struct {
	db          *sql.DB
	userHandler *handlers.UserHandler
}

func NewServer(db *sql.DB) *Server {
	userRepo := repository.NewUserRepository(db)

	userHandler := handlers.NewUserHandler(userRepo)

	return &Server{
		db:          db,
		userHandler: userHandler,
	}
}

func (s *Server) setupRoutes() {
	http.HandleFunc("/health", s.userHandler.Health)
	http.HandleFunc("/get-user", s.userHandler.User)
	http.HandleFunc("/get-users", s.userHandler.Users)
	http.HandleFunc("/create-user", s.userHandler.Create)
	http.HandleFunc("/update-user", s.userHandler.Update)
	http.HandleFunc("/delete-user", s.userHandler.Delete)
}

func (s *Server) Run(port string) {
	s.setupRoutes()

	fmt.Printf("🚀 Server is running on http://localhost%s\n", port)
	fmt.Println("📍 Endpoints:")
	fmt.Println("   GET  /health")
	fmt.Println("   GET  /get-user")
	fmt.Println("   GET  /get-users")
	fmt.Println("   POST /create-user")
	fmt.Println("   PUT  /update-user")
	fmt.Println("   DELETE  /delete-user")

	log.Fatal(http.ListenAndServe(port, nil))
}

/*

### 🔍 Пояснення:

type Server struct { ... } — структура сервера зі всіма залежностями.

func NewServer(db *sql.DB) *Server — конструктор:
1. Створює репозиторій
2. Створює handler
3. Повертає готовий сервер

Потік залежностей:**
DB → Repository → Handler → Server


s.setupRoutes() — реєструє всі HTTP маршрути.

http.HandleFunc("/health", s.userHandler.Health):
- /health — шлях
- s.userHandler.Health — функція-обробник

---

 🎯 Як це працює разом?

1. main.go запускається
2. Завантажується config
3. Підключається database
4. Створюється Server з усіма залежностями:
   DB → UserRepository → UserHandler → Server
5. Реєструються роути
6. Сервер слухає на порту :8080
7. Приходить запит POST /create-user
8. HTTP → UserHandler.Create → UserRepository.Create → MySQL
9. Відправляється JSON відповідь

*/

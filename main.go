package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/JridyFery/pikastAPI/config"
	"github.com/JridyFery/pikastAPI/handlers"
	models "github.com/JridyFery/pikastAPI/models"
	"github.com/JridyFery/pikastAPI/repository"
	router "github.com/JridyFery/pikastAPI/router"
	"github.com/rs/cors"
	"github.com/spf13/viper"
)

func main() {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost", "http://localhost:8080"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "PUT", "HEAD"},
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	})
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	var configuration config.Configuration

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	//connect to the data base
	UserName := configuration.Database.UserName
	Password := configuration.Database.Password
	DataBase := configuration.Database.DataBase
	Charset := configuration.Database.Charset
	ParseTime := configuration.Database.ParseTime
	db, err := gorm.Open("mysql", UserName+":"+Password+"@/"+DataBase+"?charset="+
		Charset+"&parseTime="+ParseTime+"&loc=Local")
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()
	db.AutoMigrate(&models.Pokemon{}, &models.Player{})

	playerRepo := repository.PlayerRepo{db}
	playerHandler := handlers.PlayerHandler{playerRepo}
	pokemonRepo := repository.PokemonRepo{db}
	pokemonHandler := handlers.PokemonHandler{pokemonRepo}
	// Init Router
	r := mux.NewRouter()
	// serve static files
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	PlayerRouterHandler := router.PlayerRouterHandler{Router: r, Handler: playerHandler}
	pokemonRouterHandler := router.PokemonRouterHandler{Router: r, Handler: pokemonHandler}
	PlayerRouterHandler.HandleFunctions()
	pokemonRouterHandler.HandleFunctions()
	// start server
	port := ":" + strconv.Itoa(configuration.Server.Port)
	handler := c.Handler(r)

	log.Fatal(http.ListenAndServe(port, handler))
}

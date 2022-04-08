//Para test
//package main
package bases

import (
	"encoding/json"
	"fmt"

	redis "github.com/garyburd/redigo/redis"
)

type Game struct {
	Nombre_Juego   string `json:"nombre_juego"`
	Nombre_Ganador string `json:"nombre_ganador"`
}

var ADDRREDIS = "35.206.76.88"
var PASSWORDREDIS = "1234"
var PORTREDIS = "6379"

/*
// Para Test
func main() {
	saveDataRedisStr("Juego_Random", "25")
}
*/

func saveDataRedisObj(game *Game) {
	// Establecer conexion
	c, err := redis.Dial("tcp", fmt.Sprintf("%s:%s", ADDRREDIS, PORTREDIS))

	if err != nil {
		fmt.Println("Not found Database 404")
	}

	// Authenticacion
	_, err = c.Do("AUTH", "1234")

	defer c.Close()

	json, err := json.Marshal(game)

	if err != nil {
		fmt.Println("No Generate JSON")
	}

	_, err = c.Do("LPUSH", "gamesobj", json)
	if err != nil {
		fmt.Println("Error Insert Data")
	} else {
		fmt.Println("Data Inserted Redis")
	}
	/*
		// See Data
		val, err := redis.Strings(c.Do("LRANGE", "gamesobj", 0, -1))

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(val)
	*/
}

func saveDataRedisStr(game, winner string) {

	c, err := redis.Dial("tcp", fmt.Sprintf("%s:%s", ADDRREDIS, PORTREDIS))

	if err != nil {
		fmt.Println("Not found Database 404")
	}

	_, autherror := c.Do("AUTH", "1234")

	if autherror != nil {
		fmt.Println("Error authentication")
	}

	defer c.Close()

	_, err = c.Do("LPUSH", "gamesstr", game, winner)
	if err != nil {
		fmt.Println("Error Insert Data")
	} else {
		fmt.Println("Data Inserted Redis")
	}
	/*
		// See Data
		val, err := redis.Strings(c.Do("LRANGE", "gamesstr", 0, -1))

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(val)
	*/

}


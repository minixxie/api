package lib

import "gopkg.in/yaml.v2"
import "github.com/dgrijalva/jwt-go"
import "api/entity"
import "time"
import "os"
import "bufio"
import "log"

func GenJWT(userId int64) string {
	jwtToken := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := jwtToken.Claims.(jwt.MapClaims)
	claims["userId"] = userId
	// claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	// Generate encoded token and send it as response.

	t, err := jwtToken.SignedString([]byte("hello123"))
	if err != nil {
	    return ""
	}

	return t
}

func LoadConfig() (entity.Config, error) {
    config := entity.Config{}

    var err error

    var env string
    env = os.Getenv("ENV")
    log.Printf("ENV = %s", env)
    if env = os.Getenv("ENV"); env == "" {
		env = "ldev"
	}


    file, err := os.Open("/config/" + env + ".yml")
    if err != nil {
        return config, err
    }
    defer file.Close()

    stats, statsErr := file.Stat()
    if statsErr != nil {
        return config, statsErr
    }

    var size int64 = stats.Size()
    bytes := make([]byte, size)

    bufr := bufio.NewReader(file)
    _, err = bufr.Read(bytes)

    err = yaml.Unmarshal(bytes, &config)
    if err != nil {
        return config, err
    }
    return config, nil
}
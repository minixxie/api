package endpoints

import "api/entity"
import "api/lib"

import "log"
import "gopkg.in/labstack/echo.v3"
import "database/sql"
// import "gopkg.in/go-on/go.uuid.v1"
import "time"
import "net/http"
// import "github.com/dgrijalva/jwt-go"
// import "github.com/labstack/echo/middleware"

func Signup(userDB *sql.DB) func(ctx echo.Context) error {
    type RequestJson struct {
        Phone string            `json:"phone"`
        Password string         `json:"password"`
    }
    type ResponseJson struct {
        userId int64             `json:"userId"`
        SignupAt int64           `json:"signupAt"`
        SignupAtISO string       `json:"signupAtISO"`
        Jwt string               `json:"jwt"`
    }
    return func(ctx echo.Context) error {
        req := RequestJson{}
        if err := ctx.Bind(&req); err != nil {
            log.Fatal("Cannot bind req: ", err)
        }
        log.Printf("req.Phone = %s", req.Phone)

        nowNano := time.Now().UnixNano()
        nowMicro := nowNano / 1000
        nowISO := time.Unix(0, nowNano).UTC().Format(time.RFC3339Nano)
        user := entity.User{
            Phone: req.Phone,
            Password: req.Password,
            CreatedAtMicroseconds: nowMicro,
            CreatedAtISO: nowISO,
            UpdatedAtMicroseconds: nowMicro,
            UpdatedAtISO: nowISO,
        }

        stmt, err := userDB.Prepare(`
            INSERT INTO "User" ("phone", "password", "createdAtMicroseconds", "updatedAtMicroseconds") 
            VALUES ($1, $2, $3, $4) 
            RETURNING id
        `)
        defer stmt.Close()
        if err != nil {
            log.Fatal("Cannot prepare DB statement: ", err)
        }

log.Printf("to insert: user = %v", user)
log.Printf("to insert, phone = \"%s\"", user.Phone)
log.Printf("to insert, Password = \"%s\"", user.Password)

        var id int64
        err = stmt.QueryRow(user.Phone, user.Password, user.CreatedAtMicroseconds, user.UpdatedAtMicroseconds).Scan(&id)
        if err != nil {
            log.Fatal("Cannot run insert statement: ", err)
        } else {
            user.Id = id
        }

        jwt := lib.GenJWT(user.Id)

        return ctx.JSON(http.StatusOK, ResponseJson{
            userId: user.Id,
            SignupAt: user.CreatedAtMicroseconds,
            SignupAtISO: user.CreatedAtISO,
            Jwt: jwt,
        })

        // if req.Phone == "+85267633535" && req.Password == "hello123" {
        //     nowNano := time.Now().UnixNano()
        //     nowMicro := nowNano / 1000
        //     nowISO := time.Unix(0, nowNano).UTC().Format(time.RFC3339Nano)

        //     jwtToken := jwt.New(jwt.SigningMethodHS256)

        //     // Set claims
        //     claims := jwtToken.Claims.(jwt.MapClaims)
        //     claims["userId"] = "1234"
        //     // claims["admin"] = true
        //     claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
        //     // Generate encoded token and send it as response.

        //     t, err := jwtToken.SignedString([]byte("secret"))
        //     if err != nil {
        //         return err
        //     }

        //     return ctx.JSON(http.StatusOK, ResponseJson{
        //         SignupAt: now,
        //         SignupAtISO: nowISO,
        //         JwtToken: t,
        //     })
        // } else {
        //     return echo.ErrUnauthorized
        // }
    }
}


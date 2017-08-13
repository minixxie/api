package endpoints

import "api/lib"
import "api/entity"

import "log"
import "gopkg.in/labstack/echo.v3"
import "database/sql"
// import "gopkg.in/go-on/go.uuid.v1"
import "time"
// import "api/entity"
import "net/http"

// import "github.com/dgrijalva/jwt-go"
// import "github.com/labstack/echo/middleware"

func Signin(userDB *sql.DB) func(ctx echo.Context) error {
    type RequestJson struct {
        Phone string            `json:"phone"`
        Password string         `json:"password"`
    }
    type ResponseJson struct {
        SigninAt int64           `json:"signinAt"`
        SigninAtISO string       `json:"signinAtISO"`
        Jwt string               `json:"jwt"`
    }
    return func(ctx echo.Context) error {
        req := RequestJson{}
        if err := ctx.Bind(&req); err != nil {
            log.Fatal("Cannot bind req: ", err)
        }
        log.Printf("req.Phone = %s", req.Phone)

        stmt, err := userDB.Prepare(`
            SELECT "id", "phone", "password", "createdAtMicroseconds", "updatedAtMicroseconds"
            FROM
            "User"
            WHERE "phone" = $1
            LIMIT 1
        `)
        defer stmt.Close()
        if err != nil {
            log.Fatal("Cannot prepare DB statement: ", err)
        }

        row := stmt.QueryRow(req.Phone)

        var id int64
        var phone string
        var password string
        var createdAtMicroseconds int64
        var updatedAtMicroseconds int64
        err = row.Scan(&id, &phone, &password, &createdAtMicroseconds, &updatedAtMicroseconds)
        if err != nil {
            log.Fatal("row.Scan has error: ", err)
        }
log.Printf("loaded phone = \"%s\"", phone)
log.Printf("loaded password = \"%s\"", password)
        createdAtISO := time.Unix(0, createdAtMicroseconds*1000).UTC().Format(time.RFC3339Nano)
        updatedAtISO := time.Unix(0, updatedAtMicroseconds*1000).UTC().Format(time.RFC3339Nano)
        user := entity.User{
            Id: id,
            Phone: phone,
            Password: password,
            CreatedAtMicroseconds: createdAtMicroseconds,
            CreatedAtISO: createdAtISO,
            UpdatedAtMicroseconds: updatedAtMicroseconds,
            UpdatedAtISO: updatedAtISO,
        }

log.Printf("req = %v", req)
log.Printf("user = %v", user)

            log.Printf("req.Phone  = \"%s\"", req.Phone)
            log.Printf("user.Phone = \"%s\"", user.Phone)
            log.Printf("req.Password  = \"%s\"", req.Password)
            log.Printf("user.Password = \"%s\"", user.Password)

        if (req.Phone == user.Phone) && (req.Password == user.Password) {
            log.Printf("both equal")
            log.Printf("req.Phone  = \"%s\"", req.Phone)
            log.Printf("user.Phone = \"%s\"", user.Phone)
            log.Printf("req.Password  = \"%s\"", req.Password)
            log.Printf("user.Password = \"%s\"", user.Password)

            nowNano := time.Now().UnixNano()
            nowMicro := nowNano / 1000
            nowISO := time.Unix(0, nowNano).UTC().Format(time.RFC3339Nano)

            jwt := lib.GenJWT(user.Id)

            return ctx.JSON(http.StatusOK, ResponseJson{
                SigninAt: nowMicro,
                SigninAtISO: nowISO,
                Jwt: jwt,
            })
        } else {
            return echo.ErrUnauthorized
        }
    }
}


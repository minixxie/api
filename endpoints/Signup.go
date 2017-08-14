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
        Error string             `json:"error"`
        UserId int64             `json:"userId"`
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

        // (1) check if phone duplicates
        stmt1, err := userDB.Prepare(`
            SELECT count(1)
            FROM
            "User"
            WHERE "phone" = $1
            LIMIT 1
        `)
        defer stmt1.Close()
        if err != nil {
            log.Fatal("Cannot prepare DB statement: ", err)
        }
        row := stmt1.QueryRow(req.Phone)
        var duplicateCount int64
        err = row.Scan(&duplicateCount)
        if err != nil {
            log.Fatal("row.Scan has error: ", err)
        }
        log.Printf("row count ==> %d", duplicateCount)
        if duplicateCount > 0 {
            return ctx.JSON(http.StatusConflict, ResponseJson{
                Error: "PHONE_ALREADY_REGISTERED",
            })
        } else {
            // (2) create user if no duplicate
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

            stmt2, err := userDB.Prepare(`
                INSERT INTO "User" ("phone", "password", "createdAtMicroseconds", "updatedAtMicroseconds") 
                VALUES ($1, $2, $3, $4) 
                RETURNING id
            `)
            defer stmt2.Close()
            if err != nil {
                log.Fatal("Cannot prepare DB statement: ", err)
            }

    log.Printf("to insert: user = %v", user)
    log.Printf("to insert, phone = \"%s\"", user.Phone)
    log.Printf("to insert, Password = \"%s\"", user.Password)

            var id int64
            err = stmt2.QueryRow(user.Phone, user.Password, user.CreatedAtMicroseconds, user.UpdatedAtMicroseconds).Scan(&id)
            if err != nil {
                log.Fatal("Cannot run insert statement: ", err)
            } else {
                user.Id = id
            }

            jwt := lib.GenJWT(user.Id)

            return ctx.JSON(http.StatusOK, ResponseJson{
                UserId: user.Id,
                SignupAt: user.CreatedAtMicroseconds,
                SignupAtISO: user.CreatedAtISO,
                Jwt: jwt,
            })
        }





    }
}


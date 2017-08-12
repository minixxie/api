package endpoints

import "log"
import "gopkg.in/labstack/echo.v3"
import "database/sql"
// import "gopkg.in/go-on/go.uuid.v1"
import "time"
import "api/entity"
import "net/http"

func CreateOrder(mainDB *sql.DB) func(ctx echo.Context) error {
    type RequestJson struct {
        Title string            `json:"title"`
        OrderCategoryId int64   `json:"orderCategoryId"`
    }
    type ResponseJson struct {
        Order entity.Order      `json:"order"`
    }
    return func(ctx echo.Context) error {
        req := RequestJson{}
        if err := ctx.Bind(&req); err != nil {
            log.Fatal("Cannot bind req: ", err)
        }
        log.Printf("req.title = %s", req.Title)

        // id := uuid.NewV4()
        nowNano := time.Now().UnixNano()
        nowMicro := nowNano / 1000
        nowISO := time.Unix(0, nowNano).UTC().Format(time.RFC3339Nano)
        order := entity.Order{
            // Id: id.String(),
            Title: req.Title,
            OrderCategoryId: req.OrderCategoryId,
            CreatedAtMicroseconds: nowMicro,
            CreatedAtISO: nowISO,
            UpdatedAtMicroseconds: nowMicro,
            UpdatedAtISO: nowISO,
        }

        // stmt, err := mainDBmainDB.Prepare("INSERT INTO \"Orders\" (\"title\", \"orderCategoryId\", \"createdAt\", \"updatedAt\") VALUES (?, ?, ?, ?)")
        // stmt, err := mainDB.Prepare(`
        //     INSERT INTO "Orders" ("title", "orderCategoryId", "createdAt", "updatedAt") VALUES ($1, (SELECT "id" FROM "OrderCategory" WHERE "key" = $2), $3, $4)
        // `)
        stmt, err := mainDB.Prepare(`
            INSERT INTO "Order" ("title", "orderCategoryId", "createdAtMicroseconds", "updatedAtMicroseconds") 
            VALUES ($1, $2, $3, $4) 
            RETURNING id
        `)
        defer stmt.Close()
        if err != nil {
            log.Fatal("Cannot prepare DB statement: ", err)
        }

log.Printf("to insert: order = %v", order)
        var id int64
        err = stmt.QueryRow(order.Title, order.OrderCategoryId, order.CreatedAtMicroseconds, order.UpdatedAtMicroseconds).Scan(&id)
        if err != nil {
            log.Fatal("Cannot run insert statement: ", err)
        } else {
            order.Id = id
        }
        // log.Printf("res = %v", res)

        return ctx.JSON(http.StatusOK, ResponseJson{
            Order: order,
        })

    }
}


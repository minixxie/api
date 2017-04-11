package endpoints

import "log"
import "gopkg.in/labstack/echo.v3"
import "database/sql"
// import _ "gopkg.in/go-sql-driver/mysql.v1"
import "time"
import "api/entity"
import "net/http"

func GetOrders(db *sql.DB) func(ctx echo.Context) error {
    type ResponseJson struct {
        Count int                  `json:"count"`
        Orders []entity.Order      `json:"orders"`
    }
    return func(ctx echo.Context) error {
        log.Printf("Hello World")
        rows, err := db.Query(`
            SELECT "id", "title", "orderCategoryId", "createdAtMicroseconds", "updatedAtMicroseconds"
            FROM
            "Order" 
            ORDER BY "createdAtMicroseconds" DESC 
            LIMIT 100
        `)
        if err != nil {
            log.Fatal("Cannot select Orders", err)
        }
        var orders []entity.Order
        for rows.Next() {
            var id int64
            var title string
            var orderCategoryId int64
            var createdAtMicroseconds int64
            var updatedAtMicroseconds int64
            err = rows.Scan(&id, &title, &orderCategoryId, &createdAtMicroseconds, &updatedAtMicroseconds)
            if err != nil {
                log.Fatal("rows.Scan has error: ", err)
            }
            // log.Printf("endpoint_get_orders: title = %s, orderCategoryId = %s", title, orderCategoryId)
            // log.Printf("endpoint_get_orders: createdAt: %d", createdAt)
            // log.Printf("endpoint_get_orders: updatedAt: %d", updatedAt)
            createdAtISO := time.Unix(0, createdAtMicroseconds*1000).UTC().Format(time.RFC3339Nano)
            updatedAtISO := time.Unix(0, updatedAtMicroseconds*1000).UTC().Format(time.RFC3339Nano)
            orders = append(orders, entity.Order{
                Id: id,
                Title: title,
                OrderCategoryId: orderCategoryId,
                CreatedAtMicroseconds: createdAtMicroseconds,
                CreatedAtISO: createdAtISO,
                UpdatedAtMicroseconds: updatedAtMicroseconds,
                UpdatedAtISO: updatedAtISO,
            })
        }

        return ctx.JSON(http.StatusCreated, ResponseJson{
            Count: len(orders),
            Orders: orders,
        })
    }
}


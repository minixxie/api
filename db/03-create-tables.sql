DROP TABLE IF EXISTS "OrderCategory";
CREATE TABLE if not exists "OrderCategory" (
    "id" SERIAL PRIMARY KEY,
    "key" CHAR(20) NOT NULL,
	"title_en_US" varchar(255),
	"title_zh_CN" varchar(255),
	"title_zh_HK" varchar(255)
);
INSERT INTO "OrderCategory" ("key", "title_en_US", "title_zh_CN", "title_zh_HK") values ('NONE', 'None', '无', '無');
INSERT INTO "OrderCategory" ("key", "title_en_US", "title_zh_CN", "title_zh_HK") values ('FOOD_DELIVERY', 'Food Delivery', '送外卖', '送外賣');
INSERT INTO "OrderCategory" ("key", "title_en_US", "title_zh_CN", "title_zh_HK") values ('COURIER', 'Courier', '信使', '信使');

DROP TABLE IF EXISTS "Order";
CREATE TABLE if not exists "Order" (
    "id" SERIAL PRIMARY KEY,
    "title" varchar(255) NOT NULL DEFAULT '',
    "orderCategoryId" bigint NOT NULL,
    "createdAtMicroseconds" bigint NOT NULL,
    "updatedAtMicroseconds" bigint NOT NULL,
    foreign key ("orderCategoryId") references "OrderCategory" ("id")
);


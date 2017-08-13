DROP TABLE IF EXISTS "OrderCategory";
CREATE TABLE if not exists "OrderCategory" (
    "id" SERIAL PRIMARY KEY,
    "key" VARCHAR(20) NOT NULL,
	"title_en_US" VARCHAR(255),
	"title_zh_CN" VARCHAR(255),
	"title_zh_HK" VARCHAR(255)
);
INSERT INTO "OrderCategory" ("key", "title_en_US", "title_zh_CN", "title_zh_HK") values ('NONE', 'None', '无', '無');
INSERT INTO "OrderCategory" ("key", "title_en_US", "title_zh_CN", "title_zh_HK") values ('FOOD_DELIVERY', 'Food Delivery', '送外卖', '送外賣');
INSERT INTO "OrderCategory" ("key", "title_en_US", "title_zh_CN", "title_zh_HK") values ('COURIER', 'Courier', '信使', '信使');

DROP TABLE IF EXISTS "Order";
CREATE TABLE if not exists "Order" (
    "id" SERIAL PRIMARY KEY,
    "title" VARCHAR(255) NOT NULL DEFAULT '',
    "orderCategoryId" BIGINT NOT NULL,
    "createdAtMicroseconds" BIGINT NOT NULL,
    "updatedAtMicroseconds" BIGINT NOT NULL,
    foreign key ("orderCategoryId") references "OrderCategory" ("id")
);


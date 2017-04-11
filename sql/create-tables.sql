CREATE TABLE if not exists Orders (
    id binary(16) NOT NULL,
    title varchar(255) NOT NULL DEFAULT '',
    orderCategoryId CHAR(20) NOT NULL DEFAULT 'NONE',
    createdAt bigint NOT NULL,
    updatedAt bigint NOT NULL,
    primary key(id),
    foreign key (orderCategoryId) references OrderCategory (id)
);

CREATE OR REPLACE VIEW OrdersView AS
	SELECT lower(hex(id)) as idHex, title, orderCategoryId,
	from_unixtime(createdAt/1000000000) as createdAtISO,
	from_unixtime(updatedAt/1000000000) as updatedAtISO
	FROM Orders ORDER BY createdAt DESC;

CREATE TABLE if not exists OrderCategory (
	id CHAR(20) NOT NULL,
	title_en_US varchar(255),
	title_zh_CN varchar(255),
	title_zh_HK varchar(255),
    primary key(id)
);
INSERT INTO OrderCategory (id, title_en_US, title_zh_CN, title_zh_HK) values ('NONE', 'None', '无', '無');
INSERT INTO OrderCategory (id, title_en_US, title_zh_CN, title_zh_HK) values ('FOOD_DELIVERY', 'Food Delivery', '送外卖', '送外賣');
INSERT INTO OrderCategory (id, title_en_US, title_zh_CN, title_zh_HK) values ('COURIER', 'Courier', '信使', '信使');

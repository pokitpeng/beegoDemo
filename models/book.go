package models

// 书本表
type Book struct {
	Id       int    `json:"id"`        // 书id
	Name     string `json:"name"`      // 书名
	AuthorId int    `json:"author_id"` // 作者 外键
	// 一对一 一本书对应一个作者
	Author Author `gorm:"ForeignKey:AuthorId;AssociationForeignKey:Id"` // AuthorId关联teacher表中的Id
}

// TableName ...
func (Book) TableName() string {
	return "book"
}

/* 测试数据
INSERT INTO book (name,author_id)VALUES('遮天',1);
INSERT INTO book (name,author_id)VALUES('不死不灭',1);
INSERT INTO book (name,author_id)VALUES('神墓',1);
INSERT INTO book (name,author_id)VALUES('长生界',1);
INSERT INTO book (name,author_id)VALUES('完美世界',1);
INSERT INTO book (name,author_id)VALUES('仙逆',2);
INSERT INTO book (name,author_id)VALUES('求魔',2);
INSERT INTO book (name,author_id)VALUES('一念永恒',2);
INSERT INTO book (name,author_id)VALUES('三寸人间',2);
INSERT INTO book (name,author_id)VALUES('凡人修仙传',3);
*/

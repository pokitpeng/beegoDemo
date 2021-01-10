package models

/*
作者表
规定：一个作者可以写多本书，但是一本书只有一个作者
*/
type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Sex  string `json:"sex"`
	Age  int    `json:"age"`
	// 	一对多 一个作者有多本书
	// Books []Book `gorm:"ForeignKey:AuthorId;AssociationForeignKey:Id"` // 外键配置与book表保持一致
	Books []Book // 实际测试发现外键关联可以不用再写
}

// TableName ...
func (Author) TableName() string {
	return "author"
}

/* 测试数据
INSERT INTO author (id,name,sex,age)VALUES(1,'辰东','男',38);
INSERT INTO author (id,name,sex,age)VALUES(2,'耳根','男',35);
INSERT INTO author (id,name,sex,age)VALUES(3,'忘语','男',30);
*/

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

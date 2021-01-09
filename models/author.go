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
	// Books []Book `gorm:"ForeignKey:AuthorId;AssociationForeignKey:Id"` // 外键配置与之前保持一致
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

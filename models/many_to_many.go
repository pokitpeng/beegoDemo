package models

/*
学生与课程，多对多关系，一个学生学多门课，一门课有多个学生
*/

// 学生表
type Student struct {
	Id      int       `json:"id"`
	Name    string    `json:"name"`
	Age     int       `json:"age"`
	Lessons []*Lesson `json:"lessons" gorm:"many2many:student_lesson;"`
}

// TableName ...
func (Student) TableName() string {
	return "student"
}

/*
测试数据
INSERT INTO student VALUES(1,'tom',18);
INSERT INTO student VALUES(2,'tony',18);
INSERT INTO student VALUES(3,'anny',19);
INSERT INTO student VALUES(4,'pokit',18);
INSERT INTO student VALUES(5,'tiny',19);
INSERT INTO student VALUES(6,'jerry',18);
*/

// 课程表
type Lesson struct {
	Id       int        `json:"id"`
	Name     string     `json:"name"`
	Students []*Student `json:"students" gorm:"many2many:student_lesson;"`
}

// TableName ...
func (Lesson) TableName() string {
	return "lesson"
}

/*
测试数据
INSERT INTO lesson VALUES(1,'golang');
INSERT INTO lesson VALUES(2,'python');
INSERT INTO lesson VALUES(3,'java');
INSERT INTO lesson VALUES(4,'js');
INSERT INTO lesson VALUES(5,'rust');
INSERT INTO lesson VALUES(6,'c');
*/

/*// 学生 课程中间表 默认会自动创建并关联
type StudentLesson struct {
	StudentId int       `json:"student_id"`
	LessonId  int       `json:"lesson_id"`

}

// TableName ...
func (StudentLesson) TableName() string {
	return "student_lesson"
}*/

/*
测试数据
INSERT INTO student_lesson (student_id,lesson_id)VALUES(1,1);
INSERT INTO student_lesson (student_id,lesson_id)VALUES(1,2);
INSERT INTO student_lesson (student_id,lesson_id)VALUES(1,3);
INSERT INTO student_lesson (student_id,lesson_id)VALUES(2,4);
INSERT INTO student_lesson (student_id,lesson_id)VALUES(2,5);
INSERT INTO student_lesson (student_id,lesson_id)VALUES(2,6);
INSERT INTO student_lesson (student_id,lesson_id)VALUES(3,2);
INSERT INTO student_lesson (student_id,lesson_id)VALUES(3,4);
INSERT INTO student_lesson (student_id,lesson_id)VALUES(3,6);
INSERT INTO student_lesson (student_id,lesson_id)VALUES(4,1);
INSERT INTO student_lesson (student_id,lesson_id)VALUES(4,3);
INSERT INTO student_lesson (student_id,lesson_id)VALUES(4,5);
*/

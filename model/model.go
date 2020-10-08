package model

type Answer struct {
	Id         int    `orm:"id" json:"id"`
	Multi      int    `orm:"multi" json:"multi"`
	QuestionId int    `orm:"question_id" json:"question_id"`
	OptionA    string `orm:"option_a" json:"option_a"`
	OptionB    string `orm:"option_b" json:"option_b"`
	OptionC    string `orm:"option_c" json:"option_c"`
	OptionD    string `orm:"option_d" json:"option_d"`
	Rights     string `orm:"rights" json:"rights"`
}

type Config struct {
	Key1   string `orm:"key1" json:"key1"`
	Key2   string `orm:"key2" json:"key2"`
	Value1 string `orm:"value1" json:"value1"`
	Value2 string `orm:"value2" json:"value2"`
}

type Note struct {
	Id        int    `orm:"id" json:"id"`
	Uid       int    `orm:"uid" json:"uid"`
	Source    int    `orm:"source" json:"source"`
	Data      string `orm:"data" json:"data"`
	Tag       string `orm:"tag" json:"tag"`
	CreatedAt string `orm:"created_at" json:"created_at"`
	UpdatedAt string `orm:"updated_at" json:"updated_at"`
	Kind      int    `orm:"kind" json:"kind"`
}

type Question struct {
	Id         int    `orm:"id" json:"id"`
	CategoryI  int    `orm:"category_i" json:"category_i"`
	CategoryII int    `orm:"category_i_i" json:"category_i_i"`
	Question   string `orm:"question" json:"question"`
}

func (q Question) GetByIdDefine(id int) {
}

func (q Question) GetById(id int) error {
	err := MyDB.QueryRow("select id from question where id = ?", 1).Scan(&q.Id)
	return err
}

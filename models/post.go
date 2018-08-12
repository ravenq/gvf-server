package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

// PostType the post type
type PostType int

const (
	Original PostType = iota
	Repost
	Translated
)

// PostStatus the post status
type PostStatus int

const (
	Normal = iota
	Draft
	Deleted
)

// Post model.
type Post struct {
	Id         int64     `orm:"auto;unique;pk" json:"id,omitempty"`
	Author     string    `json:"author,omitempty"`
	Category   *Category `orm:"rel(fk)" json:"category,omitempty"`
	Title      string    `json:"title,omitempty"`
	Tags       string    `json:"tags,omitempty"`
	IsTop      bool      `json:"is_top,omitempty"`
	Summary    string    `json:"summary,omitempty"`
	Content    string    `json:"content,omitempty"`
	Vist       int       `json:"vist,omitempty"`
	Status     int       `json:"status,omitempty"`
	PostType   PostType  `json:"post_type,omitempty"`
	RefUrl     string    `orm:"null" json:"refRrl,omitempty"`
	RefAuthor  string    `orm:"null" json:"refAuthor,omitempty"`
	Translator string    `orm:"null" json:"translator,omitempty"`
	CreateTime time.Time `orm:"auto_now_add;type(datetime)" json:"createTime,omitempty"`
	UpdateTime time.Time `orm:"auto_now;type(datetime)" json:"updateTime,omitempty"`
	PubTime    time.Time `orm:"auto_now;type(datetime)" json:"pubTime,omitempty"`
}

func init() {
	orm.RegisterModel(new(Post))
}

// AddPost insert a new Post into database and returns
// last inserted Id on success.
func AddPost(m *Post) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetPostById retrieves Post by Id. Returns error if
// Id doesn't exist
func GetPostById(id int64) (v *Post, err error) {
	o := orm.NewOrm()
	v = &Post{Id: id}
	if err = o.QueryTable(new(Post)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllPost retrieves all Post matches certain condition. Returns empty list if
// no records exist
func GetAllPost(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Post))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Post
	qs = qs.OrderBy(sortFields...).RelatedSel()
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdatePost updates Post by Id and returns error if
// the record to be updated doesn't exist
func UpdatePostById(m *Post) (err error) {
	o := orm.NewOrm()
	v := Post{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeletePost deletes Post by Id and returns error if
// the record to be deleted doesn't exist
func DeletePost(id int64) (err error) {
	o := orm.NewOrm()
	v := Post{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Post{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

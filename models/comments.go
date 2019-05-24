package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

// CommentsStatus the post status
type CommentsStatus int

const (
	// CommentsStatusNormal nomal
	CommentsStatusNormal CommentsStatus = iota
	// CommentsStatusDeleted comment soft deleted
	CommentsStatusDeleted
	// CommentsStatusHide comment has been hided
	CommentsStatusHide
)

// Comments model.
type Comments struct {
	Id         int64          `orm:"auto;unique;pk" json:"id,omitempty"`
	IsTop      bool           `json:"isTop,omitempty"`
	Parent     int64          `json:"parent,omitempty"`
	CommentId  string         `json:"commentId"`
	Author     *User          `orm:"rel(fk);null;on_delete(set_null);" json:"author,omitempty"`
	CreateTime time.Time      `orm:"auto_now_add;type(datetime)" json:"createTime,omitempty"`
	UpdateTime time.Time      `orm:"auto_now_add;type(datetime)" json:"updateTime,omitempty"`
	Status     CommentsStatus `json:"status,omitempty"`
	RemoteAddr string         `json:"remoteAddr,omitempty"`
	Content    string         `orm:"type(text)" json:"content,omitempty"`
	Likes      int64     	    `json:"likes"`
	Dislikes   int64          `json:"dislikes"`
	Replies    []*Comments    `orm:"-" json:"replies,omitempty"`
}

func init() {
	orm.RegisterModel(new(Comments))
}

// AddComments insert a new Comments into database and returns
// last inserted Id on success.
func AddComments(m *Comments) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetCommentsById retrieves Comments by Id. Returns error if
// Id doesn't exist
func GetCommentsById(id int64) (v *Comments, err error) {
	o := orm.NewOrm()
	v = &Comments{Id: id}
	if err = o.QueryTable(new(Comments)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// UpdateCommentsById updates Comments by Id and returns error if
// the record to be updated doesn't exist
func UpdateCommentsById(m *Comments) (err error) {
	o := orm.NewOrm()
	v := Comments{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		v.UpdateTime = time.Now()
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteComments deletes Comments by Id and returns error if
// the record to be deleted doesn't exist
func DeleteComments(id int64) (err error) {
	o := orm.NewOrm()
	v := Comments{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		v.Status = CommentsStatusDeleted
		if num, err = o.Update(v); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// GetCommentsByCommentId get Comments by post id
func GetCommentsByCommentId(id string) (ml []Comments, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable(new(Comments)).Filter("comment_id", id).OrderBy("create_time").RelatedSel().All(&ml)
	return
}

// GetCommentsCount get post message count.
func GetCommentsCount(id string) (count int64, err error) {
	o := orm.NewOrm()
	count, err = o.QueryTable(new(Comments)).Filter("comment_id", id).RelatedSel().Count()
	return
}

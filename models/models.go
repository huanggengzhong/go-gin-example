package models

import (
	"fmt"
	"log"
	"time"

	"github.com/huanggengzhong/go-gin-example/pkg/setting"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
	DeletedOn  int `json:"deleted_on"`
}

func init() {
	var (
		err                                               error
		dbType, dbName, user, password, host, tablePrefix string
	)

	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}

	dbType = sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()
	// 记得不需要:
	// 原因:
	// 这导致了一个新的db变量被创建，而不是更新外部作用域中的全局db变量。因此，在CloseDB()函数中使用的db实际上是一个不同的未初始化变量。要修复这个问题，你需要将:=替换为=，以便使用全局的db变量进行赋值：

	// db, err := gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))
	if err != nil {
		log.Println(err)
	}
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}
	db.SingularTable(true)
	//todo callback
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	db.Callback().Delete().Replace("gorm:delete", deleteCallback)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(10)

}

func CloseDB() {
	defer db.Close()
}

// 新增gorm回调
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		//scope.FieldByName 通过 scope.Fields() 获取所有字段
		if createTimeField, ok := scope.FieldByName("CreatedOn"); ok {
			// field.IsBlank 可判断该字段的值是否为空
			if createTimeField.IsBlank {
				createTimeField.Set(nowTime)
			}
		}

		if modifyTimeField, ok := scope.FieldByName("ModifiedOn"); ok {
			if modifyTimeField.IsBlank {
				modifyTimeField.Set(nowTime)
			}
		}

	}
}

func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	// 	scope.Get(...) 根据入参获取设置了字面值的参数，例如本文中是 gorm:update_column ，它会去查找含这个字面值的字段属性
	// scope.SetColumn(...) 假设没有指定 update_column 的字段，我们默认在更新回调设置 ModifiedOn 的值
	if _, ok := scope.Get("gorm:update_column"); !ok {
		scope.SetColumn("ModifiedOn", time.Now().Unix())
	}
}

// scope.Get("gorm:delete_option") 检查是否手动指定了 delete_option
// scope.FieldByName("DeletedOn") 获取我们约定的删除字段，若存在则 UPDATE 软删除，若不存在则 DELETE 硬删除
// scope.QuotedTableName() 返回引用的表名，这个方法 GORM 会根据自身逻辑对表名进行一些处理
// scope.CombinedConditionSql() 返回组合好的条件 SQL，看一下方法原型很明了
func deleteCallback(scope *gorm.Scope) {

	if !scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}
		deletedOnField, hasDeletedOnFied := scope.FieldByName("DeletedOn")
		if !scope.Search.Unscoped && hasDeletedOnFied {
			scope.Raw(fmt.Sprintf(
				"UPDATE %v set %v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deletedOnField.DBName),
				scope.AddToVars(time.Now().Unix()),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else {
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			))
		}
	}
}

func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}

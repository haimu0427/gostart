package main

import (
	_ "context"

	"gorm.io/driver/mysql"
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type UserInfo struct {
	gorm.Model        // id 会自动设置为主键
	Name       string `gorm:"type:varchar(100);unique_index"`
	Gender     string `gorm:"column:sex_of_user"`
	Hobby      string
}

func (UserInfo) TableName() string {
	return "profiles"
}

func main() {
	// GORM v2 已不支持 DefaultTableNameHandler，相关代码可删除或注释
	// gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	// 	return "tbl_" + defaultTableName
	// }
	// 	return "tbl_" + defaultTableName
	// }
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:root@tcp(120.26.105.128:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// GORM v2 不再支持 db.SingularTable(true)，表名复数化需通过自定义 TableName 方法实现
	// 如果需要禁用表名复数化，可以为模型添加 TableName 方法：
	// func (UserInfo) TableName() string { return "user_info" }
	// 自动迁移表结构
	err = db.AutoMigrate(&UserInfo{})
	if err != nil {
		panic("failed to migrate database: " + err.Error())
	}

	// 示例：创建一条记录（不指定ID，让数据库自动生成）
	user := UserInfo{Name: "alice", Gender: "female", Hobby: "rap"}
	result := db.Create(&user)
	if result.Error != nil {
		panic("failed to create user: " + result.Error.Error())
	}
	println("成功创建用户,ID:", user.ID)

	var u UserInfo
	result = db.First(&u, user.ID) // 查询刚创建的记录
	if result.Error != nil {
		panic("failed to find user: " + result.Error.Error())
	}
	println("查询到用户:", u.Name)

	result = db.Model(&u).Update("Hobby", "双色球")
	if result.Error != nil {
		panic("failed to update user: " + result.Error.Error())
	}
	println("更新后的爱好:", u.Hobby)

	// 注释掉删除操作，这样数据就会保留在数据库中
	// result = db.Delete(&u)
	// if result.Error != nil {
	// 	panic("failed to delete user: " + result.Error.Error())
	// }
	// println("用户已删除")
	//删除使用的都是一些软删除, 不会真真的删除, 而是标记删除
	// 例如, gorm.Model 中有一个 DeletedAt 字段, 记录删除时间, 如果该字段不为 null, 则表示已删除
	// 这样可以避免误删数据, 也方便数据恢复
	// 如果确实需要物理删除, 可以使用 Unscoped 方法, 例如:
	// db.Unscoped().Delete(&u) // 这将真正删除记录
	// 关闭数据库连接

}

//  func main() {
// 	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
// 	if err != nil {
// 		panic("failed to connect database")
// 	}

// 	ctx := context.Background()

// 	// Migrate the schema
// 	db.AutoMigrate(&Product{})

// 	// Create
// 	err = gorm.G[Product](db).Create(ctx, &Product{Code: "D42", Price: 100})

// 	// Read
// 	var product Product
// 	err = db.WithContext(ctx).First(&product, 1).Error // find product with integer primary key

// 	var products []Product
// 	err = db.WithContext(ctx).Where("code = ?", "D42").Find(&products).Error // find products with code D42

// 	// Update - update product's price to 200
// 	err = db.WithContext(ctx).Model(&product).Update("Price", 200).Error

// 	// Update - update multiple fields
// 	err = db.WithContext(ctx).Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"}).Error

// 	// Delete - delete product
// 	err = db.WithContext(ctx).Delete(&product).Error
// }

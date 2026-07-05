package dbhelper

import (
	"database/sql"
	"fmt"

	configpkg "github.com/Hungnd562k/GoRestApiTutorial/config_pkg"
	_ "github.com/lib/pq"
)

func CheckDbConnection() bool {
	// 1. Tạo chuỗi kết nối chuẩn từ cấu hình (Nên lấy từ configpkg thay vì hardcode)
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable connect_timeout=5",
		configpkg.Host, 5432, configpkg.Username, configpkg.Password, "postgres")

	// 2. Mở kết nối đến DB
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		fmt.Println("Lỗi mở kết nối:", err)
		return false
	}
	defer db.Close()

	// 3. Thực sự ping đến DB để kiểm tra kết nối sống/chết (sql.Open không tự ping)
	err = db.Ping()
	if err != nil {
		fmt.Println("Không thể ping tới Database:", err)
		return false
	}

	fmt.Println("Kết nối cơ sở dữ liệu thành công!")
	return true
}

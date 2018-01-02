package data

import (
	"crypto/rand"
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"
	// database/sql 을 위한 postgres 드라이버.
	// 왜 별명을 언더바로 하는 지 모르겠다.
	// 인데 오류를 피하려고 하는 듯?
	// https://stackoverflow.com/questions/21743841/how-to-avoid-annoying-error-declared-and-not-used-from-golang
	_ "github.com/lib/pq"
)

// db 변수에 타입 초기화
var DB *sql.DB

// 데이터 베이스 접속 생성
func init() {

	fmt.Println("DB init()")
	var err error
	db, err := sql.Open("postgres", "user=postgres dbname=chitchat password=postgres sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}
	// 정의됐지만 사용되지 않음 오류를 피하기 위해 사용을 하고..
	/// 동시에 접속에 오류가 없는지 확인하기 위해서 사용한다.
	// error 를 반환하므로 결과값이 nil 이어야 접속이 잘 된 것이다.
	if db.Ping() == nil {
		fmt.Println("sql.Open 작업이 잘 수행되었습니다.")
	}

	DB = db

	return
}

// 랜덤한 UUID 생성
// http://github.com/nu7hatch/gouuid
func UUID생성() (uuid string) {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		log.Fatalln("UUID 못 만들겠다", err)
	}

	// 0x40 은 RFC 4122 에서 예약된 변수다. -> 잘 모르는 이야기
	u[8] = (u[8] | 0x40) & 0x7F
	// Set the four most significant bits (bits 12 through 15) of the
	// time_hi_and_version field to the 4-bit version number.
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return
}

// hash plaintext with SHA-1
func Crypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return
}

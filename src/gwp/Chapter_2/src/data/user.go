package data

import (
	"fmt"
	"time"
)

type User struct {
	Id        int
	Uuid      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

type Session struct {
	Id        int
	Uuid      string
	Email     string
	UserId    int
	CreatedAt time.Time
}

func (user *User) CreateSession() (session Session, err error) {
	// 세선 테이블에 사용자의 정보를 넣는다.
	문 := "insert into sessions (uuid, email, user_id, created_at) values ($1,$2,$3,$4) returning id,uuid,email,user_id,created_at"
	접속문, 오류 := DB.Prepare(문)
	if 오류 != nil {
		return
	}
	defer 접속문.Close()
	// 위에서 생성한 오류타입의 오류변수를 재활용
	오류 = 접속문.QueryRow(UUID생성(), user.Email, user.Id, time.Now()).Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	return
}

func (user *User) Create() (err error) {

	// postgres 는 마지막으로 삽입된 id f르 자동으로 반환하지 않는다.
	// 왜냐면 그런 가정이 틀렸기 때문이다.
	// 항상 시퀸스를 사용해야 한다. RETURNING 키워드를 사용해서 이 정보를 가져온다.

	문 := "insert into users(uuid, name, email, password, created_at) values($1, $2, $3, $4, $5) returning id, uuid, created_at"
	fmt.Println("문 실행 전")
	접속문, 오류 := DB.Prepare(문)
	fmt.Println("문 실행 됨")
	if 오류 != nil {
		return
	}
	defer 접속문.Close()

	// 커넥션을 가지고 있는 접속문에다가 생성된 UUID, 이름, 이메일, 암호, 현재 시간을 조회한 다음에...
	// 사용자 구조체에다가 반환된 id 를 스캔한다.
	// values($1, $2, $3, $4, $5) 가 각각 아래의 쿼리 로우세 변수로 치환되게 된다.
	오류 = 접속문.QueryRow(UUID생성(), user.Name, user.Email, Crypt(user.Password), time.Now()).Scan(&user.Name, &user.Uuid, &user.CreatedAt)
	// scan 문은 결과가 한 줄일 때 직접 결과를 변수에 담아버리는 것이다.

	return

}

func UserByEmail(이메일 string) (user User, 오류 error) {
	user = User{}
	fmt.Println("이메일로 사용자 가져오기")
	오류 = DB.QueryRow("SELECT id, uuid, name, email, password, created_at FROM users WHERE email = $1", 이메일).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return
}

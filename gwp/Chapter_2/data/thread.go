package data

import(
	"time"
	"fmt"
)

type Thread struct {
	Id	int
	Uuid string
	Topic	string
	UserId	int
	CreatedAt	time.Time
}

func Threads() (쓰레드들 []Thread, 오류 error){

	fmt.Println("쓰레드 가져오기 진입")
	행들, 오류 := DB.Query("SELECT id, uuid, topic, user_id, created_at FROM threads ORDER BY created_at DESC")
	if 오류 != nil {
		return
	}
	for 행들.Next() {
		conv := Thread{}
		if  오류 = 행들.Scan(&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId, &conv.CreatedAt); 오류 != nil {
			return
		}
		쓰레드들 = append(쓰레드들, conv)
	}
	행들.Close()
	return
}

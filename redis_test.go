package mqproxy

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/go-redis/redis"
)

type news struct {
	Title, Content string
}

func (s *news) String() string {
	return fmt.Sprintf("标题:%s\n内容:%s", s.Title, s.Content)
}

func TestConsume(t *testing.T) {

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// type assign interface
	var mqClient MqClient
	mqClient = &RedisMq{redisCli: redisClient}

	mqClient.consume("news", "consumer", "group", func(data interface{}) error {
		n := news{}
		json.Unmarshal([]byte(data.(string)), &n)
		fmt.Println(&n)
		return nil
	}, func(err error) {
		t.Error("有错了" + err.Error())
	})

}

func TestProduce(t *testing.T) {

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// type assign interface
	var mqserver MqServer
	mqserver = &RedisMq{redisCli: redisClient}

	n := &news{
		Title:   "今天早报",
		Content: "俄罗斯与乌克兰纠纷日益严重，国际社会开始对此表示高度关注",
	}
	b, _ := json.Marshal(n)
	if err := mqserver.produce("news", string(b)); err != nil {
		t.Error(err)
	}

}

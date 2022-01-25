package mqproxy

import (
	"fmt"
	"testing"

	"github.com/go-redis/redis"
)

func TestConsume(t *testing.T) {

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// type assign interface
	var mqClient MqClient
	mqClient = &RedisMq{redisCli: redisClient}

	for {

		mqClient.consume("news", "", "", func(data interface{}) error {
			fmt.Println("今天的新闻是", data.(string))
			return nil
		})

	}

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

	if err := mqserver.produce("news", "俄罗斯与乌克兰纠纷日益严重，国际社会开始对此表示担忧"); err != nil {
		t.Error(err)
	}

}

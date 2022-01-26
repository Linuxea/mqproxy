package mqproxy

import (
	"github.com/go-redis/redis"
)

type RedisMq struct {
	redisCli *redis.Client
}

func (r *RedisMq) produce(dest string, data interface{}) error {
	return r.redisCli.Publish(dest, data).Err()
}

func (r *RedisMq) require(dest, name, group string) (interface{}, error) {
	stream := dest
	subscribe := r.redisCli.Subscribe(stream)
	return (<-subscribe.Channel()).Payload, nil
}

func (r *RedisMq) consume(dest, name, group string, bf businessFunc, eh errHandler) {

	for {
		var err error
		b, err := r.require(dest, name, group)
		if err != nil {
			eh(err)
			return
		}

		if err = bf(b); err != nil {
			eh(err)
		}
	}
}

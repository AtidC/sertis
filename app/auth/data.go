package auth

import (
	"blog/db"
	"blog/log"
	"context"
	"time"

	config "github.com/spf13/viper"
)

var dtFormat = config.GetString("datetime.format")

func selectUserData(requestId, userName string) (user, error) {
	startProcess := time.Now()
	res := user{}

	timeout := config.GetDuration("db.postgres.timeout")
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Minute)
	defer cancel()

	conn := db.GetPostgresPool()
	rows, err := conn.Query(ctx, SQLSelectUserInfo, userName)
	if err != nil {
		log.End(requestId, startProcess, err)
		return res, err
	}
	defer rows.Close()

	if rows.Next() {
		values, err := rows.Values()
		if err != nil {
			log.End(requestId, startProcess, err)
			return res, err
		}

		res = user{
			ID:   values[0].(string),
			Name: values[1].(string),
		}
	}
	log.Info(requestId, "return: %v", res)

	log.End(requestId, startProcess, nil)
	return res, nil
}

func checkPassWord(requestId, userName, passWord string) (bool, error) {
	startProcess := time.Now()
	valid := false

	timeout := config.GetDuration("db.postgres.timeout")
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Minute)
	defer cancel()

	conn := db.GetPostgresPool()
	rows, err := conn.Query(ctx, SQLSelectPassOfUser, userName)
	if err != nil {
		log.End(requestId, startProcess, err)
		return valid, err
	}
	defer rows.Close()

	if rows.Next() {
		values, err := rows.Values()
		if err != nil {
			log.End(requestId, startProcess, err)
			return valid, err
		}

		if passWord == values[0].(string) {
			valid = true
		}

	}
	log.Info(requestId, "return: %v", valid)

	log.End(requestId, startProcess, nil)
	return valid, nil
}

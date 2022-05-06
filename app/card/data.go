package card

import (
	"blog/db"
	"blog/log"
	"context"
	"fmt"
	"strings"
	"time"

	config "github.com/spf13/viper"
)

func insertCard(requestId string, req CardRequest) error {
	startProcess := time.Now()

	// Create context for set service timeout
	timeout := config.GetDuration("db.postgres.timeout")
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()

	conn := db.GetPostgresPool()
	if _, err := conn.Exec(ctx, SQLInsertCard,
		requestId,
		req.Category,
		req.Title,
		req.Status,
		req.Content,
		req.AuthorID,
	); err != nil {
		log.End(requestId, startProcess, err)
		return err
	}

	log.End(requestId, startProcess, nil)
	return nil
}

func updateCard(requestId string, req CardRequest) error {
	startProcess := time.Now()

	// Create context for set service timeout
	timeout := config.GetDuration("db.postgres.timeout")
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()

	col := ""
	if req.Category != "" {
		col += fmt.Sprintf("category = '%s',", req.Category)
	}
	if req.Title != "" {
		col += fmt.Sprintf("title = '%s',", req.Title)
	}
	if req.Status != "" {
		col += fmt.Sprintf("status = '%s',", req.Status)
	}
	if req.Content != "" {
		col += fmt.Sprintf("content = '%s',", req.Content)
	}
	if req.AuthorID != "" {
		col += fmt.Sprintf("author = '%s'", req.AuthorID)
	}
	sql := strings.Replace(SQLUpdateCard, "{column}", col, 1)
	log.Debug(requestId, "sql: %s", sql)

	conn := db.GetPostgresPool()
	if _, err := conn.Exec(ctx, sql, req.Id); err != nil {
		log.End(requestId, startProcess, err)
		return err
	}

	log.End(requestId, startProcess, nil)
	return nil
}

func deleteCard(requestId, cardId string) error {
	startProcess := time.Now()

	// Create context for set service timeout
	timeout := config.GetDuration("db.postgres.timeout")
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()

	conn := db.GetPostgresPool()
	if _, err := conn.Exec(ctx, SQLDeleteCard,
		cardId,
	); err != nil {
		log.End(requestId, startProcess, err)
		return err
	}

	log.End(requestId, startProcess, nil)
	return nil
}

func selectCardList(requestId string) ([]Card, error) {
	startProcess := time.Now()
	res := make([]Card, 0)

	timeout := config.GetDuration("db.postgres.timeout")
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Minute)
	defer cancel()

	conn := db.GetPostgresPool()
	rows, err := conn.Query(ctx, SQLSelectCardList)
	if err != nil {
		log.End(requestId, startProcess, err)
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			log.End(requestId, startProcess, err)
			return res, err
		}

		res = append(res, Card{
			//convert DB types to Go types
			Id:         values[0].(string),
			Category:   values[1].(string),
			Title:      values[2].(string),
			Status:     values[3].(string),
			Content:    db.NullString(values[4]),
			Author:     values[5].(string),
			CreateDate: values[6].(time.Time).Format(config.GetString("datetime.format")),
			UpdateDate: values[7].(time.Time).Format(config.GetString("datetime.format")),
		})
		fmt.Println(values[6].(time.Time))
	}
	log.Info(requestId, "return: %v", res)

	log.End(requestId, startProcess, nil)
	return res, nil
}

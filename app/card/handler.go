package card

import (
	"fmt"
	"net/http"
	"time"

	"blog/app/auth"
	"blog/log"

	"github.com/google/uuid"
	echo "github.com/labstack/echo/v4"
)

func Create(c echo.Context) error {
	startProcess := time.Now()
	requestId := uuid.New().String()
	log.StartAPI(requestId, c)

	// Bind context to struct
	req := CardRequest{}
	if err := c.Bind(&req); err != nil {
		log.EndAPI(requestId, startProcess, http.StatusBadRequest, err.Error())
		return c.JSON(http.StatusBadRequest, Response{
			Code:      "E000",
			Message:   fmt.Sprintf("Invalid request body, %s", err.Error()),
			RequestId: requestId,
		})
	}
	req.AuthorID = auth.TokenInfo(c)

	// Insert new card to database
	if err := insertCard(requestId, req); err != nil {
		log.EndAPI(requestId, startProcess, http.StatusInternalServerError, err.Error())
		return c.JSON(http.StatusInternalServerError, Response{
			Code:      "E100",
			Message:   fmt.Sprintf("Can't create card, %s", err.Error()),
			RequestId: requestId,
		})
	}

	log.EndAPI(requestId, startProcess, http.StatusOK, "")
	return c.JSON(http.StatusOK, Response{
		Code:      "S000",
		Message:   "Success",
		RequestId: requestId,
	})
}

func Update(c echo.Context) error {
	startProcess := time.Now()
	requestId := uuid.New().String()
	log.StartAPI(requestId, c)

	// Bind context to struct
	req := CardRequest{}
	if err := c.Bind(&req); err != nil {
		log.EndAPI(requestId, startProcess, http.StatusBadRequest, err.Error())
		return c.JSON(http.StatusBadRequest, Response{
			Code:      "E000",
			Message:   fmt.Sprintf("Invalid request body, %s", err.Error()),
			RequestId: requestId,
		})
	}
	req.AuthorID = auth.TokenInfo(c)

	if err := updateCard(requestId, req); err != nil {
		log.EndAPI(requestId, startProcess, http.StatusInternalServerError, err.Error())
		return c.JSON(http.StatusInternalServerError, Response{
			Code:      "E100",
			Message:   fmt.Sprintf("Can't update card, %s", err.Error()),
			RequestId: requestId,
		})
	}

	log.EndAPI(requestId, startProcess, http.StatusOK, "")
	return c.JSON(http.StatusOK, Response{
		Code:      "S000",
		Message:   "Success",
		RequestId: requestId,
	})
}

func Delete(c echo.Context) error {
	startProcess := time.Now()
	requestId := uuid.New().String()
	log.StartAPI(requestId, c)

	if err := deleteCard(requestId, c.Param("cardId")); err != nil {
		log.EndAPI(requestId, startProcess, http.StatusInternalServerError, err.Error())
		return c.JSON(http.StatusInternalServerError, Response{
			Code:      "E100",
			Message:   fmt.Sprintf("Can't delete card, %s", err.Error()),
			RequestId: requestId,
		})
	}

	log.EndAPI(requestId, startProcess, http.StatusOK, "")
	return c.JSON(http.StatusOK, Response{
		Code:      "S000",
		Message:   "Success",
		RequestId: requestId,
	})
}

func List(c echo.Context) error {
	startProcess := time.Now()
	requestId := uuid.New().String()
	log.StartAPI(requestId, c)

	// Select card on database
	res, err := selectCardList(requestId)
	if err != nil {
		log.EndAPI(requestId, startProcess, http.StatusInternalServerError, err.Error())
		return c.JSON(http.StatusInternalServerError, Response{
			Code:      "E001",
			Message:   fmt.Sprintf("Can't get card list, %s", err.Error()),
			RequestId: requestId,
		})
	}
	if len(res) == 0 {
		log.EndAPI(requestId, startProcess, http.StatusOK, "Data not found")
		return c.JSON(http.StatusOK, Response{
			Code:      "S001",
			Message:   "Data not found",
			RequestId: requestId,
		})
	}

	log.EndAPI(requestId, startProcess, http.StatusOK, res)
	return c.JSON(http.StatusOK, Response{
		Code:      "S000",
		Message:   "Success",
		RequestId: requestId,
		Result:    res,
	})
}

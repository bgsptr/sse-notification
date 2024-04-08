package handler

import (
	// "context"
	// "errors"
	// "log"
	// "time"
	// "encoding/json"
	"fmt"
	// "net/http"
	"time"

	// "strings"

	"github.com/labstack/echo/v4"
)

var (
    BASE_URL = "http://localhost:8000/api"
)

func EventHandler(c echo.Context) error {
    c.Response().Header().Set("Access-Control-Allow-Origin", "*")
    c.Response().Header().Set("Access-Control-Expose-Headers", "Content-Type")

    c.Response().Header().Set("Content-Type", "text/event-stream")
    c.Response().Header().Set("Cache-Control", "no-cache")
    c.Response().Header().Set("Connection", "keep-alive")

    // Calculate deadline based on request context
    // ctx := c.Request().Context()
    // deadline, ok := ctx.Deadline()
    // if !ok {
    //     return errors.New("no deadline in request context")
    // }

    data := make(chan string, 20)
 
    // stringse := []string{"hi", "ay"}

    // err := json.NewDecoder(c.Request().Body).Decode(username)
    // if err != nil {
    //     return err
    // }

    username := c.QueryParam("username")


    go getData(data, username)
    printMessage(data, c)
    

    // for {
    //     select {
    //     // case <-ctx.Done():
    //     //     // Client disconnected or context cancelled
    //     //     return nil
    //     // jika ada channel yang masuk
    //     case msg, ok := <-data:
    //         // Check if data channel is closed
    //         if !ok {
    //             return nil // Stop sending data
    //         }
    //         fmt.Println("ini datanya")
            // fmt.Fprintf(c.Response(), "data: %s\n\n", msg)
            // c.Response().Flush()
            // time.Sleep(1 * time.Second)
    //     }

    //     // Check if deadline has passed
    //     // if time.Now().After(deadline) {
    //     //     // Deadline reached, end the stream
    //     //     return nil
    //     // }
    // }

    return nil
}

func getData(data chan<- string, username string) {
    // for _, val := range stringse {
    //     // mengirim val khe channel data
    //     data <- val
    // }

    if username == "alex" {
        for i := 0; i < 20; i++ {
            data <- fmt.Sprintf("data %d", i)
        }
    } else {
        for i := 0; i < 10; i++ {
            data <- fmt.Sprintf("data %d", i)
        }
    }

    close(data)
}


func printMessage(data <-chan string, c echo.Context) {
    for {
        select {
        case msg, ok := <- data:
            if !ok {
                break
            }
            fmt.Println(msg)
            fmt.Fprintf(c.Response(), "data: %s\n\n", msg)
            c.Response().Flush()
            time.Sleep(time.Millisecond * 500)
        }
    // for message := range data {

    // }
    }
}

func GetNotifFromEvent(c echo.Context) error {

    if c.Request().Method != "GET" {
        return echo.NewHTTPError(http.StatusBadRequest, "False http verb")
    }

    apiTarget := fmt.Sprintf("%s/%s", BASE_URL, "events")
    method := "GET"

    req, err := http.NewRequest(method, apiTarget, nil)
	if err != nil {
		// http.Error(c.Writer, "400", http.StatusBadRequest)
		return echo.NewHTTPError(http.StatusInternalServerError, "Server error")
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "bad request"})
		return
	}
	defer resp.Body.Close()
} 


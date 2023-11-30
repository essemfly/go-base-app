package main

import (
	"bytes"
	"encoding/json"
	"essemfly/go_base_app/internal/domain"
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

const NUM_MACHINES = 1
const NUM_CALLS = 10000

func main() {
	rand.Seed(time.Now().UnixNano())

	var wg sync.WaitGroup

	// 100개의 goroutine 생성 및 실행
	for i := 0; i < NUM_MACHINES; i++ {
		wg.Add(1)
		go sendLogData(generateRandomLogData(), &wg)
	}

	// 모든 goroutine이 완료될 때까지 대기
	wg.Wait()
}

func generateRandomLogData() domain.LogData {
	// 무작위 LogData 생성
	return domain.LogData{
		Topic:     fmt.Sprintf("logResults"),
		Message:   fmt.Sprintf("Random Message %d", rand.Intn(1000)),
		Score:     rand.Intn(100),
		CreatedAt: time.Now(),
	}
}

func sendLogData(logData domain.LogData, wg *sync.WaitGroup) {
	defer wg.Done() // 함수 실행이 끝나면 WaitGroup 카운터 감소

	i := 0
	for i < NUM_CALLS {
		// LogData 구조체를 JSON으로 변환
		jsonData, err := json.Marshal(logData)
		if err != nil {
			fmt.Println("Error marshalling JSON:", err)
			return
		}

		// HTTP 요청 생성
		url := "http://127.0.0.1:8080/logs"
		request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
		if err != nil {
			fmt.Println("Error creating request:", err)
			return
		}
		request.Header.Set("Content-Type", "application/json")

		// HTTP 클라이언트로 요청 보내기
		client := &http.Client{}
		response, err := client.Do(request)
		if err != nil {
			fmt.Println("Error sending request:", err)
			return
		}
		defer response.Body.Close()
		i++
		time.Sleep(1 * time.Second)
	}
}

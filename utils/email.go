package utils

import "fmt" 
import data"booking-app/data_classes"
import "time"
import "sync"

func SendEmail(requests chan data.EmailRequest, wg *sync.WaitGroup) {
	defer wg.Done()

	for request := range requests {
		fmt.Printf("Sending an email to %v\n", request.To)
		time.Sleep(2 * time.Second)
		fmt.Printf("Email sent to: %s\n", request.To)
	}
}
package main

import (
	"sync"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

const maxMessages = 5

func startScheduler(stopChan <-chan struct{}, messageChan chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			message := "Contrary to popular belief, Lorem Ipsum is not simply random text. It has roots in a piece of classical Latin literature from 45 BC, making it over 2000 years old. Richard McClintock, a Latin professor at Hampden-Sydney College in Virginia, looked up one of the more obscure Latin words, consectetur, from a Lorem Ipsum passage, and going through the cites of the word in classical literature, discovered the undoubtable source. Lorem Ipsum comes from sections 1.10.32 and 1.10.33 of de Finibus Bonorum et Malorum (The Extremes of Good and Evil) by Cicero, written in 45 BC. This book is a treatise on the theory of ethics, very popular during the Renaissance. The first line of Lorem Ipsum, Lorem ipsum dolor sit amet.., comes from a line in section 1.10.32."
			messageChan <- message
		case <-stopChan:
			messageChan <- "Scheduler stopped"
			return
		}
	}
}

func main() {
	a := app.New()
	w := a.NewWindow("Scheduler GUI")
	w.Resize(fyne.NewSize(600, 400))

	stopChan := make(chan struct{})
	messageChan := make(chan string, 100)
	var wg sync.WaitGroup
	var schedulerRunning bool

	messages := make([]string, 0, maxMessages)
	output := widget.NewLabel("")
	scrollContainer := container.NewScroll(output)
	scrollContainer.SetMinSize(fyne.NewSize(600, 300))

	updateTicker := time.NewTicker(1 * time.Second)
	defer updateTicker.Stop()

	go func() {
		for {
			select {
			case message := <-messageChan:
				if len(messages) >= maxMessages {
					messages = messages[1:] // Remove the oldest message
				}
				messages = append(messages, message)
			case <-updateTicker.C:
				output.SetText("")
				for _, msg := range messages {
					output.SetText(output.Text + msg + "\n")
				}
				output.Refresh()
			}
		}
	}()

	startButton := widget.NewButton("Start", func() {
		if !schedulerRunning {
			schedulerRunning = true
			messageChan <- "Starting scheduler..."
			wg.Add(1)
			go startScheduler(stopChan, messageChan, &wg)
		}
	})

	stopButton := widget.NewButton("Stop", func() {
		if schedulerRunning {
			stopChan <- struct{}{}
			wg.Wait()
			schedulerRunning = false
		}
	})

	w.SetCloseIntercept(func() {
		if !schedulerRunning {
			w.Close()
			close(messageChan) // Close the channel when the window is closed
		} else {
			messageChan <- "Stop the scheduler before closing."
		}
	})

	w.SetContent(container.NewVBox(
		startButton,
		stopButton,
		scrollContainer,
	))

	w.ShowAndRun()
}

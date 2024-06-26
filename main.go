package main

import (
	"sync"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func startScheduler(stopChan <-chan struct{}, appendMessage func(string), wg *sync.WaitGroup) {
	defer wg.Done()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			message := "Hello World"
			appendMessage(message)
		case <-stopChan:
			appendMessage("Scheduler stopped")
			return
		}
	}
}

func main() {
	a := app.New()
	w := a.NewWindow("Scheduler GUI")
	w.Resize(fyne.NewSize(600, 400))

	stopChan := make(chan struct{})
	var wg sync.WaitGroup
	var schedulerRunning bool

	output := widget.NewLabel("")
	scrollContainer := container.NewScroll(output)
	scrollContainer.SetMinSize(fyne.NewSize(600, 300))

	appendMessage := func(message string) {
		output.SetText(output.Text + message + "\n")
		scrollContainer.ScrollToBottom()
	}

	startButton := widget.NewButton("Start", func() {
		if !schedulerRunning {
			schedulerRunning = true
			appendMessage("Starting scheduler...")
			wg.Add(1)
			go startScheduler(stopChan, appendMessage, &wg)
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
		} else {
			appendMessage("Stop the scheduler before closing.")
		}
	})

	w.SetContent(container.NewVBox(
		startButton,
		stopButton,
		scrollContainer,
	))

	w.ShowAndRun()
}

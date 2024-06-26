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
			// message := "Hello World"
			message := " Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed quis nisl justo. Cras at convallis est. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Donec a lacus vitae libero vehicula auctor fringilla ac metus. Curabitur mauris nisi, feugiat et faucibus id, interdum eu risus. Pellentesque vehicula sodales accumsan. Proin id scelerisque diam. Vestibulum hendrerit erat eget sapien auctor, nec egestas nisl scelerisque. Morbi auctor malesuada imperdiet. Maecenas tempor viverra ligula dictum auctor. Suspendisse aliquet consequat finibus. Vestibulum gravida dolor id nulla tincidunt ultricies. Quisque congue sit amet nisi at faucibus. Nullam quis velit tortor. Vestibulum congue et velit sed finibus. Sed et velit semper, cursus ex in, ultricies orci. Praesent pulvinar semper libero, id sodales lectus luctus id. Pellentesque scelerisque velit ipsum, quis luctus leo faucibus eu. Integer vitae tortor ut justo sollicitudin ullamcorper. Phasellus eu orci a metus gravida lacinia. Suspendisse in neque vel tortor porta rhoncus id sit amet velit. Sed ut metus mauris. Mauris consectetur fringilla ultrices. Pellentesque malesuada pulvinar efficitur. Suspendisse ante urna, congue eget neque vitae, placerat dictum ligula. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aenean ullamcorper aliquam velit ac eleifend. Cras ut dui in nisl laoreet pretium non scelerisque mi. Vestibulum odio purus, ullamcorper sed turpis ut, venenatis ultricies nulla. Quisque eget posuere ligula. In eget cursus ex. Suspendisse elementum et purus ut tempus. Nullam eu turpis fringilla, finibus urna eleifend, iaculis lacus. Aliquam vehicula volutpat quam, non placerat neque suscipit non. Nulla pretium eros nec augue faucibus dictum. Maecenas nec placerat velit, ut sagittis libero. Nam sed vehicula tellus. Praesent at sapien ante. Donec id mattis turpis. Integer tincidunt sit amet mauris nec ultricies. Aenean fermentum lorem leo, non commodo nunc ullamcorper vel. Aliquam vel velit at enim convallis ultricies et pharetra turpis. Morbi fringilla diam erat. Praesent lacinia pretium tellus et pellentesque. Phasellus placerat ex id quam ultricies egestas. Ut hendrerit ipsum ut tellus blandit, a vulputate enim pharetra. In aliquet dapibus felis id laoreet. Nullam aliquet quis dolor non feugiat. Vivamus vitae tortor a purus eleifend pellentesque sed non urna. Sed mattis erat ut sapien consequat, molestie finibus velit ornare. Nam blandit nec sapien at imperdiet. Donec congue rhoncus ante in porta. Proin rhoncus metus turpis, eu molestie mi rutrum id. Nulla quis massa eu ipsum mattis commodo. Duis rhoncus elit nec dui ultricies, fermentum porta metus tempor. Vestibulum posuere, nisi nec dignissim vulputate, mi nulla imperdiet metus, vitae convallis leo leo nec urna. Vivamus ultrices, eros a placerat congue, felis nisi vulputate justo, eget lacinia quam diam vulputate elit. Nunc diam dolor, consectetur in accumsan id, pellentesque ac odio. In sollicitudin urna sit amet ligula molestie malesuada. Nullam at lectus luctus, venenatis sem id, sagittis est. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas cursus fermentum metus eu laoreet. In placerat urna at felis pharetra bibendum eu ut lectus. Pellentesque dapibus sit amet dolor ac malesuada. Praesent nec porta enim. Nullam sed sagittis lectus. Curabitur gravida felis vitae turpis suscipit pellentesque. Mauris vel ante et nibh auctor scelerisque. Fusce libero nunc, placerat in justo non, pretium hendrerit lorem. Mauris non ullamcorper eros. Cras ac tellus quis elit faucibus maximus. Proin tincidunt tristique sem in ultrices. Donec nec urna ut erat scelerisque finibus at ac nulla. Phasellus justo tortor, placerat eu ligula et, euismod euismod nisl. Aliquam rutrum mi egestas sem efficitur vestibulum. Suspendisse facilisis dui eget magna tristique convallis. Phasellus eget venenatis libero. Fusce et ultrices enim. Nulla commodo est et dui gravida, ut vulputate quam auctor. Duis ut libero vel ipsum imperdiet tempus fermentum eget lorem. Aliquam placerat, diam non sollicitudin interdum, ligula nisl venenatis mi, ut rhoncus velit turpis non dui. Vestibulum iaculis eros sit amet quam pretium, ac laoreet ligula tincidunt. Vestibulum molestie sagittis justo, eu scelerisque leo facilisis ac. Integer viverra non mauris sit amet sollicitudin. Nullam efficitur sapien non erat vehicula varius. Aliquam erat volutpat. Donec maximus justo sit amet pretium rutrum. Nulla eu sem varius, tristique felis sed, scelerisque purus. Aliquam in libero imperdiet eros efficitur ultricies. Morbi ut pretium nibh, nec varius ligula. Proin iaculis nisi in diam porttitor, ut bibendum turpis mattis. Nulla facilisis eu ligula et iaculis. Quisque a augue leo. Sed sed justo bibendum, vestibulum libero vitae, ullamcorper velit. Aenean et imperdiet tortor. Morbi accumsan nunc non dolor posuere porta. Nullam fermentum a turpis a viverra. Maecenas ornare id dui non dictum. Nulla nulla purus, egestas at hendrerit sed, mollis id tortor. Duis viverra metus nec diam consequat, vel fringilla nulla tincidunt. Praesent tincidunt molestie aliquam. Mauris in justo tincidunt, eleifend sem et, dignissim odio. Morbi ut ultrices mi. Cras at scelerisque dui. Proin bibendum vitae dui at laoreet. Proin sed augue sed augue eleifend maximus. Vestibulum convallis lectus sit amet quam cursus volutpat id at nulla. Nulla nec dolor condimentum, gravida nibh eu, luctus sapien.Donec vel sem sed mauris blandit imperdiet. Sed vehicula odio id ullamcorper tristique. Etiam viverra leo eget lorem posuere, vitae venenatis augue molestie. Quisque et lorem porta ipsum dignissim sagittis. Nulla ut sodales lorem. Suspendisse potenti. Quisque ultricies lacus eu suscipit cursus. Ut sagittis massa at bibendum vehicula. Integer sit amet sagittis velit. Sed interdum porttitor arcu id tincidunt. Praesent pharetra euismod imperdiet. Suspendisse quis lectus ligula. Aenean a felis nec nulla finibus finibus. Donec fringilla porttitor arcu, vitae faucibus arcu pulvinar nec. Fusce in posuere ex. Vestibulum eget mi id velit euismod maximus."
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

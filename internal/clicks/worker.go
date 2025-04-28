package clicks

import (
	"ad-server/pkg/metrics"
)

type ClickWorker struct {
	service *ClickService
	queue   chan ClickEvent
}

func NewClickWorker(service *ClickService) *ClickWorker {
	return &ClickWorker{
		service: service,
		queue:   make(chan ClickEvent, 10000),
	}
}

func (w *ClickWorker) Start() {
	for click := range w.queue {
		_ = w.service.SaveClick(click)
	}
}

func (w *ClickWorker) Enqueue(click ClickEvent) {
	select {
	case w.queue <- click:
		metrics.ClickCounter.Inc()
	default:
		// drop or log if queue full
	}
}

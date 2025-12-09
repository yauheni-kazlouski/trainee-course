package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
)

// CompositeHandler holds multiple slog handlers
type CompositeHandler struct{
	handlers []slog.Handler
}

// NewCompositeHandler creates handler with given handlers
func NewCompositeHandler(source ...slog.Handler) *CompositeHandler {
	res := make([]slog.Handler, len(source))
	copy(res, source)
	return &CompositeHandler{res}
}

// Enabled returns true if ALL handlers are enabled for this level
func (h *CompositeHandler) Enabled(ctx context.Context, l slog.Level) bool {
	f := true

	for _, el := range(h.handlers){
		if !el.Enabled(ctx, l) {
			f = false
		}
	}

	return f
}

// Handle passes record to all handlers
func (h *CompositeHandler) Handle(ctx context.Context, rec slog.Record) error {

	for _, el := range(h.handlers){
		err := el.Handle(ctx, rec)
		if err != nil {
			fmt.Printf("Error with %#v handler: %s", el, err.Error())
		}
	}

	return nil
}

// WithAttrs creates new handler with added attributes
func (h *CompositeHandler) WithAttrs(attrs []slog.Attr) slog.Handler{
	result := make([]slog.Handler, len(h.handlers))

	for i, el := range(h.handlers){
		result[i] = el.WithAttrs(attrs)
	}

	return &CompositeHandler{result}
}

// WithGroup creates new handler with group name
func (h *CompositeHandler) WithGroup(name string) slog.Handler{
	result := make([]slog.Handler, len(h.handlers))

	for i, el := range(h.handlers){
		result[i] = el.WithGroup(name)
	}

	return &CompositeHandler{result}
}

func main() {
	// Open (create) log file
	logFile, err := os.OpenFile("test.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer logFile.Close()

	// Create composite handler for JSON file and text output and logger with this handler
	compHandler := NewCompositeHandler(slog.NewJSONHandler(logFile, nil), slog.NewTextHandler(os.Stdout, nil))
	compLogger := slog.New(compHandler)

	// Test message to both outputs
	compLogger.Info("testing composite logger")

}
package storage

import (
	"context"
	"errors"
	"testing"

	"github.com/pedrocvaranda/portfolio-tracker/portfolio/models"
)

func TestMemoryPortfolio(t *testing.T) {
	store := NewMemory()
	p := models.NewPortfolio("p1", "Main", "USD")
	if err := store.SavePortfolio(context.Background(), p); err != nil {
		t.Fatal(err)
	}
	got, err := store.GetPortfolio(context.Background(), "p1")
	if err != nil {
		t.Fatal(err)
	}
	if got.ID != p.ID {
		t.Fatalf("got %q, want %q", got.ID, p.ID)
	}
}

func TestMemoryNotFound(t *testing.T) {
	_, err := NewMemory().GetPortfolio(context.Background(), "missing")
	if !errors.Is(err, ErrNotFound) {
		t.Fatalf("err = %v, want ErrNotFound", err)
	}
}

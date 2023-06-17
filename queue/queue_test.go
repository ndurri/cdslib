package queue

import (
	"testing"
)

func TestNew(t *testing.T) {
	q := NewQueue("queue")
    if q.Name != "queue" {
    	t.Errorf("q.Name = %s, want 'queue'", q.Name)
    }
}

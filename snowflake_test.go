package snowflake

import (
	"testing"
	"time"
)

func Test_New(t *testing.T) {
	S := New()

	if S.epoch != 0 {
		t.Logf("expected 0, got %d", S.epoch)
		t.FailNow()
	}

	if S.InternalWorkerID != 0 {
		t.Logf("expected 0, got %d", S.InternalWorkerID)
		t.FailNow()
	}

	if S.InternalProcessID != 0 {
		t.Logf("expected 0, got %d", S.InternalProcessID)
		t.FailNow()
	}

	if S.IncrimentID != 0 {
		t.Logf("expected 0, got %d", S.IncrimentID)
		t.FailNow()
	}

	t.Log(S)
}

func Test_NewWithOptions(t *testing.T) {
	now := time.Now()
	S := NewWithOptions(
		WithEpoch(12345),
		WithSpecificTime(now),
		WithSpecificWorkerID(12345),
		WithSpecificProcessID(12345),
		WithSpecificIncrimentID(12345))

	if S.epoch != 12345 {
		t.Logf("expected 12345, got %d", S.epoch)
		t.FailNow()
	}

	if S.Timestamp != now {
		t.Logf("expected %s, got %s", now, S.Timestamp)
		t.FailNow()
	}

	if S.InternalWorkerID != 12345 {
		t.Logf("expected 12345, got %d", S.InternalWorkerID)
		t.FailNow()
	}

	if S.InternalProcessID != 12345 {
		t.Logf("expected 12345, got %d", S.InternalProcessID)
		t.FailNow()
	}

	if S.IncrimentID != 12345 {
		t.Logf("expected 12345, got %d", S.IncrimentID)
		t.FailNow()
	}

	t.Log(S)
}

func Test_Unmarshal(t *testing.T) {
	S := New()

	if err := S.Unmarshal([]int{19}); err != nil {
		t.Log(err)
		t.FailNow()
	}
}

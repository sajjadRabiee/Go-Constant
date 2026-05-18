package constant

import (
	"testing"
)

// --- string-based type ---

type direction string

func (d direction) String() string { return string(d) }

const (
	north direction = "north"
	south direction = "south"
	east  direction = "east"
)

var dirSet = NewConstSet(north, north, south, east)

// --- struct-based type ---

type color struct {
	name string
	hex  string
}

func (c color) String() string { return c.name }

var (
	red  = color{name: "red", hex: "#FF0000"}
	blue = color{name: "blue", hex: "#0000FF"}

	colorSet = NewConstSet(red, red, blue)
)

// --- tests ---

func TestParse_Found(t *testing.T) {
	got, err := dirSet.Parse("south")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != south {
		t.Fatalf("got %q, want %q", got, south)
	}
}

func TestParse_NotFound_ReturnsDefault(t *testing.T) {
	got, err := dirSet.Parse("up")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if got != north {
		t.Fatalf("got %q, want default %q", got, north)
	}
}

func TestParse_ErrorMessage(t *testing.T) {
	_, err := dirSet.Parse("up")
	want := `"up" is not a valid constant`
	if err.Error() != want {
		t.Fatalf("got error %q, want %q", err.Error(), want)
	}
}

func TestParse_EmptyString(t *testing.T) {
	_, err := dirSet.Parse("")
	if err == nil {
		t.Fatal("expected error for empty string, got nil")
	}
}

func TestParse_Struct_Found(t *testing.T) {
	got, err := colorSet.Parse("blue")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.hex != "#0000FF" {
		t.Fatalf("got hex %q, want %q", got.hex, "#0000FF")
	}
}

func TestParse_Struct_NotFound_ReturnsDefault(t *testing.T) {
	got, err := colorSet.Parse("green")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if got != red {
		t.Fatalf("got %v, want default %v", got, red)
	}
}

func TestMembers(t *testing.T) {
	members := dirSet.Members()
	if len(members) != 3 {
		t.Fatalf("got %d members, want 3", len(members))
	}
	want := []direction{north, south, east}
	for i, m := range members {
		if m != want[i] {
			t.Errorf("members[%d] = %q, want %q", i, m, want[i])
		}
	}
}

func TestDefault(t *testing.T) {
	if dirSet.Default() != north {
		t.Fatalf("got %q, want %q", dirSet.Default(), north)
	}
}

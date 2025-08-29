package common

import (
	"errors"
	"testing"
)

type SrcSimple struct {
	ID   int
	Name string
}

type DstSimple struct {
	ID   int
	Name string
}

type SrcNested struct {
	ID   int
	Info SrcInfo
}

type SrcInfo struct {
	Email string
	Age   int
}

type DstNested struct {
	ID   int
	Info DstInfo
}

type DstInfo struct {
	Email string
	Age   int
}

type SrcPartial struct {
	ID int
}

type DstPartial struct {
	ID   int
	Name string
}

func TestConvert_SimpleStruct(t *testing.T) {
	src := SrcSimple{ID: 1, Name: "Alice"}
	got, err := Convert[DstSimple](src)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.ID != src.ID || got.Name != src.Name {
		t.Errorf("expected %+v, got %+v", src, got)
	}
}

func TestConvert_StructPointer(t *testing.T) {
	src := &SrcSimple{ID: 2, Name: "Bob"}
	got, err := Convert[DstSimple](src)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.ID != src.ID || got.Name != src.Name {
		t.Errorf("expected %+v, got %+v", src, got)
	}
}

func TestConvert_NestedStruct(t *testing.T) {
	src := SrcNested{ID: 3, Info: SrcInfo{Email: "a@b.com", Age: 30}}
	got, err := Convert[DstNested](src)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.ID != src.ID || got.Info.Email != src.Info.Email || got.Info.Age != src.Info.Age {
		t.Errorf("expected %+v, got %+v", src, got)
	}
}

func TestConvert_PartialFields(t *testing.T) {
	src := SrcPartial{ID: 4}
	got, err := Convert[DstPartial](src)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.ID != src.ID {
		t.Errorf("expected ID %d, got %d", src.ID, got.ID)
	}
	if got.Name != "" {
		t.Errorf("expected Name to be empty, got %s", got.Name)
	}
}

func TestConvert_InvalidType(t *testing.T) {
	_, err := Convert[DstSimple](123)
	if err == nil {
		t.Error("expected error for non-struct input, got nil")
	}
	if !errors.Is(err, ErrNotStruct) {
		t.Errorf("unexpected error: %v", err)
	}
}

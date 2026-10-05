package registry

import "testing"

func TestAssignThenOwner(t *testing.T) {
	book := New()
	code := book.Assign("alice")
	if book.Owner(code) != "alice" {
		t.Fatalf("码的归属应该是 alice，得到 %q", book.Owner(code))
	}
}

func TestUnknownCodeHasNoOwner(t *testing.T) {
	if New().Owner("ZZZZZZ") != "" {
		t.Fatal("没分配过的码不该有归属")
	}
}

func TestPlayersStartEmpty(t *testing.T) {
	if len(New().Players()) != 0 {
		t.Fatal("空台账不该有玩家")
	}
}

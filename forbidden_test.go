package forbidden

import "testing"

func TestNoDuplicates(t *testing.T) {
	if len(names) != len(nameMap) {
		t.Fatal("There are duplicates in the list")
	}
}

func TestForbidden(t *testing.T) {
	names := []string{
		"ADMIN",
		"Admin",
		"admin",
		".json",
		".JSON",
		"  admin   ",
		"admin   ",
		" admin",
	}

	for _, name := range names {
		if !Name(name) {
			t.Fatal("Expected forbidden but got fine for " + name)
		}
	}
}

func TestNotForbidden(t *testing.T) {
	names := []string{
		"ADMI",
		".admin",
		"..json",
		".json.",
		".json.XML",
		"tar.gz",
		"JSONN",
	}

	for _, name := range names {
		if Name(name) {
			t.Fatal("Expected fine but got forbidden for " + name)
		}
	}
}

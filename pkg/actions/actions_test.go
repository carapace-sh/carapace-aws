package actions

import "testing"

func TestMacrosExposed(t *testing.T) {
	for _, name := range []string{"Profiles", "Regions"} {
		macro, ok := Macros[name]
		if !ok {
			t.Fatalf("expected macro %q to be exposed", name)
		}
		if macro.Name != name {
			t.Fatalf("expected macro %q to use matching name, got %q", name, macro.Name)
		}
		if macro.Function == "" {
			t.Fatalf("expected macro %q to expose its source function", name)
		}
	}
}

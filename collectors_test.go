package main

import "testing"

func TestIsValidCollector(t *testing.T) {
	for _, collector := range validCollectors {
		if !isValidCollector(collector) {
			t.Errorf("Collector %s detected as non-valid", collector)
		}
	}

	randomCollectors := []string{"toto", "tata", "test", "hello"}
	for _, collector := range randomCollectors {
		if isValidCollector(collector) {
			t.Errorf("Collector %s is detected as valid", collector)
		}
	}
}

func TestCollectorSliceContains(t *testing.T) {
	cs := collectorSlice{"toto", "tata", "test", "hello"}

	for _, item := range cs {
		if !cs.Contains(item) {
			t.Errorf("%s is not detected in the slice", item)
		}
	}

	for _, item := range []string{"another", "set"} {
		if cs.Contains(item) {
			t.Errorf("%s is detected in the slice", item)
		}
	}
}

func TestCollectorSliceString(t *testing.T) {
	cs := collectorSlice{"toto", "tata", "test", "hello"}
	str := cs.String()
	if str != "Collectors: toto, tata, test, hello" {
		t.Errorf("Invalid String transformation of the collectorSlice")
	}
}

func TestCollectorSliceSet(t *testing.T) {
	var cs collectorSlice

	err := cs.Set("toto")
	if err == nil {
		t.Error("`toto` is not a valid collector")
	}

	err = cs.Set(validCollectors[0])
	if err != nil {
		t.Errorf("%s is a valid collector", validCollectors[0])
	}
}

func TestCollectorSliceSetDefaultCollector(t *testing.T) {
	var cs collectorSlice

	cs.SetDefaultCollector()

	if len(cs) > 1 {
		t.Errorf("More than one default collector set: %s", cs)
	}

	if cs[0] != validCollectors[0] {
		t.Errorf("Wrong default collector. Want: %s, Got: %s", validCollectors[0], cs[0])
	}
}

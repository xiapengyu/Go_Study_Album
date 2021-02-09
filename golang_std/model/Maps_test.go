package main

import "testing"

func search(m map[string]string, word string) string {
	return m[word]
}

func Test_search(t *testing.T){
	m := map[string]string{"name": "Tom", "address": "China"}
	got := search(m, "name")

	want := "Just do it"

	assertStrings(t, got, want)
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

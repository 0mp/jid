package jig

import (
	"bytes"
	"github.com/bitly/go-simplejson"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestNewSuggestion(t *testing.T) {
	var assert = assert.New(t)
	assert.Equal(NewSuggestion(), &Suggestion{})
}

func TestSuggestionGet(t *testing.T) {
	var assert = assert.New(t)
	j := createJson(`{"name":"simeji-github", "naming":"simeji", "nickname":"simejisimeji"}`)
	s := NewSuggestion()
	assert.Equal([]string{"m", "nam"}, s.Get(j, "na"))

	j = createJson(`{"abcde":"simeji-github", "abcdef":"simeji", "ab":"simejisimeji"}`)
	assert.Equal([]string{".", "."}, s.Get(j, ""))
	assert.Equal([]string{"b", "ab"}, s.Get(j, "a"))
	assert.Equal([]string{"de", "abcde"}, s.Get(j, "abc"))
	assert.Equal([]string{"", "abcde"}, s.Get(j, "abcde"))

	j = createJson(`["zero"]`)
	assert.Equal([]string{"[0]", "[0]"}, s.Get(j, ""))
	assert.Equal([]string{"0]", "[0]"}, s.Get(j, "["))
	assert.Equal([]string{"]", "[0]"}, s.Get(j, "[0"))

	j = createJson(`["zero", "one"]`)
	assert.Equal([]string{"[", "["}, s.Get(j, ""))

	assert.Equal([]string{"", "["}, s.Get(j, "["))
}

func TestSuggestionGetCurrentType(t *testing.T) {
	var assert = assert.New(t)
	s := NewSuggestion()

	j := createJson(`[1,2,3]`)
	assert.Equal(ARRAY, s.GetCurrentType(j))
	j = createJson(`{"name":[1,2,3], "naming":{"account":"simeji"}, "test":"simeji", "testing":"ok"}`)
	assert.Equal(MAP, s.GetCurrentType(j))
	j = createJson(`"name"`)
	assert.Equal(STRING, s.GetCurrentType(j))
}

func TestSuggestionGetCandidateKeys(t *testing.T) {
	var assert = assert.New(t)
	j := createJson(`{"naming":"simeji", "nickname":"simejisimeji", "city":"tokyo", "name":"simeji-github" }`)
	s := NewSuggestion()

	assert.Equal([]string{"city", "name", "naming", "nickname"}, s.GetCandidateKeys(j, ""))
	assert.Equal([]string{"name", "naming", "nickname"}, s.GetCandidateKeys(j, "n"))
	assert.Equal([]string{"name", "naming"}, s.GetCandidateKeys(j, "na"))

	j = createJson(`{"abcde":"simeji-github", "abcdef":"simeji", "ab":"simejisimeji"}`)
	assert.Equal([]string{"abcde", "abcdef"}, s.GetCandidateKeys(j, "abcde"))

	j = createJson(`[1,2,"aa"]`)
	s = NewSuggestion()
	assert.Equal([]string{}, s.GetCandidateKeys(j, "["))
}

func createJson(s string) *simplejson.Json {
	r := bytes.NewBufferString(s)
	buf, _ := ioutil.ReadAll(r)
	j, _ := simplejson.NewJson(buf)
	return j
}

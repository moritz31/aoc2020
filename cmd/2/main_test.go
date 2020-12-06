package main

import "testing"

const input = `1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc`

func TestParsing(t *testing.T) {
	count := parsePasswordList(input)

	if count != 1 {
		t.Errorf("Expected 1 got %d", count)
	}
}

func TestRuleGeneration(t *testing.T) {
	sampleInput := "1-3 a: abcde"
	createdRule := generateRule(sampleInput)

	if createdRule.min != 1 && createdRule.max != 3 && createdRule.character != "a" {
		t.Errorf("Rule does not match")
	}
}

func TestRuleMatching(t *testing.T) {
	sampleRule := rule{min: 1, max: 3, character: "a"}
	samplePassword := " abcde"

	valid := validatePasswordWithRule(sampleRule, samplePassword)

	if valid != true {
		t.Errorf("Invalid password")
	}
}

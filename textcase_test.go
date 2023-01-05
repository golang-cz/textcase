package textcase

import (
	"strings"
	"testing"
)

func TestTextCases(t *testing.T) {
	tt := []struct {
		in    string
		camel string
		snake string
	}{
		{in: "", camel: "", snake: ""},
		{in: "_", camel: "", snake: ""},
		{in: "a", camel: "a", snake: "a"},
		{in: "a___", camel: "a", snake: "a"},
		{in: "___a", camel: "a", snake: "a"},
		{in: "a_b", camel: "aB", snake: "a_b"},
		{in: "a___b", camel: "aB", snake: "a_b"},
		{in: "ax___by", camel: "axBy", snake: "ax_by"},
		{in: "someText", camel: "someText", snake: "some_text"},
		{in: "someTEXT", camel: "someText", snake: "some_text"},
		{in: "*postgresql.JSONBConverter", camel: "postgresqlJsonbConverter", snake: "postgresql_jsonb_converter"},
		{in: "NeXT", camel: "neXt", snake: "ne_xt"},
		{in: "Add updated_at to users table", camel: "addUpdatedAtToUsersTable", snake: "add_updated_at_to_users_table"},
		{in: "Hey, this TEXT will have to obey some rules!!", camel: "heyThisTextWillHaveToObeySomeRules", snake: "hey_this_text_will_have_to_obey_some_rules"},
		{in: "Háčky, čárky. Příliš žluťoučký kůň úpěl ďábelské ódy.", camel: "háčkyČárkyPřílišŽluťoučkýKůňÚpělĎábelskéÓdy", snake: "háčky_čárky_příliš_žluťoučký_kůň_úpěl_ďábelské_ódy"},
		{in: "here comes O'Brian", camel: "hereComesOBrian", snake: "here_comes_o_brian"},
		{in: "thisIsCamelCase", camel: "thisIsCamelCase", snake: "this_is_camel_case"},
		{in: "this_is_snake_case", camel: "thisIsSnakeCase", snake: "this_is_snake_case"},
		{in: "__snake_case__", camel: "snakeCase", snake: "snake_case"},
		{in: "fromCamelCaseToCamelCase", camel: "fromCamelCaseToCamelCase", snake: "from_camel_case_to_camel_case"},
		{in: "$()&^%(_--crazy__--input$)", camel: "crazyInput", snake: "crazy_input"},
		{in: "_$$_This is some text, OK?!", camel: "thisIsSomeTextOk", snake: "this_is_some_text_ok"},
		{in: "$(((*&^%$#@!)))#$%^&*", camel: "", snake: ""},
	}

	for _, test := range tt {
		// camelCase
		if got := CamelCase(test.in); got != test.camel {
			t.Errorf("unexpected camelCase for %q: got %q, want %q", test.in, got, test.camel)
		}

		// PascalCase
		testPascal := strings.Title(test.camel)
		if got := PascalCase(test.in); got != testPascal {
			t.Errorf("unexpected PascalCase for %q: got %q, want %q", test.in, got, testPascal)
		}

		// snake_case
		if got := SnakeCase(test.in); got != test.snake {
			t.Errorf("unexpected snake_case for %q: got %q, want %q", test.in, got, test.snake)
		}

		// kebab-case
		testKebab := strings.ReplaceAll(test.snake, "_", "-")
		if got := KebabCase(test.in); got != testKebab {
			t.Errorf("unexpected kebab-case for %q: got %q, want %q", test.in, got, testKebab)
		}
	}
}

func TestMarkLetterCaseChanges(t *testing.T) {
	tt := []struct {
		in  string
		out string
	}{
		{in: "detectUpperLowerChanges", out: "detect_Upper_Lower_Changes"},
		{in: "detectUPPERchange", out: "detect_UPPER_change"},
		{in: "detect_UPPER_change", out: "detect_UPPER_change"},
		{in: "Some camelCase and PascalCase text, OK?", out: "Some camel_Case and Pascal_Case text, OK?"},
	}

	for _, test := range tt {
		if got := markLetterCaseChanges(test.in); got != test.out {
			t.Errorf("unexpected markLowerUpperChanges for %q: got %q, want %q", test.in, got, test.out)
		}
	}
}

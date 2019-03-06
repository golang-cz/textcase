package camelsnake

import (
	"testing"
)

func TestCamelSnake(t *testing.T) {
	t.Parallel()

	tt := []struct {
		in    string
		camel string
		snake string
	}{
		{in: "Add updated_at to users table", camel: "AddUpdatedAtToUsersTable", snake: "add_updated_at_to_users_table"},
		{in: "$()&^%(_--crazy__--input$)", camel: "CrazyInput", snake: "crazy_input"},
		{in: "Hey, this TEXT will have to obey some rules!!", camel: "HeyThisTextWillHaveToObeySomeRules", snake: "hey_this_text_will_have_to_obey_some_rules"},
		{in: "_$$_This is some text, OK?!", camel: "ThisIsSomeTextOk", snake: "this_is_some_text_ok"},
		{in: "_", camel: "", snake: ""},
		{in: "$(((*&^%$#@!)))#$%^&*", camel: "", snake: ""},
		{in: "", camel: "", snake: ""},
		{in: "a", camel: "A", snake: "a"},
		{in: "a___", camel: "A", snake: "a"},
		{in: "a___b", camel: "AB", snake: "a_b"},
		{in: "ax___by", camel: "AxBy", snake: "ax_by"},
	}

	for _, test := range tt {
		if got := CamelCase(test.in); got != test.camel {
			t.Errorf("unexpected CamelCase for input(%q), got %q, want %q", test.in, got, test.camel)
		}
		if got := SnakeCase(test.in); got != test.snake {
			t.Errorf("unexpected snake_case for input(%q), got %q, want %q", test.in, got, test.snake)
		}
	}
}

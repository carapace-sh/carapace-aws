package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var waf_listSqlInjectionMatchSetsCmd = &cobra.Command{
	Use:   "list-sql-injection-match-sets",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waf_listSqlInjectionMatchSetsCmd).Standalone()

	waf_listSqlInjectionMatchSetsCmd.Flags().String("limit", "", "Specifies the number of [SqlInjectionMatchSet]() objects that you want AWS WAF to return for this request.")
	waf_listSqlInjectionMatchSetsCmd.Flags().String("next-marker", "", "If you specify a value for `Limit` and you have more [SqlInjectionMatchSet]() objects than the value of `Limit`, AWS WAF returns a `NextMarker` value in the response that allows you to list another group of `SqlInjectionMatchSets`.")
	wafCmd.AddCommand(waf_listSqlInjectionMatchSetsCmd)
}

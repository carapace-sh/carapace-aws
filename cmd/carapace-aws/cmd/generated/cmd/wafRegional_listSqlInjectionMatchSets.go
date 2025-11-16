package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafRegional_listSqlInjectionMatchSetsCmd = &cobra.Command{
	Use:   "list-sql-injection-match-sets",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafRegional_listSqlInjectionMatchSetsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wafRegional_listSqlInjectionMatchSetsCmd).Standalone()

		wafRegional_listSqlInjectionMatchSetsCmd.Flags().String("limit", "", "Specifies the number of [SqlInjectionMatchSet]() objects that you want AWS WAF to return for this request.")
		wafRegional_listSqlInjectionMatchSetsCmd.Flags().String("next-marker", "", "If you specify a value for `Limit` and you have more [SqlInjectionMatchSet]() objects than the value of `Limit`, AWS WAF returns a `NextMarker` value in the response that allows you to list another group of `SqlInjectionMatchSets`.")
	})
	wafRegionalCmd.AddCommand(wafRegional_listSqlInjectionMatchSetsCmd)
}

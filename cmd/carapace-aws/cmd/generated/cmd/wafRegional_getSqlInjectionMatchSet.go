package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafRegional_getSqlInjectionMatchSetCmd = &cobra.Command{
	Use:   "get-sql-injection-match-set",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafRegional_getSqlInjectionMatchSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wafRegional_getSqlInjectionMatchSetCmd).Standalone()

		wafRegional_getSqlInjectionMatchSetCmd.Flags().String("sql-injection-match-set-id", "", "The `SqlInjectionMatchSetId` of the [SqlInjectionMatchSet]() that you want to get.")
		wafRegional_getSqlInjectionMatchSetCmd.MarkFlagRequired("sql-injection-match-set-id")
	})
	wafRegionalCmd.AddCommand(wafRegional_getSqlInjectionMatchSetCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafRegional_deleteSqlInjectionMatchSetCmd = &cobra.Command{
	Use:   "delete-sql-injection-match-set",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafRegional_deleteSqlInjectionMatchSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wafRegional_deleteSqlInjectionMatchSetCmd).Standalone()

		wafRegional_deleteSqlInjectionMatchSetCmd.Flags().String("change-token", "", "The value returned by the most recent call to [GetChangeToken]().")
		wafRegional_deleteSqlInjectionMatchSetCmd.Flags().String("sql-injection-match-set-id", "", "The `SqlInjectionMatchSetId` of the [SqlInjectionMatchSet]() that you want to delete.")
		wafRegional_deleteSqlInjectionMatchSetCmd.MarkFlagRequired("change-token")
		wafRegional_deleteSqlInjectionMatchSetCmd.MarkFlagRequired("sql-injection-match-set-id")
	})
	wafRegionalCmd.AddCommand(wafRegional_deleteSqlInjectionMatchSetCmd)
}

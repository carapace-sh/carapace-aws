package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var waf_deleteSqlInjectionMatchSetCmd = &cobra.Command{
	Use:   "delete-sql-injection-match-set",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waf_deleteSqlInjectionMatchSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(waf_deleteSqlInjectionMatchSetCmd).Standalone()

		waf_deleteSqlInjectionMatchSetCmd.Flags().String("change-token", "", "The value returned by the most recent call to [GetChangeToken]().")
		waf_deleteSqlInjectionMatchSetCmd.Flags().String("sql-injection-match-set-id", "", "The `SqlInjectionMatchSetId` of the [SqlInjectionMatchSet]() that you want to delete.")
		waf_deleteSqlInjectionMatchSetCmd.MarkFlagRequired("change-token")
		waf_deleteSqlInjectionMatchSetCmd.MarkFlagRequired("sql-injection-match-set-id")
	})
	wafCmd.AddCommand(waf_deleteSqlInjectionMatchSetCmd)
}

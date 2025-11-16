package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafRegional_updateSqlInjectionMatchSetCmd = &cobra.Command{
	Use:   "update-sql-injection-match-set",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafRegional_updateSqlInjectionMatchSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wafRegional_updateSqlInjectionMatchSetCmd).Standalone()

		wafRegional_updateSqlInjectionMatchSetCmd.Flags().String("change-token", "", "The value returned by the most recent call to [GetChangeToken]().")
		wafRegional_updateSqlInjectionMatchSetCmd.Flags().String("sql-injection-match-set-id", "", "The `SqlInjectionMatchSetId` of the `SqlInjectionMatchSet` that you want to update.")
		wafRegional_updateSqlInjectionMatchSetCmd.Flags().String("updates", "", "An array of `SqlInjectionMatchSetUpdate` objects that you want to insert into or delete from a [SqlInjectionMatchSet]().")
		wafRegional_updateSqlInjectionMatchSetCmd.MarkFlagRequired("change-token")
		wafRegional_updateSqlInjectionMatchSetCmd.MarkFlagRequired("sql-injection-match-set-id")
		wafRegional_updateSqlInjectionMatchSetCmd.MarkFlagRequired("updates")
	})
	wafRegionalCmd.AddCommand(wafRegional_updateSqlInjectionMatchSetCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var waf_updateSqlInjectionMatchSetCmd = &cobra.Command{
	Use:   "update-sql-injection-match-set",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waf_updateSqlInjectionMatchSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(waf_updateSqlInjectionMatchSetCmd).Standalone()

		waf_updateSqlInjectionMatchSetCmd.Flags().String("change-token", "", "The value returned by the most recent call to [GetChangeToken]().")
		waf_updateSqlInjectionMatchSetCmd.Flags().String("sql-injection-match-set-id", "", "The `SqlInjectionMatchSetId` of the `SqlInjectionMatchSet` that you want to update.")
		waf_updateSqlInjectionMatchSetCmd.Flags().String("updates", "", "An array of `SqlInjectionMatchSetUpdate` objects that you want to insert into or delete from a [SqlInjectionMatchSet]().")
		waf_updateSqlInjectionMatchSetCmd.MarkFlagRequired("change-token")
		waf_updateSqlInjectionMatchSetCmd.MarkFlagRequired("sql-injection-match-set-id")
		waf_updateSqlInjectionMatchSetCmd.MarkFlagRequired("updates")
	})
	wafCmd.AddCommand(waf_updateSqlInjectionMatchSetCmd)
}

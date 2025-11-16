package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafRegional_createSqlInjectionMatchSetCmd = &cobra.Command{
	Use:   "create-sql-injection-match-set",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafRegional_createSqlInjectionMatchSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wafRegional_createSqlInjectionMatchSetCmd).Standalone()

		wafRegional_createSqlInjectionMatchSetCmd.Flags().String("change-token", "", "The value returned by the most recent call to [GetChangeToken]().")
		wafRegional_createSqlInjectionMatchSetCmd.Flags().String("name", "", "A friendly name or description for the [SqlInjectionMatchSet]() that you're creating.")
		wafRegional_createSqlInjectionMatchSetCmd.MarkFlagRequired("change-token")
		wafRegional_createSqlInjectionMatchSetCmd.MarkFlagRequired("name")
	})
	wafRegionalCmd.AddCommand(wafRegional_createSqlInjectionMatchSetCmd)
}

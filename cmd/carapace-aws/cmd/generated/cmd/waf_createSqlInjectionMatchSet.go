package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var waf_createSqlInjectionMatchSetCmd = &cobra.Command{
	Use:   "create-sql-injection-match-set",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waf_createSqlInjectionMatchSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(waf_createSqlInjectionMatchSetCmd).Standalone()

		waf_createSqlInjectionMatchSetCmd.Flags().String("change-token", "", "The value returned by the most recent call to [GetChangeToken]().")
		waf_createSqlInjectionMatchSetCmd.Flags().String("name", "", "A friendly name or description for the [SqlInjectionMatchSet]() that you're creating.")
		waf_createSqlInjectionMatchSetCmd.MarkFlagRequired("change-token")
		waf_createSqlInjectionMatchSetCmd.MarkFlagRequired("name")
	})
	wafCmd.AddCommand(waf_createSqlInjectionMatchSetCmd)
}

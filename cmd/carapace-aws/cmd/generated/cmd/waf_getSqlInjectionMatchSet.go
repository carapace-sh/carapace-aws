package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var waf_getSqlInjectionMatchSetCmd = &cobra.Command{
	Use:   "get-sql-injection-match-set",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waf_getSqlInjectionMatchSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(waf_getSqlInjectionMatchSetCmd).Standalone()

		waf_getSqlInjectionMatchSetCmd.Flags().String("sql-injection-match-set-id", "", "The `SqlInjectionMatchSetId` of the [SqlInjectionMatchSet]() that you want to get.")
		waf_getSqlInjectionMatchSetCmd.MarkFlagRequired("sql-injection-match-set-id")
	})
	wafCmd.AddCommand(waf_getSqlInjectionMatchSetCmd)
}

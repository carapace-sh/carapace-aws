package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafRegional_getChangeTokenStatusCmd = &cobra.Command{
	Use:   "get-change-token-status",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafRegional_getChangeTokenStatusCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wafRegional_getChangeTokenStatusCmd).Standalone()

		wafRegional_getChangeTokenStatusCmd.Flags().String("change-token", "", "The change token for which you want to get the status.")
		wafRegional_getChangeTokenStatusCmd.MarkFlagRequired("change-token")
	})
	wafRegionalCmd.AddCommand(wafRegional_getChangeTokenStatusCmd)
}

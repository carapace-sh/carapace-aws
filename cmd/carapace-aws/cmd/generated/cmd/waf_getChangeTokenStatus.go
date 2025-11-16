package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var waf_getChangeTokenStatusCmd = &cobra.Command{
	Use:   "get-change-token-status",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waf_getChangeTokenStatusCmd).Standalone()

	waf_getChangeTokenStatusCmd.Flags().String("change-token", "", "The change token for which you want to get the status.")
	waf_getChangeTokenStatusCmd.MarkFlagRequired("change-token")
	wafCmd.AddCommand(waf_getChangeTokenStatusCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var waf_deleteIpsetCmd = &cobra.Command{
	Use:   "delete-ipset",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waf_deleteIpsetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(waf_deleteIpsetCmd).Standalone()

		waf_deleteIpsetCmd.Flags().String("change-token", "", "The value returned by the most recent call to [GetChangeToken]().")
		waf_deleteIpsetCmd.Flags().String("ipset-id", "", "The `IPSetId` of the [IPSet]() that you want to delete.")
		waf_deleteIpsetCmd.MarkFlagRequired("change-token")
		waf_deleteIpsetCmd.MarkFlagRequired("ipset-id")
	})
	wafCmd.AddCommand(waf_deleteIpsetCmd)
}

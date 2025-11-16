package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafRegional_deleteIpsetCmd = &cobra.Command{
	Use:   "delete-ipset",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafRegional_deleteIpsetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wafRegional_deleteIpsetCmd).Standalone()

		wafRegional_deleteIpsetCmd.Flags().String("change-token", "", "The value returned by the most recent call to [GetChangeToken]().")
		wafRegional_deleteIpsetCmd.Flags().String("ipset-id", "", "The `IPSetId` of the [IPSet]() that you want to delete.")
		wafRegional_deleteIpsetCmd.MarkFlagRequired("change-token")
		wafRegional_deleteIpsetCmd.MarkFlagRequired("ipset-id")
	})
	wafRegionalCmd.AddCommand(wafRegional_deleteIpsetCmd)
}

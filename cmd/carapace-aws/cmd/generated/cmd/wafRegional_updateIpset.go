package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafRegional_updateIpsetCmd = &cobra.Command{
	Use:   "update-ipset",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafRegional_updateIpsetCmd).Standalone()

	wafRegional_updateIpsetCmd.Flags().String("change-token", "", "The value returned by the most recent call to [GetChangeToken]().")
	wafRegional_updateIpsetCmd.Flags().String("ipset-id", "", "The `IPSetId` of the [IPSet]() that you want to update.")
	wafRegional_updateIpsetCmd.Flags().String("updates", "", "An array of `IPSetUpdate` objects that you want to insert into or delete from an [IPSet]().")
	wafRegional_updateIpsetCmd.MarkFlagRequired("change-token")
	wafRegional_updateIpsetCmd.MarkFlagRequired("ipset-id")
	wafRegional_updateIpsetCmd.MarkFlagRequired("updates")
	wafRegionalCmd.AddCommand(wafRegional_updateIpsetCmd)
}

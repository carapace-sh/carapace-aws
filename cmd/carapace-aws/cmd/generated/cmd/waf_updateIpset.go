package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var waf_updateIpsetCmd = &cobra.Command{
	Use:   "update-ipset",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waf_updateIpsetCmd).Standalone()

	waf_updateIpsetCmd.Flags().String("change-token", "", "The value returned by the most recent call to [GetChangeToken]().")
	waf_updateIpsetCmd.Flags().String("ipset-id", "", "The `IPSetId` of the [IPSet]() that you want to update.")
	waf_updateIpsetCmd.Flags().String("updates", "", "An array of `IPSetUpdate` objects that you want to insert into or delete from an [IPSet]().")
	waf_updateIpsetCmd.MarkFlagRequired("change-token")
	waf_updateIpsetCmd.MarkFlagRequired("ipset-id")
	waf_updateIpsetCmd.MarkFlagRequired("updates")
	wafCmd.AddCommand(waf_updateIpsetCmd)
}

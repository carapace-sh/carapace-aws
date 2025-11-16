package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafRegional_getIpsetCmd = &cobra.Command{
	Use:   "get-ipset",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafRegional_getIpsetCmd).Standalone()

	wafRegional_getIpsetCmd.Flags().String("ipset-id", "", "The `IPSetId` of the [IPSet]() that you want to get.")
	wafRegional_getIpsetCmd.MarkFlagRequired("ipset-id")
	wafRegionalCmd.AddCommand(wafRegional_getIpsetCmd)
}

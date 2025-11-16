package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafRegional_createIpsetCmd = &cobra.Command{
	Use:   "create-ipset",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafRegional_createIpsetCmd).Standalone()

	wafRegional_createIpsetCmd.Flags().String("change-token", "", "The value returned by the most recent call to [GetChangeToken]().")
	wafRegional_createIpsetCmd.Flags().String("name", "", "A friendly name or description of the [IPSet]().")
	wafRegional_createIpsetCmd.MarkFlagRequired("change-token")
	wafRegional_createIpsetCmd.MarkFlagRequired("name")
	wafRegionalCmd.AddCommand(wafRegional_createIpsetCmd)
}

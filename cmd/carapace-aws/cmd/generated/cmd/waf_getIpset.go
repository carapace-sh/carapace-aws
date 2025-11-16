package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var waf_getIpsetCmd = &cobra.Command{
	Use:   "get-ipset",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waf_getIpsetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(waf_getIpsetCmd).Standalone()

		waf_getIpsetCmd.Flags().String("ipset-id", "", "The `IPSetId` of the [IPSet]() that you want to get.")
		waf_getIpsetCmd.MarkFlagRequired("ipset-id")
	})
	wafCmd.AddCommand(waf_getIpsetCmd)
}

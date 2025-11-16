package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var waf_createIpsetCmd = &cobra.Command{
	Use:   "create-ipset",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waf_createIpsetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(waf_createIpsetCmd).Standalone()

		waf_createIpsetCmd.Flags().String("change-token", "", "The value returned by the most recent call to [GetChangeToken]().")
		waf_createIpsetCmd.Flags().String("name", "", "A friendly name or description of the [IPSet]().")
		waf_createIpsetCmd.MarkFlagRequired("change-token")
		waf_createIpsetCmd.MarkFlagRequired("name")
	})
	wafCmd.AddCommand(waf_createIpsetCmd)
}

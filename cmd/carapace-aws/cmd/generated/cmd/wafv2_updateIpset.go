package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafv2_updateIpsetCmd = &cobra.Command{
	Use:   "update-ipset",
	Short: "Updates the specified [IPSet]().",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafv2_updateIpsetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wafv2_updateIpsetCmd).Standalone()

		wafv2_updateIpsetCmd.Flags().String("addresses", "", "Contains an array of strings that specifies zero or more IP addresses or blocks of IP addresses that you want WAF to inspect for in incoming requests.")
		wafv2_updateIpsetCmd.Flags().String("description", "", "A description of the IP set that helps with identification.")
		wafv2_updateIpsetCmd.Flags().String("id", "", "A unique identifier for the set.")
		wafv2_updateIpsetCmd.Flags().String("lock-token", "", "A token used for optimistic locking.")
		wafv2_updateIpsetCmd.Flags().String("name", "", "The name of the IP set.")
		wafv2_updateIpsetCmd.Flags().String("scope", "", "Specifies whether this is for a global resource type, such as a Amazon CloudFront distribution.")
		wafv2_updateIpsetCmd.MarkFlagRequired("addresses")
		wafv2_updateIpsetCmd.MarkFlagRequired("id")
		wafv2_updateIpsetCmd.MarkFlagRequired("lock-token")
		wafv2_updateIpsetCmd.MarkFlagRequired("name")
		wafv2_updateIpsetCmd.MarkFlagRequired("scope")
	})
	wafv2Cmd.AddCommand(wafv2_updateIpsetCmd)
}

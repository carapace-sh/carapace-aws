package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafv2_createIpsetCmd = &cobra.Command{
	Use:   "create-ipset",
	Short: "Creates an [IPSet](), which you use to identify web requests that originate from specific IP addresses or ranges of IP addresses.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafv2_createIpsetCmd).Standalone()

	wafv2_createIpsetCmd.Flags().String("addresses", "", "Contains an array of strings that specifies zero or more IP addresses or blocks of IP addresses that you want WAF to inspect for in incoming requests.")
	wafv2_createIpsetCmd.Flags().String("description", "", "A description of the IP set that helps with identification.")
	wafv2_createIpsetCmd.Flags().String("ipaddress-version", "", "The version of the IP addresses, either `IPV4` or `IPV6`.")
	wafv2_createIpsetCmd.Flags().String("name", "", "The name of the IP set.")
	wafv2_createIpsetCmd.Flags().String("scope", "", "Specifies whether this is for a global resource type, such as a Amazon CloudFront distribution.")
	wafv2_createIpsetCmd.Flags().String("tags", "", "An array of key:value pairs to associate with the resource.")
	wafv2_createIpsetCmd.MarkFlagRequired("addresses")
	wafv2_createIpsetCmd.MarkFlagRequired("ipaddress-version")
	wafv2_createIpsetCmd.MarkFlagRequired("name")
	wafv2_createIpsetCmd.MarkFlagRequired("scope")
	wafv2Cmd.AddCommand(wafv2_createIpsetCmd)
}

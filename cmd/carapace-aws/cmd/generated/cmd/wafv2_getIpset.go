package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafv2_getIpsetCmd = &cobra.Command{
	Use:   "get-ipset",
	Short: "Retrieves the specified [IPSet]().",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafv2_getIpsetCmd).Standalone()

	wafv2_getIpsetCmd.Flags().String("id", "", "A unique identifier for the set.")
	wafv2_getIpsetCmd.Flags().String("name", "", "The name of the IP set.")
	wafv2_getIpsetCmd.Flags().String("scope", "", "Specifies whether this is for a global resource type, such as a Amazon CloudFront distribution.")
	wafv2_getIpsetCmd.MarkFlagRequired("id")
	wafv2_getIpsetCmd.MarkFlagRequired("name")
	wafv2_getIpsetCmd.MarkFlagRequired("scope")
	wafv2Cmd.AddCommand(wafv2_getIpsetCmd)
}

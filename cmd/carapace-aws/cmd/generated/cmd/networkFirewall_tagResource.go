package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkFirewall_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds the specified tags to the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkFirewall_tagResourceCmd).Standalone()

	networkFirewall_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
	networkFirewall_tagResourceCmd.Flags().String("tags", "", "")
	networkFirewall_tagResourceCmd.MarkFlagRequired("resource-arn")
	networkFirewall_tagResourceCmd.MarkFlagRequired("tags")
	networkFirewallCmd.AddCommand(networkFirewall_tagResourceCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkFirewall_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes the tags with the specified keys from the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkFirewall_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkFirewall_untagResourceCmd).Standalone()

		networkFirewall_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
		networkFirewall_untagResourceCmd.Flags().String("tag-keys", "", "")
		networkFirewall_untagResourceCmd.MarkFlagRequired("resource-arn")
		networkFirewall_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	networkFirewallCmd.AddCommand(networkFirewall_untagResourceCmd)
}

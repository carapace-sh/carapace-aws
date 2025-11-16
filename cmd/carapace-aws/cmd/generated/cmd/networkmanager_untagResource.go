package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes tags from a specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkmanager_untagResourceCmd).Standalone()

		networkmanager_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
		networkmanager_untagResourceCmd.Flags().String("tag-keys", "", "The tag keys to remove from the specified resource.")
		networkmanager_untagResourceCmd.MarkFlagRequired("resource-arn")
		networkmanager_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	networkmanagerCmd.AddCommand(networkmanager_untagResourceCmd)
}

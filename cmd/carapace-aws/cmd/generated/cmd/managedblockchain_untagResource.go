package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var managedblockchain_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes the specified tags from the Amazon Managed Blockchain resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(managedblockchain_untagResourceCmd).Standalone()

	managedblockchain_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
	managedblockchain_untagResourceCmd.Flags().String("tag-keys", "", "The tag keys.")
	managedblockchain_untagResourceCmd.MarkFlagRequired("resource-arn")
	managedblockchain_untagResourceCmd.MarkFlagRequired("tag-keys")
	managedblockchainCmd.AddCommand(managedblockchain_untagResourceCmd)
}

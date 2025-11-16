package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transfer_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Detaches a key-value pair from a resource, as identified by its Amazon Resource Name (ARN).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transfer_untagResourceCmd).Standalone()

	transfer_untagResourceCmd.Flags().String("arn", "", "The value of the resource that will have the tag removed.")
	transfer_untagResourceCmd.Flags().String("tag-keys", "", "TagKeys are key-value pairs assigned to ARNs that can be used to group and search for resources by type.")
	transfer_untagResourceCmd.MarkFlagRequired("arn")
	transfer_untagResourceCmd.MarkFlagRequired("tag-keys")
	transferCmd.AddCommand(transfer_untagResourceCmd)
}

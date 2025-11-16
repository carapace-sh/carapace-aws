package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var braket_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Remove tags from a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(braket_untagResourceCmd).Standalone()

	braket_untagResourceCmd.Flags().String("resource-arn", "", "Specify the `resourceArn` for the resource from which to remove the tags.")
	braket_untagResourceCmd.Flags().String("tag-keys", "", "Specify the keys for the tags to remove from the resource.")
	braket_untagResourceCmd.MarkFlagRequired("resource-arn")
	braket_untagResourceCmd.MarkFlagRequired("tag-keys")
	braketCmd.AddCommand(braket_untagResourceCmd)
}

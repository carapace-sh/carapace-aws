package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var directconnect_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes one or more tags from the specified Direct Connect resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(directconnect_untagResourceCmd).Standalone()

	directconnect_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
	directconnect_untagResourceCmd.Flags().String("tag-keys", "", "The tag keys of the tags to remove.")
	directconnect_untagResourceCmd.MarkFlagRequired("resource-arn")
	directconnect_untagResourceCmd.MarkFlagRequired("tag-keys")
	directconnectCmd.AddCommand(directconnect_untagResourceCmd)
}

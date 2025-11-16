package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mpa_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes a resource tag.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mpa_untagResourceCmd).Standalone()

	mpa_untagResourceCmd.Flags().String("resource-arn", "", "Amazon Resource Name (ARN) for the resource you want to untag.")
	mpa_untagResourceCmd.Flags().String("tag-keys", "", "Array of tag key-value pairs that you want to untag.")
	mpa_untagResourceCmd.MarkFlagRequired("resource-arn")
	mpa_untagResourceCmd.MarkFlagRequired("tag-keys")
	mpaCmd.AddCommand(mpa_untagResourceCmd)
}

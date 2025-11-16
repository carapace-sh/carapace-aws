package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fsx_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "This action removes a tag from an Amazon FSx resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fsx_untagResourceCmd).Standalone()

	fsx_untagResourceCmd.Flags().String("resource-arn", "", "The ARN of the Amazon FSx resource to untag.")
	fsx_untagResourceCmd.Flags().String("tag-keys", "", "A list of keys of tags on the resource to untag.")
	fsx_untagResourceCmd.MarkFlagRequired("resource-arn")
	fsx_untagResourceCmd.MarkFlagRequired("tag-keys")
	fsxCmd.AddCommand(fsx_untagResourceCmd)
}

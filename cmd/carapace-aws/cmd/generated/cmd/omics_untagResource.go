package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes tags from a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_untagResourceCmd).Standalone()

	omics_untagResourceCmd.Flags().String("resource-arn", "", "The resource's ARN.")
	omics_untagResourceCmd.Flags().String("tag-keys", "", "Keys of tags to remove.")
	omics_untagResourceCmd.MarkFlagRequired("resource-arn")
	omics_untagResourceCmd.MarkFlagRequired("tag-keys")
	omicsCmd.AddCommand(omics_untagResourceCmd)
}

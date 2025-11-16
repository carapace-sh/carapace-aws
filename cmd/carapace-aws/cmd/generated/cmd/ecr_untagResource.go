package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecr_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Deletes specified tags from a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecr_untagResourceCmd).Standalone()

	ecr_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource from which to remove tags.")
	ecr_untagResourceCmd.Flags().String("tag-keys", "", "The keys of the tags to be removed.")
	ecr_untagResourceCmd.MarkFlagRequired("resource-arn")
	ecr_untagResourceCmd.MarkFlagRequired("tag-keys")
	ecrCmd.AddCommand(ecr_untagResourceCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecr_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "List the tags for an Amazon ECR resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecr_listTagsForResourceCmd).Standalone()

	ecr_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) that identifies the resource for which to list the tags.")
	ecr_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	ecrCmd.AddCommand(ecr_listTagsForResourceCmd)
}

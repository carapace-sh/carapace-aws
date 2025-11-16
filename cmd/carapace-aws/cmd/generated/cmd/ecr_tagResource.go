package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecr_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds specified tags to a resource with the specified ARN.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecr_tagResourceCmd).Standalone()

	ecr_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the the resource to which to add tags.")
	ecr_tagResourceCmd.Flags().String("tags", "", "The tags to add to the resource.")
	ecr_tagResourceCmd.MarkFlagRequired("resource-arn")
	ecr_tagResourceCmd.MarkFlagRequired("tags")
	ecrCmd.AddCommand(ecr_tagResourceCmd)
}

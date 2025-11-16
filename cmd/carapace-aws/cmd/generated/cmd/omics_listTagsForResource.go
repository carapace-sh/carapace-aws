package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Retrieves a list of tags for a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_listTagsForResourceCmd).Standalone()

	omics_listTagsForResourceCmd.Flags().String("resource-arn", "", "The resource's ARN.")
	omics_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	omicsCmd.AddCommand(omics_listTagsForResourceCmd)
}

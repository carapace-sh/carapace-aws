package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var entityresolution_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Displays the tags associated with an Entity Resolution resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(entityresolution_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(entityresolution_listTagsForResourceCmd).Standalone()

		entityresolution_listTagsForResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource for which you want to view tags.")
		entityresolution_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	entityresolutionCmd.AddCommand(entityresolution_listTagsForResourceCmd)
}

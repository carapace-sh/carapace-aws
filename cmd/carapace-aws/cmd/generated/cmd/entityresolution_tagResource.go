package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var entityresolution_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Assigns one or more tags (key-value pairs) to the specified Entity Resolution resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(entityresolution_tagResourceCmd).Standalone()

	entityresolution_tagResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource for which you want to view tags.")
	entityresolution_tagResourceCmd.Flags().String("tags", "", "The tags used to organize, track, or control access for this resource.")
	entityresolution_tagResourceCmd.MarkFlagRequired("resource-arn")
	entityresolution_tagResourceCmd.MarkFlagRequired("tags")
	entityresolutionCmd.AddCommand(entityresolution_tagResourceCmd)
}

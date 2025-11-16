package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var accessanalyzer_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Retrieves a list of tags applied to the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(accessanalyzer_listTagsForResourceCmd).Standalone()

	accessanalyzer_listTagsForResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource to retrieve tags from.")
	accessanalyzer_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	accessanalyzerCmd.AddCommand(accessanalyzer_listTagsForResourceCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var accessanalyzer_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds a tag to the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(accessanalyzer_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(accessanalyzer_tagResourceCmd).Standalone()

		accessanalyzer_tagResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource to add the tag to.")
		accessanalyzer_tagResourceCmd.Flags().String("tags", "", "The tags to add to the resource.")
		accessanalyzer_tagResourceCmd.MarkFlagRequired("resource-arn")
		accessanalyzer_tagResourceCmd.MarkFlagRequired("tags")
	})
	accessanalyzerCmd.AddCommand(accessanalyzer_tagResourceCmd)
}

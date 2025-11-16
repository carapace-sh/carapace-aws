package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeartifact_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Gets information about Amazon Web Services tags for a specified Amazon Resource Name (ARN) in CodeArtifact.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeartifact_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codeartifact_listTagsForResourceCmd).Standalone()

		codeartifact_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource to get tags for.")
		codeartifact_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	codeartifactCmd.AddCommand(codeartifact_listTagsForResourceCmd)
}

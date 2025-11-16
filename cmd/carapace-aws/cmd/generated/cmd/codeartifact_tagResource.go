package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeartifact_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds or updates tags for a resource in CodeArtifact.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeartifact_tagResourceCmd).Standalone()

	codeartifact_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource that you want to add or update tags for.")
	codeartifact_tagResourceCmd.Flags().String("tags", "", "The tags you want to modify or add to the resource.")
	codeartifact_tagResourceCmd.MarkFlagRequired("resource-arn")
	codeartifact_tagResourceCmd.MarkFlagRequired("tags")
	codeartifactCmd.AddCommand(codeartifact_tagResourceCmd)
}

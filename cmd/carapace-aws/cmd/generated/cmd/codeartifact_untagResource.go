package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeartifact_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes tags from a resource in CodeArtifact.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeartifact_untagResourceCmd).Standalone()

	codeartifact_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource that you want to remove tags from.")
	codeartifact_untagResourceCmd.Flags().String("tag-keys", "", "The tag key for each tag that you want to remove from the resource.")
	codeartifact_untagResourceCmd.MarkFlagRequired("resource-arn")
	codeartifact_untagResourceCmd.MarkFlagRequired("tag-keys")
	codeartifactCmd.AddCommand(codeartifact_untagResourceCmd)
}

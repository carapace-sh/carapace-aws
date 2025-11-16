package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var accessanalyzer_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes a tag from the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(accessanalyzer_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(accessanalyzer_untagResourceCmd).Standalone()

		accessanalyzer_untagResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource to remove the tag from.")
		accessanalyzer_untagResourceCmd.Flags().String("tag-keys", "", "The key for the tag to add.")
		accessanalyzer_untagResourceCmd.MarkFlagRequired("resource-arn")
		accessanalyzer_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	accessanalyzerCmd.AddCommand(accessanalyzer_untagResourceCmd)
}

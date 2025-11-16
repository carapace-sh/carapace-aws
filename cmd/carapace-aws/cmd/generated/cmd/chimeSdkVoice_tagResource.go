package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds a tag to the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkVoice_tagResourceCmd).Standalone()

		chimeSdkVoice_tagResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource being tagged.")
		chimeSdkVoice_tagResourceCmd.Flags().String("tags", "", "A list of the tags being added to the resource.")
		chimeSdkVoice_tagResourceCmd.MarkFlagRequired("resource-arn")
		chimeSdkVoice_tagResourceCmd.MarkFlagRequired("tags")
	})
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_tagResourceCmd)
}

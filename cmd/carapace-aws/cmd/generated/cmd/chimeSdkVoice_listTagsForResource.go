package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Returns a list of the tags in a given resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_listTagsForResourceCmd).Standalone()

	chimeSdkVoice_listTagsForResourceCmd.Flags().String("resource-arn", "", "The resource ARN.")
	chimeSdkVoice_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_listTagsForResourceCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds one or more tags (keys and values) to an application, campaign, message template, or segment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpoint_tagResourceCmd).Standalone()

		pinpoint_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
		pinpoint_tagResourceCmd.Flags().String("tags-model", "", "")
		pinpoint_tagResourceCmd.MarkFlagRequired("resource-arn")
		pinpoint_tagResourceCmd.MarkFlagRequired("tags-model")
	})
	pinpointCmd.AddCommand(pinpoint_tagResourceCmd)
}

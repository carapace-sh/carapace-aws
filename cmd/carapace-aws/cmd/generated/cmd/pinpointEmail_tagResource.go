package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointEmail_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Add one or more tags (keys and values) to a specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointEmail_tagResourceCmd).Standalone()

	pinpointEmail_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource that you want to add one or more tags to.")
	pinpointEmail_tagResourceCmd.Flags().String("tags", "", "A list of the tags that you want to add to the resource.")
	pinpointEmail_tagResourceCmd.MarkFlagRequired("resource-arn")
	pinpointEmail_tagResourceCmd.MarkFlagRequired("tags")
	pinpointEmailCmd.AddCommand(pinpointEmail_tagResourceCmd)
}

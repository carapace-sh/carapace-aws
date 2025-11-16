package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointEmail_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Retrieve a list of the tags (keys and values) that are associated with a specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointEmail_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpointEmail_listTagsForResourceCmd).Standalone()

		pinpointEmail_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource that you want to retrieve tag information for.")
		pinpointEmail_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	pinpointEmailCmd.AddCommand(pinpointEmail_listTagsForResourceCmd)
}

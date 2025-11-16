package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists the tags (metadata) you have assigned to the resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_listTagsForResourceCmd).Standalone()

		iot_listTagsForResourceCmd.Flags().String("next-token", "", "To retrieve the next set of results, the `nextToken` value from a previous response; otherwise **null** to receive the first set of results.")
		iot_listTagsForResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource.")
		iot_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	iotCmd.AddCommand(iot_listTagsForResourceCmd)
}

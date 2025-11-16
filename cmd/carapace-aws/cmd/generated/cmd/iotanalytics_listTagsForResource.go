package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotanalytics_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists the tags (metadata) that you have assigned to the resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotanalytics_listTagsForResourceCmd).Standalone()

	iotanalytics_listTagsForResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource whose tags you want to list.")
	iotanalytics_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	iotanalyticsCmd.AddCommand(iotanalytics_listTagsForResourceCmd)
}

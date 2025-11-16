package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotanalytics_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds to or modifies the tags of the given resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotanalytics_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotanalytics_tagResourceCmd).Standalone()

		iotanalytics_tagResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource whose tags you want to modify.")
		iotanalytics_tagResourceCmd.Flags().String("tags", "", "The new or modified tags for the resource.")
		iotanalytics_tagResourceCmd.MarkFlagRequired("resource-arn")
		iotanalytics_tagResourceCmd.MarkFlagRequired("tags")
	})
	iotanalyticsCmd.AddCommand(iotanalytics_tagResourceCmd)
}

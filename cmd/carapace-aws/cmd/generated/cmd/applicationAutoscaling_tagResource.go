package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationAutoscaling_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds or edits tags on an Application Auto Scaling scalable target.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationAutoscaling_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(applicationAutoscaling_tagResourceCmd).Standalone()

		applicationAutoscaling_tagResourceCmd.Flags().String("resource-arn", "", "Identifies the Application Auto Scaling scalable target that you want to apply tags to.")
		applicationAutoscaling_tagResourceCmd.Flags().String("tags", "", "The tags assigned to the resource.")
		applicationAutoscaling_tagResourceCmd.MarkFlagRequired("resource-arn")
		applicationAutoscaling_tagResourceCmd.MarkFlagRequired("tags")
	})
	applicationAutoscalingCmd.AddCommand(applicationAutoscaling_tagResourceCmd)
}

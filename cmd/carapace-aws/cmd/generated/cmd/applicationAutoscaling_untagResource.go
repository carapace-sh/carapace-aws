package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationAutoscaling_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Deletes tags from an Application Auto Scaling scalable target.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationAutoscaling_untagResourceCmd).Standalone()

	applicationAutoscaling_untagResourceCmd.Flags().String("resource-arn", "", "Identifies the Application Auto Scaling scalable target from which to remove tags.")
	applicationAutoscaling_untagResourceCmd.Flags().String("tag-keys", "", "One or more tag keys.")
	applicationAutoscaling_untagResourceCmd.MarkFlagRequired("resource-arn")
	applicationAutoscaling_untagResourceCmd.MarkFlagRequired("tag-keys")
	applicationAutoscalingCmd.AddCommand(applicationAutoscaling_untagResourceCmd)
}

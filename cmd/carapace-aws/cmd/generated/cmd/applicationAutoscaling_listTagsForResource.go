package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationAutoscaling_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Returns all the tags on the specified Application Auto Scaling scalable target.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationAutoscaling_listTagsForResourceCmd).Standalone()

	applicationAutoscaling_listTagsForResourceCmd.Flags().String("resource-arn", "", "Specify the ARN of the scalable target.")
	applicationAutoscaling_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	applicationAutoscalingCmd.AddCommand(applicationAutoscaling_listTagsForResourceCmd)
}

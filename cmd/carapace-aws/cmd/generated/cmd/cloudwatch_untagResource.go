package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudwatch_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes one or more tags from the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudwatch_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudwatch_untagResourceCmd).Standalone()

		cloudwatch_untagResourceCmd.Flags().String("resource-arn", "", "The ARN of the CloudWatch resource that you're removing tags from.")
		cloudwatch_untagResourceCmd.Flags().String("tag-keys", "", "The list of tag keys to remove from the resource.")
		cloudwatch_untagResourceCmd.MarkFlagRequired("resource-arn")
		cloudwatch_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	cloudwatchCmd.AddCommand(cloudwatch_untagResourceCmd)
}

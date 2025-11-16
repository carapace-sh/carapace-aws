package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudwatch_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Assigns one or more tags (key-value pairs) to the specified CloudWatch resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudwatch_tagResourceCmd).Standalone()

	cloudwatch_tagResourceCmd.Flags().String("resource-arn", "", "The ARN of the CloudWatch resource that you're adding tags to.")
	cloudwatch_tagResourceCmd.Flags().String("tags", "", "The list of key-value pairs to associate with the alarm.")
	cloudwatch_tagResourceCmd.MarkFlagRequired("resource-arn")
	cloudwatch_tagResourceCmd.MarkFlagRequired("tags")
	cloudwatchCmd.AddCommand(cloudwatch_tagResourceCmd)
}

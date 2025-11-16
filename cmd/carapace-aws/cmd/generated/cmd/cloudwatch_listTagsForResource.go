package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudwatch_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Displays the tags associated with a CloudWatch resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudwatch_listTagsForResourceCmd).Standalone()

	cloudwatch_listTagsForResourceCmd.Flags().String("resource-arn", "", "The ARN of the CloudWatch resource that you want to view tags for.")
	cloudwatch_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	cloudwatchCmd.AddCommand(cloudwatch_listTagsForResourceCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Displays the tags associated with a CloudWatch Logs resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(logs_listTagsForResourceCmd).Standalone()

		logs_listTagsForResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource that you want to view tags for.")
		logs_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	logsCmd.AddCommand(logs_listTagsForResourceCmd)
}

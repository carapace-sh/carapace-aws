package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Assigns one or more tags (key-value pairs) to the specified CloudWatch Logs resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(logs_tagResourceCmd).Standalone()

		logs_tagResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource that you're adding tags to.")
		logs_tagResourceCmd.Flags().String("tags", "", "The list of key-value pairs to associate with the resource.")
		logs_tagResourceCmd.MarkFlagRequired("resource-arn")
		logs_tagResourceCmd.MarkFlagRequired("tags")
	})
	logsCmd.AddCommand(logs_tagResourceCmd)
}

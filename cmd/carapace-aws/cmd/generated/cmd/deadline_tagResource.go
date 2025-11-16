package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Tags a resource using the resource's ARN and desired tags.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(deadline_tagResourceCmd).Standalone()

		deadline_tagResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource to apply tags to.")
		deadline_tagResourceCmd.Flags().String("tags", "", "Each tag consists of a tag key and a tag value.")
		deadline_tagResourceCmd.MarkFlagRequired("resource-arn")
	})
	deadlineCmd.AddCommand(deadline_tagResourceCmd)
}

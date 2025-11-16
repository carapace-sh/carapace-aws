package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stepfunctions_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Add a tag to a Step Functions resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stepfunctions_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(stepfunctions_tagResourceCmd).Standalone()

		stepfunctions_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) for the Step Functions state machine or activity.")
		stepfunctions_tagResourceCmd.Flags().String("tags", "", "The list of tags to add to a resource.")
		stepfunctions_tagResourceCmd.MarkFlagRequired("resource-arn")
		stepfunctions_tagResourceCmd.MarkFlagRequired("tags")
	})
	stepfunctionsCmd.AddCommand(stepfunctions_tagResourceCmd)
}

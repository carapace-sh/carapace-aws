package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stepfunctions_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "List tags for a given resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stepfunctions_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(stepfunctions_listTagsForResourceCmd).Standalone()

		stepfunctions_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) for the Step Functions state machine or activity.")
		stepfunctions_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	stepfunctionsCmd.AddCommand(stepfunctions_listTagsForResourceCmd)
}

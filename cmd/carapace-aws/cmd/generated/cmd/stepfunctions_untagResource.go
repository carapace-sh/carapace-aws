package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stepfunctions_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Remove a tag from a Step Functions resource",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stepfunctions_untagResourceCmd).Standalone()

	stepfunctions_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) for the Step Functions state machine or activity.")
	stepfunctions_untagResourceCmd.Flags().String("tag-keys", "", "The list of tags to remove from the resource.")
	stepfunctions_untagResourceCmd.MarkFlagRequired("resource-arn")
	stepfunctions_untagResourceCmd.MarkFlagRequired("tag-keys")
	stepfunctionsCmd.AddCommand(stepfunctions_untagResourceCmd)
}

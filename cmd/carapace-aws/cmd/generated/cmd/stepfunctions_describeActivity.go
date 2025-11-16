package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stepfunctions_describeActivityCmd = &cobra.Command{
	Use:   "describe-activity",
	Short: "Describes an activity.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stepfunctions_describeActivityCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(stepfunctions_describeActivityCmd).Standalone()

		stepfunctions_describeActivityCmd.Flags().String("activity-arn", "", "The Amazon Resource Name (ARN) of the activity to describe.")
		stepfunctions_describeActivityCmd.MarkFlagRequired("activity-arn")
	})
	stepfunctionsCmd.AddCommand(stepfunctions_describeActivityCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stepfunctions_deleteActivityCmd = &cobra.Command{
	Use:   "delete-activity",
	Short: "Deletes an activity.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stepfunctions_deleteActivityCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(stepfunctions_deleteActivityCmd).Standalone()

		stepfunctions_deleteActivityCmd.Flags().String("activity-arn", "", "The Amazon Resource Name (ARN) of the activity to delete.")
		stepfunctions_deleteActivityCmd.MarkFlagRequired("activity-arn")
	})
	stepfunctionsCmd.AddCommand(stepfunctions_deleteActivityCmd)
}

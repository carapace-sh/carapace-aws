package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var swf_respondActivityTaskCompletedCmd = &cobra.Command{
	Use:   "respond-activity-task-completed",
	Short: "Used by workers to tell the service that the [ActivityTask]() identified by the `taskToken` completed successfully with a `result` (if provided).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(swf_respondActivityTaskCompletedCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(swf_respondActivityTaskCompletedCmd).Standalone()

		swf_respondActivityTaskCompletedCmd.Flags().String("result", "", "The result of the activity task.")
		swf_respondActivityTaskCompletedCmd.Flags().String("task-token", "", "The `taskToken` of the [ActivityTask]().")
		swf_respondActivityTaskCompletedCmd.MarkFlagRequired("task-token")
	})
	swfCmd.AddCommand(swf_respondActivityTaskCompletedCmd)
}

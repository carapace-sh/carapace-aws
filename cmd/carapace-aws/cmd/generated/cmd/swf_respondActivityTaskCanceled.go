package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var swf_respondActivityTaskCanceledCmd = &cobra.Command{
	Use:   "respond-activity-task-canceled",
	Short: "Used by workers to tell the service that the [ActivityTask]() identified by the `taskToken` was successfully canceled.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(swf_respondActivityTaskCanceledCmd).Standalone()

	swf_respondActivityTaskCanceledCmd.Flags().String("details", "", "Information about the cancellation.")
	swf_respondActivityTaskCanceledCmd.Flags().String("task-token", "", "The `taskToken` of the [ActivityTask]().")
	swf_respondActivityTaskCanceledCmd.MarkFlagRequired("task-token")
	swfCmd.AddCommand(swf_respondActivityTaskCanceledCmd)
}

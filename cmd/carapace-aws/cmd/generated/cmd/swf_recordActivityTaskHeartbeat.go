package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var swf_recordActivityTaskHeartbeatCmd = &cobra.Command{
	Use:   "record-activity-task-heartbeat",
	Short: "Used by activity workers to report to the service that the [ActivityTask]() represented by the specified `taskToken` is still making progress.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(swf_recordActivityTaskHeartbeatCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(swf_recordActivityTaskHeartbeatCmd).Standalone()

		swf_recordActivityTaskHeartbeatCmd.Flags().String("details", "", "If specified, contains details about the progress of the task.")
		swf_recordActivityTaskHeartbeatCmd.Flags().String("task-token", "", "The `taskToken` of the [ActivityTask]().")
		swf_recordActivityTaskHeartbeatCmd.MarkFlagRequired("task-token")
	})
	swfCmd.AddCommand(swf_recordActivityTaskHeartbeatCmd)
}

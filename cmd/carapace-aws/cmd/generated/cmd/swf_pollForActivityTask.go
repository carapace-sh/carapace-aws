package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var swf_pollForActivityTaskCmd = &cobra.Command{
	Use:   "poll-for-activity-task",
	Short: "Used by workers to get an [ActivityTask]() from the specified activity `taskList`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(swf_pollForActivityTaskCmd).Standalone()

	swf_pollForActivityTaskCmd.Flags().String("domain", "", "The name of the domain that contains the task lists being polled.")
	swf_pollForActivityTaskCmd.Flags().String("identity", "", "Identity of the worker making the request, recorded in the `ActivityTaskStarted` event in the workflow history.")
	swf_pollForActivityTaskCmd.Flags().String("task-list", "", "Specifies the task list to poll for activity tasks.")
	swf_pollForActivityTaskCmd.MarkFlagRequired("domain")
	swf_pollForActivityTaskCmd.MarkFlagRequired("task-list")
	swfCmd.AddCommand(swf_pollForActivityTaskCmd)
}

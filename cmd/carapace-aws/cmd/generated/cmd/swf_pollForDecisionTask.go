package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var swf_pollForDecisionTaskCmd = &cobra.Command{
	Use:   "poll-for-decision-task",
	Short: "Used by deciders to get a [DecisionTask]() from the specified decision `taskList`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(swf_pollForDecisionTaskCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(swf_pollForDecisionTaskCmd).Standalone()

		swf_pollForDecisionTaskCmd.Flags().String("domain", "", "The name of the domain containing the task lists to poll.")
		swf_pollForDecisionTaskCmd.Flags().String("identity", "", "Identity of the decider making the request, which is recorded in the DecisionTaskStarted event in the workflow history.")
		swf_pollForDecisionTaskCmd.Flags().String("maximum-page-size", "", "The maximum number of results that are returned per call.")
		swf_pollForDecisionTaskCmd.Flags().String("next-page-token", "", "If `NextPageToken` is returned there are more results available.")
		swf_pollForDecisionTaskCmd.Flags().String("reverse-order", "", "When set to `true`, returns the events in reverse order.")
		swf_pollForDecisionTaskCmd.Flags().String("start-at-previous-started-event", "", "When set to `true`, returns the events with `eventTimestamp` greater than or equal to `eventTimestamp` of the most recent `DecisionTaskStarted` event.")
		swf_pollForDecisionTaskCmd.Flags().String("task-list", "", "Specifies the task list to poll for decision tasks.")
		swf_pollForDecisionTaskCmd.MarkFlagRequired("domain")
		swf_pollForDecisionTaskCmd.MarkFlagRequired("task-list")
	})
	swfCmd.AddCommand(swf_pollForDecisionTaskCmd)
}

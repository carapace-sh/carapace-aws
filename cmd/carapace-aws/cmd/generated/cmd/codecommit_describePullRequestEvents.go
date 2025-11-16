package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecommit_describePullRequestEventsCmd = &cobra.Command{
	Use:   "describe-pull-request-events",
	Short: "Returns information about one or more pull request events.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecommit_describePullRequestEventsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codecommit_describePullRequestEventsCmd).Standalone()

		codecommit_describePullRequestEventsCmd.Flags().String("actor-arn", "", "The Amazon Resource Name (ARN) of the user whose actions resulted in the event.")
		codecommit_describePullRequestEventsCmd.Flags().String("max-results", "", "A non-zero, non-negative integer used to limit the number of returned results.")
		codecommit_describePullRequestEventsCmd.Flags().String("next-token", "", "An enumeration token that, when provided in a request, returns the next batch of the results.")
		codecommit_describePullRequestEventsCmd.Flags().String("pull-request-event-type", "", "Optional.")
		codecommit_describePullRequestEventsCmd.Flags().String("pull-request-id", "", "The system-generated ID of the pull request.")
		codecommit_describePullRequestEventsCmd.MarkFlagRequired("pull-request-id")
	})
	codecommitCmd.AddCommand(codecommit_describePullRequestEventsCmd)
}

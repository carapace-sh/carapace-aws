package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_describeSessionsCmd = &cobra.Command{
	Use:   "describe-sessions",
	Short: "Retrieves a list of all active sessions (both connected and disconnected) or terminated sessions from the past 30 days.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_describeSessionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssm_describeSessionsCmd).Standalone()

		ssm_describeSessionsCmd.Flags().String("filters", "", "One or more filters to limit the type of sessions returned by the request.")
		ssm_describeSessionsCmd.Flags().String("max-results", "", "The maximum number of items to return for this call.")
		ssm_describeSessionsCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
		ssm_describeSessionsCmd.Flags().String("state", "", "The session status to retrieve a list of sessions for.")
		ssm_describeSessionsCmd.MarkFlagRequired("state")
	})
	ssmCmd.AddCommand(ssm_describeSessionsCmd)
}

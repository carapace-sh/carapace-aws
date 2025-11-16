package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_describeSessionsCmd = &cobra.Command{
	Use:   "describe-sessions",
	Short: "Retrieves a list that describes the streaming sessions for a specified stack and fleet.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_describeSessionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appstream_describeSessionsCmd).Standalone()

		appstream_describeSessionsCmd.Flags().String("authentication-type", "", "The authentication method.")
		appstream_describeSessionsCmd.Flags().String("fleet-name", "", "The name of the fleet.")
		appstream_describeSessionsCmd.Flags().String("instance-id", "", "The identifier for the instance hosting the session.")
		appstream_describeSessionsCmd.Flags().String("limit", "", "The size of each page of results.")
		appstream_describeSessionsCmd.Flags().String("next-token", "", "The pagination token to use to retrieve the next page of results for this operation.")
		appstream_describeSessionsCmd.Flags().String("stack-name", "", "The name of the stack.")
		appstream_describeSessionsCmd.Flags().String("user-id", "", "The user identifier (ID).")
		appstream_describeSessionsCmd.MarkFlagRequired("fleet-name")
		appstream_describeSessionsCmd.MarkFlagRequired("stack-name")
	})
	appstreamCmd.AddCommand(appstream_describeSessionsCmd)
}

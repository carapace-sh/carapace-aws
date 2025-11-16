package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_describeAccountModificationsCmd = &cobra.Command{
	Use:   "describe-account-modifications",
	Short: "Retrieves a list that describes modifications to the configuration of Bring Your Own License (BYOL) for the specified account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_describeAccountModificationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspaces_describeAccountModificationsCmd).Standalone()

		workspaces_describeAccountModificationsCmd.Flags().String("next-token", "", "If you received a `NextToken` from a previous call that was paginated, provide this token to receive the next set of results.")
	})
	workspacesCmd.AddCommand(workspaces_describeAccountModificationsCmd)
}

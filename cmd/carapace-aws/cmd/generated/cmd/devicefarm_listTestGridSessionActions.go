package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicefarm_listTestGridSessionActionsCmd = &cobra.Command{
	Use:   "list-test-grid-session-actions",
	Short: "Returns a list of the actions taken in a [TestGridSession]().",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicefarm_listTestGridSessionActionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(devicefarm_listTestGridSessionActionsCmd).Standalone()

		devicefarm_listTestGridSessionActionsCmd.Flags().String("max-result", "", "The maximum number of sessions to return per response.")
		devicefarm_listTestGridSessionActionsCmd.Flags().String("next-token", "", "Pagination token.")
		devicefarm_listTestGridSessionActionsCmd.Flags().String("session-arn", "", "The ARN of the session to retrieve.")
		devicefarm_listTestGridSessionActionsCmd.MarkFlagRequired("session-arn")
	})
	devicefarmCmd.AddCommand(devicefarm_listTestGridSessionActionsCmd)
}

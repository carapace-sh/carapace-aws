package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicefarm_listTestGridSessionsCmd = &cobra.Command{
	Use:   "list-test-grid-sessions",
	Short: "Retrieves a list of sessions for a [TestGridProject]().",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicefarm_listTestGridSessionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(devicefarm_listTestGridSessionsCmd).Standalone()

		devicefarm_listTestGridSessionsCmd.Flags().String("creation-time-after", "", "Return only sessions created after this time.")
		devicefarm_listTestGridSessionsCmd.Flags().String("creation-time-before", "", "Return only sessions created before this time.")
		devicefarm_listTestGridSessionsCmd.Flags().String("end-time-after", "", "Return only sessions that ended after this time.")
		devicefarm_listTestGridSessionsCmd.Flags().String("end-time-before", "", "Return only sessions that ended before this time.")
		devicefarm_listTestGridSessionsCmd.Flags().String("max-result", "", "Return only this many results at a time.")
		devicefarm_listTestGridSessionsCmd.Flags().String("next-token", "", "Pagination token.")
		devicefarm_listTestGridSessionsCmd.Flags().String("project-arn", "", "ARN of a [TestGridProject]().")
		devicefarm_listTestGridSessionsCmd.Flags().String("status", "", "Return only sessions in this state.")
		devicefarm_listTestGridSessionsCmd.MarkFlagRequired("project-arn")
	})
	devicefarmCmd.AddCommand(devicefarm_listTestGridSessionsCmd)
}

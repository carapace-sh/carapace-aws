package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_deleteFleetAdvisorDatabasesCmd = &cobra.Command{
	Use:   "delete-fleet-advisor-databases",
	Short: "End of support notice: On May 20, 2026, Amazon Web Services will end support for Amazon Web Services DMS Fleet Advisor;.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_deleteFleetAdvisorDatabasesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dms_deleteFleetAdvisorDatabasesCmd).Standalone()

		dms_deleteFleetAdvisorDatabasesCmd.Flags().String("database-ids", "", "The IDs of the Fleet Advisor collector databases to delete.")
		dms_deleteFleetAdvisorDatabasesCmd.MarkFlagRequired("database-ids")
	})
	dmsCmd.AddCommand(dms_deleteFleetAdvisorDatabasesCmd)
}

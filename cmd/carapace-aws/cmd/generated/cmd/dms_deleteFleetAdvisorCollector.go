package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_deleteFleetAdvisorCollectorCmd = &cobra.Command{
	Use:   "delete-fleet-advisor-collector",
	Short: "End of support notice: On May 20, 2026, Amazon Web Services will end support for Amazon Web Services DMS Fleet Advisor;.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_deleteFleetAdvisorCollectorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dms_deleteFleetAdvisorCollectorCmd).Standalone()

		dms_deleteFleetAdvisorCollectorCmd.Flags().String("collector-referenced-id", "", "The reference ID of the Fleet Advisor collector to delete.")
		dms_deleteFleetAdvisorCollectorCmd.MarkFlagRequired("collector-referenced-id")
	})
	dmsCmd.AddCommand(dms_deleteFleetAdvisorCollectorCmd)
}

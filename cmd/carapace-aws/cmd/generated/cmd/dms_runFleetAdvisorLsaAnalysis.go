package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_runFleetAdvisorLsaAnalysisCmd = &cobra.Command{
	Use:   "run-fleet-advisor-lsa-analysis",
	Short: "End of support notice: On May 20, 2026, Amazon Web Services will end support for Amazon Web Services DMS Fleet Advisor;.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_runFleetAdvisorLsaAnalysisCmd).Standalone()

	dmsCmd.AddCommand(dms_runFleetAdvisorLsaAnalysisCmd)
}

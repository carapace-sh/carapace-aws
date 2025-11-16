package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var healthlake_describeFhirexportJobCmd = &cobra.Command{
	Use:   "describe-fhirexport-job",
	Short: "Get FHIR export job properties.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(healthlake_describeFhirexportJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(healthlake_describeFhirexportJobCmd).Standalone()

		healthlake_describeFhirexportJobCmd.Flags().String("datastore-id", "", "The data store identifier from which FHIR data is being exported from.")
		healthlake_describeFhirexportJobCmd.Flags().String("job-id", "", "The export job identifier.")
		healthlake_describeFhirexportJobCmd.MarkFlagRequired("datastore-id")
		healthlake_describeFhirexportJobCmd.MarkFlagRequired("job-id")
	})
	healthlakeCmd.AddCommand(healthlake_describeFhirexportJobCmd)
}

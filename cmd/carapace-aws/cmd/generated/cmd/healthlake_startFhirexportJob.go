package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var healthlake_startFhirexportJobCmd = &cobra.Command{
	Use:   "start-fhirexport-job",
	Short: "Start a FHIR export job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(healthlake_startFhirexportJobCmd).Standalone()

	healthlake_startFhirexportJobCmd.Flags().String("client-token", "", "An optional user provided token used for ensuring API idempotency.")
	healthlake_startFhirexportJobCmd.Flags().String("data-access-role-arn", "", "The Amazon Resource Name (ARN) used during initiation of the export job.")
	healthlake_startFhirexportJobCmd.Flags().String("datastore-id", "", "The data store identifier from which files are being exported.")
	healthlake_startFhirexportJobCmd.Flags().String("job-name", "", "The export job name.")
	healthlake_startFhirexportJobCmd.Flags().String("output-data-config", "", "The output data configuration supplied when the export job was started.")
	healthlake_startFhirexportJobCmd.MarkFlagRequired("data-access-role-arn")
	healthlake_startFhirexportJobCmd.MarkFlagRequired("datastore-id")
	healthlake_startFhirexportJobCmd.MarkFlagRequired("output-data-config")
	healthlakeCmd.AddCommand(healthlake_startFhirexportJobCmd)
}

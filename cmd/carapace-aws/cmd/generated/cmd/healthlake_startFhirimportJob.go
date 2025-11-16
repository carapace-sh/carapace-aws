package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var healthlake_startFhirimportJobCmd = &cobra.Command{
	Use:   "start-fhirimport-job",
	Short: "Start importing bulk FHIR data into an ACTIVE data store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(healthlake_startFhirimportJobCmd).Standalone()

	healthlake_startFhirimportJobCmd.Flags().String("client-token", "", "The optional user-provided token used for ensuring API idempotency.")
	healthlake_startFhirimportJobCmd.Flags().String("data-access-role-arn", "", "The Amazon Resource Name (ARN) that grants access permission to AWS HealthLake.")
	healthlake_startFhirimportJobCmd.Flags().String("datastore-id", "", "The data store identifier.")
	healthlake_startFhirimportJobCmd.Flags().String("input-data-config", "", "The input properties for the import job request.")
	healthlake_startFhirimportJobCmd.Flags().String("job-name", "", "The import job name.")
	healthlake_startFhirimportJobCmd.Flags().String("job-output-data-config", "", "")
	healthlake_startFhirimportJobCmd.Flags().String("validation-level", "", "The validation level of the import job.")
	healthlake_startFhirimportJobCmd.MarkFlagRequired("data-access-role-arn")
	healthlake_startFhirimportJobCmd.MarkFlagRequired("datastore-id")
	healthlake_startFhirimportJobCmd.MarkFlagRequired("input-data-config")
	healthlake_startFhirimportJobCmd.MarkFlagRequired("job-output-data-config")
	healthlakeCmd.AddCommand(healthlake_startFhirimportJobCmd)
}

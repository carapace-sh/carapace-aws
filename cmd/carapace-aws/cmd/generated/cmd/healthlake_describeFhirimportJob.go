package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var healthlake_describeFhirimportJobCmd = &cobra.Command{
	Use:   "describe-fhirimport-job",
	Short: "Get the import job properties to learn more about the job or job progress.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(healthlake_describeFhirimportJobCmd).Standalone()

	healthlake_describeFhirimportJobCmd.Flags().String("datastore-id", "", "The data store identifier.")
	healthlake_describeFhirimportJobCmd.Flags().String("job-id", "", "The import job identifier.")
	healthlake_describeFhirimportJobCmd.MarkFlagRequired("datastore-id")
	healthlake_describeFhirimportJobCmd.MarkFlagRequired("job-id")
	healthlakeCmd.AddCommand(healthlake_describeFhirimportJobCmd)
}

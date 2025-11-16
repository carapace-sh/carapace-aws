package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medicalImaging_getDicomimportJobCmd = &cobra.Command{
	Use:   "get-dicomimport-job",
	Short: "Get the import job properties to learn more about the job or job progress.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medicalImaging_getDicomimportJobCmd).Standalone()

	medicalImaging_getDicomimportJobCmd.Flags().String("datastore-id", "", "The data store identifier.")
	medicalImaging_getDicomimportJobCmd.Flags().String("job-id", "", "The import job identifier.")
	medicalImaging_getDicomimportJobCmd.MarkFlagRequired("datastore-id")
	medicalImaging_getDicomimportJobCmd.MarkFlagRequired("job-id")
	medicalImagingCmd.AddCommand(medicalImaging_getDicomimportJobCmd)
}

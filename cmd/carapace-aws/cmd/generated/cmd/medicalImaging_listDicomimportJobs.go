package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medicalImaging_listDicomimportJobsCmd = &cobra.Command{
	Use:   "list-dicomimport-jobs",
	Short: "List import jobs created for a specific data store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medicalImaging_listDicomimportJobsCmd).Standalone()

	medicalImaging_listDicomimportJobsCmd.Flags().String("datastore-id", "", "The data store identifier.")
	medicalImaging_listDicomimportJobsCmd.Flags().String("job-status", "", "The filters for listing import jobs based on status.")
	medicalImaging_listDicomimportJobsCmd.Flags().String("max-results", "", "The max results count.")
	medicalImaging_listDicomimportJobsCmd.Flags().String("next-token", "", "The pagination token used to request the list of import jobs on the next page.")
	medicalImaging_listDicomimportJobsCmd.MarkFlagRequired("datastore-id")
	medicalImagingCmd.AddCommand(medicalImaging_listDicomimportJobsCmd)
}

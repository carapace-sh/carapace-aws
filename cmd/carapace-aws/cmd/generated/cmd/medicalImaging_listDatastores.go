package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medicalImaging_listDatastoresCmd = &cobra.Command{
	Use:   "list-datastores",
	Short: "List data stores.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medicalImaging_listDatastoresCmd).Standalone()

	medicalImaging_listDatastoresCmd.Flags().String("datastore-status", "", "The data store status.")
	medicalImaging_listDatastoresCmd.Flags().String("max-results", "", "Valid Range: Minimum value of 1. Maximum value of 50.")
	medicalImaging_listDatastoresCmd.Flags().String("next-token", "", "The pagination token used to request the list of data stores on the next page.")
	medicalImagingCmd.AddCommand(medicalImaging_listDatastoresCmd)
}

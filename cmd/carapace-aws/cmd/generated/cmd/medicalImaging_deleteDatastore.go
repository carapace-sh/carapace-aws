package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medicalImaging_deleteDatastoreCmd = &cobra.Command{
	Use:   "delete-datastore",
	Short: "Delete a data store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medicalImaging_deleteDatastoreCmd).Standalone()

	medicalImaging_deleteDatastoreCmd.Flags().String("datastore-id", "", "The data store identifier.")
	medicalImaging_deleteDatastoreCmd.MarkFlagRequired("datastore-id")
	medicalImagingCmd.AddCommand(medicalImaging_deleteDatastoreCmd)
}

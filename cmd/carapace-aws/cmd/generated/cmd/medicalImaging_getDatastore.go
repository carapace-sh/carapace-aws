package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medicalImaging_getDatastoreCmd = &cobra.Command{
	Use:   "get-datastore",
	Short: "Get data store properties.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medicalImaging_getDatastoreCmd).Standalone()

	medicalImaging_getDatastoreCmd.Flags().String("datastore-id", "", "The data store identifier.")
	medicalImaging_getDatastoreCmd.MarkFlagRequired("datastore-id")
	medicalImagingCmd.AddCommand(medicalImaging_getDatastoreCmd)
}

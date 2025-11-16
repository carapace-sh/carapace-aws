package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_deleteReferenceStoreCmd = &cobra.Command{
	Use:   "delete-reference-store",
	Short: "Deletes a reference store and returns a response with no body if the operation is successful.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_deleteReferenceStoreCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(omics_deleteReferenceStoreCmd).Standalone()

		omics_deleteReferenceStoreCmd.Flags().String("id", "", "The store's ID.")
		omics_deleteReferenceStoreCmd.MarkFlagRequired("id")
	})
	omicsCmd.AddCommand(omics_deleteReferenceStoreCmd)
}

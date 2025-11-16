package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_deleteReferenceCmd = &cobra.Command{
	Use:   "delete-reference",
	Short: "Deletes a reference genome and returns a response with no body if the operation is successful.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_deleteReferenceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(omics_deleteReferenceCmd).Standalone()

		omics_deleteReferenceCmd.Flags().String("id", "", "The reference's ID.")
		omics_deleteReferenceCmd.Flags().String("reference-store-id", "", "The reference's store ID.")
		omics_deleteReferenceCmd.MarkFlagRequired("id")
		omics_deleteReferenceCmd.MarkFlagRequired("reference-store-id")
	})
	omicsCmd.AddCommand(omics_deleteReferenceCmd)
}

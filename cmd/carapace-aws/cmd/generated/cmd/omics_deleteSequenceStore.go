package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_deleteSequenceStoreCmd = &cobra.Command{
	Use:   "delete-sequence-store",
	Short: "Deletes a sequence store and returns a response with no body if the operation is successful.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_deleteSequenceStoreCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(omics_deleteSequenceStoreCmd).Standalone()

		omics_deleteSequenceStoreCmd.Flags().String("id", "", "The sequence store's ID.")
		omics_deleteSequenceStoreCmd.MarkFlagRequired("id")
	})
	omicsCmd.AddCommand(omics_deleteSequenceStoreCmd)
}

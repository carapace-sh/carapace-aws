package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_getSequenceStoreCmd = &cobra.Command{
	Use:   "get-sequence-store",
	Short: "Retrieves metadata for a sequence store using its ID and returns it in JSON format.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_getSequenceStoreCmd).Standalone()

	omics_getSequenceStoreCmd.Flags().String("id", "", "The store's ID.")
	omics_getSequenceStoreCmd.MarkFlagRequired("id")
	omicsCmd.AddCommand(omics_getSequenceStoreCmd)
}

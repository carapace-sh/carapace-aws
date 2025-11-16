package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_getReferenceStoreCmd = &cobra.Command{
	Use:   "get-reference-store",
	Short: "Gets information about a reference store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_getReferenceStoreCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(omics_getReferenceStoreCmd).Standalone()

		omics_getReferenceStoreCmd.Flags().String("id", "", "The store's ID.")
		omics_getReferenceStoreCmd.MarkFlagRequired("id")
	})
	omicsCmd.AddCommand(omics_getReferenceStoreCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_getReferenceCmd = &cobra.Command{
	Use:   "get-reference",
	Short: "Downloads parts of data from a reference genome and returns the reference file in the same format that it was uploaded.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_getReferenceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(omics_getReferenceCmd).Standalone()

		omics_getReferenceCmd.Flags().String("file", "", "The file to retrieve.")
		omics_getReferenceCmd.Flags().String("id", "", "The reference's ID.")
		omics_getReferenceCmd.Flags().String("part-number", "", "The part number to retrieve.")
		omics_getReferenceCmd.Flags().String("range", "", "The range to retrieve.")
		omics_getReferenceCmd.Flags().String("reference-store-id", "", "The reference's store ID.")
		omics_getReferenceCmd.MarkFlagRequired("id")
		omics_getReferenceCmd.MarkFlagRequired("part-number")
		omics_getReferenceCmd.MarkFlagRequired("reference-store-id")
	})
	omicsCmd.AddCommand(omics_getReferenceCmd)
}

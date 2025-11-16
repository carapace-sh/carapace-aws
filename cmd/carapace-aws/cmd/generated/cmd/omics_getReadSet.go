package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_getReadSetCmd = &cobra.Command{
	Use:   "get-read-set",
	Short: "Retrieves detailed information from parts of a read set and returns the read set in the same format that it was uploaded.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_getReadSetCmd).Standalone()

	omics_getReadSetCmd.Flags().String("file", "", "The file to retrieve.")
	omics_getReadSetCmd.Flags().String("id", "", "The read set's ID.")
	omics_getReadSetCmd.Flags().String("part-number", "", "The part number to retrieve.")
	omics_getReadSetCmd.Flags().String("sequence-store-id", "", "The read set's sequence store ID.")
	omics_getReadSetCmd.MarkFlagRequired("id")
	omics_getReadSetCmd.MarkFlagRequired("part-number")
	omics_getReadSetCmd.MarkFlagRequired("sequence-store-id")
	omicsCmd.AddCommand(omics_getReadSetCmd)
}

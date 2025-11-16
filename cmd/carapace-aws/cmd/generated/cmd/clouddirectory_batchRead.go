package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clouddirectory_batchReadCmd = &cobra.Command{
	Use:   "batch-read",
	Short: "Performs all the read operations in a batch.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clouddirectory_batchReadCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(clouddirectory_batchReadCmd).Standalone()

		clouddirectory_batchReadCmd.Flags().String("consistency-level", "", "Represents the manner and timing in which the successful write or update of an object is reflected in a subsequent read operation of that same object.")
		clouddirectory_batchReadCmd.Flags().String("directory-arn", "", "The Amazon Resource Name (ARN) that is associated with the [Directory]().")
		clouddirectory_batchReadCmd.Flags().String("operations", "", "A list of operations that are part of the batch.")
		clouddirectory_batchReadCmd.MarkFlagRequired("directory-arn")
		clouddirectory_batchReadCmd.MarkFlagRequired("operations")
	})
	clouddirectoryCmd.AddCommand(clouddirectory_batchReadCmd)
}

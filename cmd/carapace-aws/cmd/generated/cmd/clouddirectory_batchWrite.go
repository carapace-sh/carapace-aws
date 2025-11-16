package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clouddirectory_batchWriteCmd = &cobra.Command{
	Use:   "batch-write",
	Short: "Performs all the write operations in a batch.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clouddirectory_batchWriteCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(clouddirectory_batchWriteCmd).Standalone()

		clouddirectory_batchWriteCmd.Flags().String("directory-arn", "", "The Amazon Resource Name (ARN) that is associated with the [Directory]().")
		clouddirectory_batchWriteCmd.Flags().String("operations", "", "A list of operations that are part of the batch.")
		clouddirectory_batchWriteCmd.MarkFlagRequired("directory-arn")
		clouddirectory_batchWriteCmd.MarkFlagRequired("operations")
	})
	clouddirectoryCmd.AddCommand(clouddirectory_batchWriteCmd)
}

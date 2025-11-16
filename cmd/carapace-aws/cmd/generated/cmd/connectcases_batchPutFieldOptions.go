package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcases_batchPutFieldOptionsCmd = &cobra.Command{
	Use:   "batch-put-field-options",
	Short: "Creates and updates a set of field options for a single select field in a Cases domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcases_batchPutFieldOptionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connectcases_batchPutFieldOptionsCmd).Standalone()

		connectcases_batchPutFieldOptionsCmd.Flags().String("domain-id", "", "The unique identifier of the Cases domain.")
		connectcases_batchPutFieldOptionsCmd.Flags().String("field-id", "", "The unique identifier of a field.")
		connectcases_batchPutFieldOptionsCmd.Flags().String("options", "", "A list of `FieldOption` objects.")
		connectcases_batchPutFieldOptionsCmd.MarkFlagRequired("domain-id")
		connectcases_batchPutFieldOptionsCmd.MarkFlagRequired("field-id")
		connectcases_batchPutFieldOptionsCmd.MarkFlagRequired("options")
	})
	connectcasesCmd.AddCommand(connectcases_batchPutFieldOptionsCmd)
}

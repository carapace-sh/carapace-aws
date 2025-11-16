package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcases_batchGetFieldCmd = &cobra.Command{
	Use:   "batch-get-field",
	Short: "Returns the description for the list of fields in the request parameters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcases_batchGetFieldCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connectcases_batchGetFieldCmd).Standalone()

		connectcases_batchGetFieldCmd.Flags().String("domain-id", "", "The unique identifier of the Cases domain.")
		connectcases_batchGetFieldCmd.Flags().String("fields", "", "A list of unique field identifiers.")
		connectcases_batchGetFieldCmd.MarkFlagRequired("domain-id")
		connectcases_batchGetFieldCmd.MarkFlagRequired("fields")
	})
	connectcasesCmd.AddCommand(connectcases_batchGetFieldCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var athena_batchGetPreparedStatementCmd = &cobra.Command{
	Use:   "batch-get-prepared-statement",
	Short: "Returns the details of a single prepared statement or a list of up to 256 prepared statements for the array of prepared statement names that you provide.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(athena_batchGetPreparedStatementCmd).Standalone()

	athena_batchGetPreparedStatementCmd.Flags().String("prepared-statement-names", "", "A list of prepared statement names to return.")
	athena_batchGetPreparedStatementCmd.Flags().String("work-group", "", "The name of the workgroup to which the prepared statements belong.")
	athena_batchGetPreparedStatementCmd.MarkFlagRequired("prepared-statement-names")
	athena_batchGetPreparedStatementCmd.MarkFlagRequired("work-group")
	athenaCmd.AddCommand(athena_batchGetPreparedStatementCmd)
}

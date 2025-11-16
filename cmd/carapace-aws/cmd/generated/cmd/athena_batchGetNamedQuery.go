package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var athena_batchGetNamedQueryCmd = &cobra.Command{
	Use:   "batch-get-named-query",
	Short: "Returns the details of a single named query or a list of up to 50 queries, which you provide as an array of query ID strings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(athena_batchGetNamedQueryCmd).Standalone()

	athena_batchGetNamedQueryCmd.Flags().String("named-query-ids", "", "An array of query IDs.")
	athena_batchGetNamedQueryCmd.MarkFlagRequired("named-query-ids")
	athenaCmd.AddCommand(athena_batchGetNamedQueryCmd)
}

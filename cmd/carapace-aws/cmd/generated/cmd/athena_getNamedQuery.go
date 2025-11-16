package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var athena_getNamedQueryCmd = &cobra.Command{
	Use:   "get-named-query",
	Short: "Returns information about a single query.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(athena_getNamedQueryCmd).Standalone()

	athena_getNamedQueryCmd.Flags().String("named-query-id", "", "The unique ID of the query.")
	athena_getNamedQueryCmd.MarkFlagRequired("named-query-id")
	athenaCmd.AddCommand(athena_getNamedQueryCmd)
}

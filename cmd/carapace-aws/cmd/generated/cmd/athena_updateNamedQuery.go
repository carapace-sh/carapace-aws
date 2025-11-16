package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var athena_updateNamedQueryCmd = &cobra.Command{
	Use:   "update-named-query",
	Short: "Updates a [NamedQuery]() object.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(athena_updateNamedQueryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(athena_updateNamedQueryCmd).Standalone()

		athena_updateNamedQueryCmd.Flags().String("description", "", "The query description.")
		athena_updateNamedQueryCmd.Flags().String("name", "", "The name of the query.")
		athena_updateNamedQueryCmd.Flags().String("named-query-id", "", "The unique identifier (UUID) of the query.")
		athena_updateNamedQueryCmd.Flags().String("query-string", "", "The contents of the query with all query statements.")
		athena_updateNamedQueryCmd.MarkFlagRequired("name")
		athena_updateNamedQueryCmd.MarkFlagRequired("named-query-id")
		athena_updateNamedQueryCmd.MarkFlagRequired("query-string")
	})
	athenaCmd.AddCommand(athena_updateNamedQueryCmd)
}

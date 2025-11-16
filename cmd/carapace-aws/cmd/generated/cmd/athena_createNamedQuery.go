package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var athena_createNamedQueryCmd = &cobra.Command{
	Use:   "create-named-query",
	Short: "Creates a named query in the specified workgroup.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(athena_createNamedQueryCmd).Standalone()

	athena_createNamedQueryCmd.Flags().String("client-request-token", "", "A unique case-sensitive string used to ensure the request to create the query is idempotent (executes only once).")
	athena_createNamedQueryCmd.Flags().String("database", "", "The database to which the query belongs.")
	athena_createNamedQueryCmd.Flags().String("description", "", "The query description.")
	athena_createNamedQueryCmd.Flags().String("name", "", "The query name.")
	athena_createNamedQueryCmd.Flags().String("query-string", "", "The contents of the query with all query statements.")
	athena_createNamedQueryCmd.Flags().String("work-group", "", "The name of the workgroup in which the named query is being created.")
	athena_createNamedQueryCmd.MarkFlagRequired("database")
	athena_createNamedQueryCmd.MarkFlagRequired("name")
	athena_createNamedQueryCmd.MarkFlagRequired("query-string")
	athenaCmd.AddCommand(athena_createNamedQueryCmd)
}

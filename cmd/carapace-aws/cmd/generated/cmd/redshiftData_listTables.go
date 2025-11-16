package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshiftData_listTablesCmd = &cobra.Command{
	Use:   "list-tables",
	Short: "List the tables in a database.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshiftData_listTablesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshiftData_listTablesCmd).Standalone()

		redshiftData_listTablesCmd.Flags().String("cluster-identifier", "", "The cluster identifier.")
		redshiftData_listTablesCmd.Flags().String("connected-database", "", "A database name.")
		redshiftData_listTablesCmd.Flags().String("database", "", "The name of the database that contains the tables to list.")
		redshiftData_listTablesCmd.Flags().String("db-user", "", "The database user name.")
		redshiftData_listTablesCmd.Flags().String("max-results", "", "The maximum number of tables to return in the response.")
		redshiftData_listTablesCmd.Flags().String("next-token", "", "A value that indicates the starting point for the next set of response records in a subsequent request.")
		redshiftData_listTablesCmd.Flags().String("schema-pattern", "", "A pattern to filter results by schema name.")
		redshiftData_listTablesCmd.Flags().String("secret-arn", "", "The name or ARN of the secret that enables access to the database.")
		redshiftData_listTablesCmd.Flags().String("table-pattern", "", "A pattern to filter results by table name.")
		redshiftData_listTablesCmd.Flags().String("workgroup-name", "", "The serverless workgroup name or Amazon Resource Name (ARN).")
		redshiftData_listTablesCmd.MarkFlagRequired("database")
	})
	redshiftDataCmd.AddCommand(redshiftData_listTablesCmd)
}

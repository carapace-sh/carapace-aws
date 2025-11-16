package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshiftData_listSchemasCmd = &cobra.Command{
	Use:   "list-schemas",
	Short: "Lists the schemas in a database.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshiftData_listSchemasCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshiftData_listSchemasCmd).Standalone()

		redshiftData_listSchemasCmd.Flags().String("cluster-identifier", "", "The cluster identifier.")
		redshiftData_listSchemasCmd.Flags().String("connected-database", "", "A database name.")
		redshiftData_listSchemasCmd.Flags().String("database", "", "The name of the database that contains the schemas to list.")
		redshiftData_listSchemasCmd.Flags().String("db-user", "", "The database user name.")
		redshiftData_listSchemasCmd.Flags().String("max-results", "", "The maximum number of schemas to return in the response.")
		redshiftData_listSchemasCmd.Flags().String("next-token", "", "A value that indicates the starting point for the next set of response records in a subsequent request.")
		redshiftData_listSchemasCmd.Flags().String("schema-pattern", "", "A pattern to filter results by schema name.")
		redshiftData_listSchemasCmd.Flags().String("secret-arn", "", "The name or ARN of the secret that enables access to the database.")
		redshiftData_listSchemasCmd.Flags().String("workgroup-name", "", "The serverless workgroup name or Amazon Resource Name (ARN).")
		redshiftData_listSchemasCmd.MarkFlagRequired("database")
	})
	redshiftDataCmd.AddCommand(redshiftData_listSchemasCmd)
}

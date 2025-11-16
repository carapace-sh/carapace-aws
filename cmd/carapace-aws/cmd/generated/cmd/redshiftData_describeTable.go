package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshiftData_describeTableCmd = &cobra.Command{
	Use:   "describe-table",
	Short: "Describes the detailed information about a table from metadata in the cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshiftData_describeTableCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshiftData_describeTableCmd).Standalone()

		redshiftData_describeTableCmd.Flags().String("cluster-identifier", "", "The cluster identifier.")
		redshiftData_describeTableCmd.Flags().String("connected-database", "", "A database name.")
		redshiftData_describeTableCmd.Flags().String("database", "", "The name of the database that contains the tables to be described.")
		redshiftData_describeTableCmd.Flags().String("db-user", "", "The database user name.")
		redshiftData_describeTableCmd.Flags().String("max-results", "", "The maximum number of tables to return in the response.")
		redshiftData_describeTableCmd.Flags().String("next-token", "", "A value that indicates the starting point for the next set of response records in a subsequent request.")
		redshiftData_describeTableCmd.Flags().String("schema", "", "The schema that contains the table.")
		redshiftData_describeTableCmd.Flags().String("secret-arn", "", "The name or ARN of the secret that enables access to the database.")
		redshiftData_describeTableCmd.Flags().String("table", "", "The table name.")
		redshiftData_describeTableCmd.Flags().String("workgroup-name", "", "The serverless workgroup name or Amazon Resource Name (ARN).")
		redshiftData_describeTableCmd.MarkFlagRequired("database")
	})
	redshiftDataCmd.AddCommand(redshiftData_describeTableCmd)
}

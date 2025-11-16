package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshiftData_listDatabasesCmd = &cobra.Command{
	Use:   "list-databases",
	Short: "List the databases in a cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshiftData_listDatabasesCmd).Standalone()

	redshiftData_listDatabasesCmd.Flags().String("cluster-identifier", "", "The cluster identifier.")
	redshiftData_listDatabasesCmd.Flags().String("database", "", "The name of the database.")
	redshiftData_listDatabasesCmd.Flags().String("db-user", "", "The database user name.")
	redshiftData_listDatabasesCmd.Flags().String("max-results", "", "The maximum number of databases to return in the response.")
	redshiftData_listDatabasesCmd.Flags().String("next-token", "", "A value that indicates the starting point for the next set of response records in a subsequent request.")
	redshiftData_listDatabasesCmd.Flags().String("secret-arn", "", "The name or ARN of the secret that enables access to the database.")
	redshiftData_listDatabasesCmd.Flags().String("workgroup-name", "", "The serverless workgroup name or Amazon Resource Name (ARN).")
	redshiftData_listDatabasesCmd.MarkFlagRequired("database")
	redshiftDataCmd.AddCommand(redshiftData_listDatabasesCmd)
}

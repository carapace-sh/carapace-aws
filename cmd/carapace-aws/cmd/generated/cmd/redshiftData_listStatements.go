package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshiftData_listStatementsCmd = &cobra.Command{
	Use:   "list-statements",
	Short: "List of SQL statements.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshiftData_listStatementsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshiftData_listStatementsCmd).Standalone()

		redshiftData_listStatementsCmd.Flags().String("cluster-identifier", "", "The cluster identifier.")
		redshiftData_listStatementsCmd.Flags().String("database", "", "The name of the database when listing statements run against a `ClusterIdentifier` or `WorkgroupName`.")
		redshiftData_listStatementsCmd.Flags().String("max-results", "", "The maximum number of SQL statements to return in the response.")
		redshiftData_listStatementsCmd.Flags().String("next-token", "", "A value that indicates the starting point for the next set of response records in a subsequent request.")
		redshiftData_listStatementsCmd.Flags().Bool("no-role-level", false, "A value that filters which statements to return in the response.")
		redshiftData_listStatementsCmd.Flags().Bool("role-level", false, "A value that filters which statements to return in the response.")
		redshiftData_listStatementsCmd.Flags().String("statement-name", "", "The name of the SQL statement specified as input to `BatchExecuteStatement` or `ExecuteStatement` to identify the query.")
		redshiftData_listStatementsCmd.Flags().String("status", "", "The status of the SQL statement to list.")
		redshiftData_listStatementsCmd.Flags().String("workgroup-name", "", "The serverless workgroup name or Amazon Resource Name (ARN).")
		redshiftData_listStatementsCmd.Flag("no-role-level").Hidden = true
	})
	redshiftDataCmd.AddCommand(redshiftData_listStatementsCmd)
}

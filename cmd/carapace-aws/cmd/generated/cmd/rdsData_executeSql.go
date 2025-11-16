package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rdsData_executeSqlCmd = &cobra.Command{
	Use:   "execute-sql",
	Short: "Runs one or more SQL statements.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rdsData_executeSqlCmd).Standalone()

	rdsData_executeSqlCmd.Flags().String("aws-secret-store-arn", "", "The Amazon Resource Name (ARN) of the secret that enables access to the DB cluster.")
	rdsData_executeSqlCmd.Flags().String("database", "", "The name of the database.")
	rdsData_executeSqlCmd.Flags().String("db-cluster-or-instance-arn", "", "The ARN of the Aurora Serverless DB cluster.")
	rdsData_executeSqlCmd.Flags().String("schema", "", "The name of the database schema.")
	rdsData_executeSqlCmd.Flags().String("sql-statements", "", "One or more SQL statements to run on the DB cluster.")
	rdsData_executeSqlCmd.MarkFlagRequired("aws-secret-store-arn")
	rdsData_executeSqlCmd.MarkFlagRequired("db-cluster-or-instance-arn")
	rdsData_executeSqlCmd.MarkFlagRequired("sql-statements")
	rdsDataCmd.AddCommand(rdsData_executeSqlCmd)
}

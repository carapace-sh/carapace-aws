package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rdsData_beginTransactionCmd = &cobra.Command{
	Use:   "begin-transaction",
	Short: "Starts a SQL transaction.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rdsData_beginTransactionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rdsData_beginTransactionCmd).Standalone()

		rdsData_beginTransactionCmd.Flags().String("database", "", "The name of the database.")
		rdsData_beginTransactionCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the Aurora Serverless DB cluster.")
		rdsData_beginTransactionCmd.Flags().String("schema", "", "The name of the database schema.")
		rdsData_beginTransactionCmd.Flags().String("secret-arn", "", "The name or ARN of the secret that enables access to the DB cluster.")
		rdsData_beginTransactionCmd.MarkFlagRequired("resource-arn")
		rdsData_beginTransactionCmd.MarkFlagRequired("secret-arn")
	})
	rdsDataCmd.AddCommand(rdsData_beginTransactionCmd)
}

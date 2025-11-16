package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dynamodb_describeContinuousBackupsCmd = &cobra.Command{
	Use:   "describe-continuous-backups",
	Short: "Checks the status of continuous backups and point in time recovery on the specified table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dynamodb_describeContinuousBackupsCmd).Standalone()

	dynamodb_describeContinuousBackupsCmd.Flags().String("table-name", "", "Name of the table for which the customer wants to check the continuous backups and point in time recovery settings.")
	dynamodb_describeContinuousBackupsCmd.MarkFlagRequired("table-name")
	dynamodbCmd.AddCommand(dynamodb_describeContinuousBackupsCmd)
}

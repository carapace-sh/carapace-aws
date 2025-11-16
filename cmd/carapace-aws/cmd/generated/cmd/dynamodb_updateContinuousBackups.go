package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dynamodb_updateContinuousBackupsCmd = &cobra.Command{
	Use:   "update-continuous-backups",
	Short: "`UpdateContinuousBackups` enables or disables point in time recovery for the specified table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dynamodb_updateContinuousBackupsCmd).Standalone()

	dynamodb_updateContinuousBackupsCmd.Flags().String("point-in-time-recovery-specification", "", "Represents the settings used to enable point in time recovery.")
	dynamodb_updateContinuousBackupsCmd.Flags().String("table-name", "", "The name of the table.")
	dynamodb_updateContinuousBackupsCmd.MarkFlagRequired("point-in-time-recovery-specification")
	dynamodb_updateContinuousBackupsCmd.MarkFlagRequired("table-name")
	dynamodbCmd.AddCommand(dynamodb_updateContinuousBackupsCmd)
}

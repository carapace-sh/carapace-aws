package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var timestreamWrite_updateDatabaseCmd = &cobra.Command{
	Use:   "update-database",
	Short: "Modifies the KMS key for an existing database.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(timestreamWrite_updateDatabaseCmd).Standalone()

	timestreamWrite_updateDatabaseCmd.Flags().String("database-name", "", "The name of the database.")
	timestreamWrite_updateDatabaseCmd.Flags().String("kms-key-id", "", "The identifier of the new KMS key (`KmsKeyId`) to be used to encrypt the data stored in the database.")
	timestreamWrite_updateDatabaseCmd.MarkFlagRequired("database-name")
	timestreamWrite_updateDatabaseCmd.MarkFlagRequired("kms-key-id")
	timestreamWriteCmd.AddCommand(timestreamWrite_updateDatabaseCmd)
}

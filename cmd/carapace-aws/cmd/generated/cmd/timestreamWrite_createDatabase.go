package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var timestreamWrite_createDatabaseCmd = &cobra.Command{
	Use:   "create-database",
	Short: "Creates a new Timestream database.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(timestreamWrite_createDatabaseCmd).Standalone()

	timestreamWrite_createDatabaseCmd.Flags().String("database-name", "", "The name of the Timestream database.")
	timestreamWrite_createDatabaseCmd.Flags().String("kms-key-id", "", "The KMS key for the database.")
	timestreamWrite_createDatabaseCmd.Flags().String("tags", "", "A list of key-value pairs to label the table.")
	timestreamWrite_createDatabaseCmd.MarkFlagRequired("database-name")
	timestreamWriteCmd.AddCommand(timestreamWrite_createDatabaseCmd)
}

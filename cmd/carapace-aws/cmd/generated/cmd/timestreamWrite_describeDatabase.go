package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var timestreamWrite_describeDatabaseCmd = &cobra.Command{
	Use:   "describe-database",
	Short: "Returns information about the database, including the database name, time that the database was created, and the total number of tables found within the database.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(timestreamWrite_describeDatabaseCmd).Standalone()

	timestreamWrite_describeDatabaseCmd.Flags().String("database-name", "", "The name of the Timestream database.")
	timestreamWrite_describeDatabaseCmd.MarkFlagRequired("database-name")
	timestreamWriteCmd.AddCommand(timestreamWrite_describeDatabaseCmd)
}

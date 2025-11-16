package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var timestreamWrite_deleteDatabaseCmd = &cobra.Command{
	Use:   "delete-database",
	Short: "Deletes a given Timestream database.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(timestreamWrite_deleteDatabaseCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(timestreamWrite_deleteDatabaseCmd).Standalone()

		timestreamWrite_deleteDatabaseCmd.Flags().String("database-name", "", "The name of the Timestream database to be deleted.")
		timestreamWrite_deleteDatabaseCmd.MarkFlagRequired("database-name")
	})
	timestreamWriteCmd.AddCommand(timestreamWrite_deleteDatabaseCmd)
}

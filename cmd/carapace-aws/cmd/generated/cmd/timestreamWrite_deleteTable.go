package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var timestreamWrite_deleteTableCmd = &cobra.Command{
	Use:   "delete-table",
	Short: "Deletes a given Timestream table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(timestreamWrite_deleteTableCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(timestreamWrite_deleteTableCmd).Standalone()

		timestreamWrite_deleteTableCmd.Flags().String("database-name", "", "The name of the database where the Timestream database is to be deleted.")
		timestreamWrite_deleteTableCmd.Flags().String("table-name", "", "The name of the Timestream table to be deleted.")
		timestreamWrite_deleteTableCmd.MarkFlagRequired("database-name")
		timestreamWrite_deleteTableCmd.MarkFlagRequired("table-name")
	})
	timestreamWriteCmd.AddCommand(timestreamWrite_deleteTableCmd)
}

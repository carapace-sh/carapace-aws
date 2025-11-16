package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var timestreamWrite_createTableCmd = &cobra.Command{
	Use:   "create-table",
	Short: "Adds a new table to an existing database in your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(timestreamWrite_createTableCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(timestreamWrite_createTableCmd).Standalone()

		timestreamWrite_createTableCmd.Flags().String("database-name", "", "The name of the Timestream database.")
		timestreamWrite_createTableCmd.Flags().String("magnetic-store-write-properties", "", "Contains properties to set on the table when enabling magnetic store writes.")
		timestreamWrite_createTableCmd.Flags().String("retention-properties", "", "The duration for which your time-series data must be stored in the memory store and the magnetic store.")
		timestreamWrite_createTableCmd.Flags().String("schema", "", "The schema of the table.")
		timestreamWrite_createTableCmd.Flags().String("table-name", "", "The name of the Timestream table.")
		timestreamWrite_createTableCmd.Flags().String("tags", "", "A list of key-value pairs to label the table.")
		timestreamWrite_createTableCmd.MarkFlagRequired("database-name")
		timestreamWrite_createTableCmd.MarkFlagRequired("table-name")
	})
	timestreamWriteCmd.AddCommand(timestreamWrite_createTableCmd)
}

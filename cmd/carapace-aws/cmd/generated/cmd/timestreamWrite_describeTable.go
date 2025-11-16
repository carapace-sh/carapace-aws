package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var timestreamWrite_describeTableCmd = &cobra.Command{
	Use:   "describe-table",
	Short: "Returns information about the table, including the table name, database name, retention duration of the memory store and the magnetic store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(timestreamWrite_describeTableCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(timestreamWrite_describeTableCmd).Standalone()

		timestreamWrite_describeTableCmd.Flags().String("database-name", "", "The name of the Timestream database.")
		timestreamWrite_describeTableCmd.Flags().String("table-name", "", "The name of the Timestream table.")
		timestreamWrite_describeTableCmd.MarkFlagRequired("database-name")
		timestreamWrite_describeTableCmd.MarkFlagRequired("table-name")
	})
	timestreamWriteCmd.AddCommand(timestreamWrite_describeTableCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var timestreamWrite_updateTableCmd = &cobra.Command{
	Use:   "update-table",
	Short: "Modifies the retention duration of the memory store and magnetic store for your Timestream table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(timestreamWrite_updateTableCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(timestreamWrite_updateTableCmd).Standalone()

		timestreamWrite_updateTableCmd.Flags().String("database-name", "", "The name of the Timestream database.")
		timestreamWrite_updateTableCmd.Flags().String("magnetic-store-write-properties", "", "Contains properties to set on the table when enabling magnetic store writes.")
		timestreamWrite_updateTableCmd.Flags().String("retention-properties", "", "The retention duration of the memory store and the magnetic store.")
		timestreamWrite_updateTableCmd.Flags().String("schema", "", "The schema of the table.")
		timestreamWrite_updateTableCmd.Flags().String("table-name", "", "The name of the Timestream table.")
		timestreamWrite_updateTableCmd.MarkFlagRequired("database-name")
		timestreamWrite_updateTableCmd.MarkFlagRequired("table-name")
	})
	timestreamWriteCmd.AddCommand(timestreamWrite_updateTableCmd)
}

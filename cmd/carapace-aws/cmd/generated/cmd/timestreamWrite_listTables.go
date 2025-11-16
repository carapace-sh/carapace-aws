package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var timestreamWrite_listTablesCmd = &cobra.Command{
	Use:   "list-tables",
	Short: "Provides a list of tables, along with the name, status, and retention properties of each table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(timestreamWrite_listTablesCmd).Standalone()

	timestreamWrite_listTablesCmd.Flags().String("database-name", "", "The name of the Timestream database.")
	timestreamWrite_listTablesCmd.Flags().String("max-results", "", "The total number of items to return in the output.")
	timestreamWrite_listTablesCmd.Flags().String("next-token", "", "The pagination token.")
	timestreamWriteCmd.AddCommand(timestreamWrite_listTablesCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bcmDataExports_getTableCmd = &cobra.Command{
	Use:   "get-table",
	Short: "Returns the metadata for the specified table and table properties.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bcmDataExports_getTableCmd).Standalone()

	bcmDataExports_getTableCmd.Flags().String("table-name", "", "The name of the table.")
	bcmDataExports_getTableCmd.Flags().String("table-properties", "", "TableProperties are additional configurations you can provide to change the data and schema of a table.")
	bcmDataExports_getTableCmd.MarkFlagRequired("table-name")
	bcmDataExportsCmd.AddCommand(bcmDataExports_getTableCmd)
}

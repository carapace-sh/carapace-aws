package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bcmDataExports_listTablesCmd = &cobra.Command{
	Use:   "list-tables",
	Short: "Lists all available tables in data exports.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bcmDataExports_listTablesCmd).Standalone()

	bcmDataExports_listTablesCmd.Flags().String("max-results", "", "The maximum number of objects that are returned for the request.")
	bcmDataExports_listTablesCmd.Flags().String("next-token", "", "The token to retrieve the next set of results.")
	bcmDataExportsCmd.AddCommand(bcmDataExports_listTablesCmd)
}

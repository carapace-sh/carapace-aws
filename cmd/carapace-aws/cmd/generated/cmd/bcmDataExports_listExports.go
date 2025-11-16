package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bcmDataExports_listExportsCmd = &cobra.Command{
	Use:   "list-exports",
	Short: "Lists all data export definitions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bcmDataExports_listExportsCmd).Standalone()

	bcmDataExports_listExportsCmd.Flags().String("max-results", "", "The maximum number of objects that are returned for the request.")
	bcmDataExports_listExportsCmd.Flags().String("next-token", "", "The token to retrieve the next set of results.")
	bcmDataExportsCmd.AddCommand(bcmDataExports_listExportsCmd)
}

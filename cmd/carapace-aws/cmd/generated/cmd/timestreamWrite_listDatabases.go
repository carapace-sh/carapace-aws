package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var timestreamWrite_listDatabasesCmd = &cobra.Command{
	Use:   "list-databases",
	Short: "Returns a list of your Timestream databases.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(timestreamWrite_listDatabasesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(timestreamWrite_listDatabasesCmd).Standalone()

		timestreamWrite_listDatabasesCmd.Flags().String("max-results", "", "The total number of items to return in the output.")
		timestreamWrite_listDatabasesCmd.Flags().String("next-token", "", "The pagination token.")
	})
	timestreamWriteCmd.AddCommand(timestreamWrite_listDatabasesCmd)
}

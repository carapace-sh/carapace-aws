package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appsync_getDataSourceCmd = &cobra.Command{
	Use:   "get-data-source",
	Short: "Retrieves a `DataSource` object.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appsync_getDataSourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appsync_getDataSourceCmd).Standalone()

		appsync_getDataSourceCmd.Flags().String("api-id", "", "The API ID.")
		appsync_getDataSourceCmd.Flags().String("name", "", "The name of the data source.")
		appsync_getDataSourceCmd.MarkFlagRequired("api-id")
		appsync_getDataSourceCmd.MarkFlagRequired("name")
	})
	appsyncCmd.AddCommand(appsync_getDataSourceCmd)
}

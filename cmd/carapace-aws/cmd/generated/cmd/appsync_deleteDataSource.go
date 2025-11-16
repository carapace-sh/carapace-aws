package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appsync_deleteDataSourceCmd = &cobra.Command{
	Use:   "delete-data-source",
	Short: "Deletes a `DataSource` object.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appsync_deleteDataSourceCmd).Standalone()

	appsync_deleteDataSourceCmd.Flags().String("api-id", "", "The API ID.")
	appsync_deleteDataSourceCmd.Flags().String("name", "", "The name of the data source.")
	appsync_deleteDataSourceCmd.MarkFlagRequired("api-id")
	appsync_deleteDataSourceCmd.MarkFlagRequired("name")
	appsyncCmd.AddCommand(appsync_deleteDataSourceCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_deleteDataSourceCmd = &cobra.Command{
	Use:   "delete-data-source",
	Short: "Deletes an Amazon Q Business data source connector.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_deleteDataSourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qbusiness_deleteDataSourceCmd).Standalone()

		qbusiness_deleteDataSourceCmd.Flags().String("application-id", "", "The identifier of the Amazon Q Business application used with the data source connector.")
		qbusiness_deleteDataSourceCmd.Flags().String("data-source-id", "", "The identifier of the data source connector that you want to delete.")
		qbusiness_deleteDataSourceCmd.Flags().String("index-id", "", "The identifier of the index used with the data source connector.")
		qbusiness_deleteDataSourceCmd.MarkFlagRequired("application-id")
		qbusiness_deleteDataSourceCmd.MarkFlagRequired("data-source-id")
		qbusiness_deleteDataSourceCmd.MarkFlagRequired("index-id")
	})
	qbusinessCmd.AddCommand(qbusiness_deleteDataSourceCmd)
}

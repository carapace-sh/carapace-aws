package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_getDataSourceCmd = &cobra.Command{
	Use:   "get-data-source",
	Short: "Gets information about an existing Amazon Q Business data source connector.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_getDataSourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qbusiness_getDataSourceCmd).Standalone()

		qbusiness_getDataSourceCmd.Flags().String("application-id", "", "The identifier of the Amazon Q Business application.")
		qbusiness_getDataSourceCmd.Flags().String("data-source-id", "", "The identifier of the data source connector.")
		qbusiness_getDataSourceCmd.Flags().String("index-id", "", "The identfier of the index used with the data source connector.")
		qbusiness_getDataSourceCmd.MarkFlagRequired("application-id")
		qbusiness_getDataSourceCmd.MarkFlagRequired("data-source-id")
		qbusiness_getDataSourceCmd.MarkFlagRequired("index-id")
	})
	qbusinessCmd.AddCommand(qbusiness_getDataSourceCmd)
}

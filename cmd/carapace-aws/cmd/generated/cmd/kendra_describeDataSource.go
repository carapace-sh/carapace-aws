package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kendra_describeDataSourceCmd = &cobra.Command{
	Use:   "describe-data-source",
	Short: "Gets information about an Amazon Kendra data source connector.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kendra_describeDataSourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kendra_describeDataSourceCmd).Standalone()

		kendra_describeDataSourceCmd.Flags().String("id", "", "The identifier of the data source connector.")
		kendra_describeDataSourceCmd.Flags().String("index-id", "", "The identifier of the index used with the data source connector.")
		kendra_describeDataSourceCmd.MarkFlagRequired("id")
		kendra_describeDataSourceCmd.MarkFlagRequired("index-id")
	})
	kendraCmd.AddCommand(kendra_describeDataSourceCmd)
}

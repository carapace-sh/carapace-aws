package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kendra_deleteDataSourceCmd = &cobra.Command{
	Use:   "delete-data-source",
	Short: "Deletes an Amazon Kendra data source connector.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kendra_deleteDataSourceCmd).Standalone()

	kendra_deleteDataSourceCmd.Flags().String("id", "", "The identifier of the data source connector you want to delete.")
	kendra_deleteDataSourceCmd.Flags().String("index-id", "", "The identifier of the index used with the data source connector.")
	kendra_deleteDataSourceCmd.MarkFlagRequired("id")
	kendra_deleteDataSourceCmd.MarkFlagRequired("index-id")
	kendraCmd.AddCommand(kendra_deleteDataSourceCmd)
}

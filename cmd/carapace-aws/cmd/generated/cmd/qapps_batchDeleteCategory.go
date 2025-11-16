package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qapps_batchDeleteCategoryCmd = &cobra.Command{
	Use:   "batch-delete-category",
	Short: "Deletes Categories for the Amazon Q Business application environment instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qapps_batchDeleteCategoryCmd).Standalone()

	qapps_batchDeleteCategoryCmd.Flags().String("categories", "", "The list of IDs of the categories to be deleted.")
	qapps_batchDeleteCategoryCmd.Flags().String("instance-id", "", "The unique identifier of the Amazon Q Business application environment instance.")
	qapps_batchDeleteCategoryCmd.MarkFlagRequired("categories")
	qapps_batchDeleteCategoryCmd.MarkFlagRequired("instance-id")
	qappsCmd.AddCommand(qapps_batchDeleteCategoryCmd)
}

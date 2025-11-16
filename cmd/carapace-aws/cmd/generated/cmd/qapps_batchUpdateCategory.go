package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qapps_batchUpdateCategoryCmd = &cobra.Command{
	Use:   "batch-update-category",
	Short: "Updates Categories for the Amazon Q Business application environment instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qapps_batchUpdateCategoryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qapps_batchUpdateCategoryCmd).Standalone()

		qapps_batchUpdateCategoryCmd.Flags().String("categories", "", "The list of categories to be updated with their new values.")
		qapps_batchUpdateCategoryCmd.Flags().String("instance-id", "", "The unique identifier of the Amazon Q Business application environment instance.")
		qapps_batchUpdateCategoryCmd.MarkFlagRequired("categories")
		qapps_batchUpdateCategoryCmd.MarkFlagRequired("instance-id")
	})
	qappsCmd.AddCommand(qapps_batchUpdateCategoryCmd)
}

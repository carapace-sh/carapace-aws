package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qapps_batchCreateCategoryCmd = &cobra.Command{
	Use:   "batch-create-category",
	Short: "Creates Categories for the Amazon Q Business application environment instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qapps_batchCreateCategoryCmd).Standalone()

	qapps_batchCreateCategoryCmd.Flags().String("categories", "", "The list of category objects to be created")
	qapps_batchCreateCategoryCmd.Flags().String("instance-id", "", "The unique identifier of the Amazon Q Business application environment instance.")
	qapps_batchCreateCategoryCmd.MarkFlagRequired("categories")
	qapps_batchCreateCategoryCmd.MarkFlagRequired("instance-id")
	qappsCmd.AddCommand(qapps_batchCreateCategoryCmd)
}

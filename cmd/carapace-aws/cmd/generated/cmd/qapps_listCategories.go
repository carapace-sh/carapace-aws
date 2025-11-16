package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qapps_listCategoriesCmd = &cobra.Command{
	Use:   "list-categories",
	Short: "Lists the categories of a Amazon Q Business application environment instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qapps_listCategoriesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qapps_listCategoriesCmd).Standalone()

		qapps_listCategoriesCmd.Flags().String("instance-id", "", "The unique identifier of the Amazon Q Business application environment instance.")
		qapps_listCategoriesCmd.MarkFlagRequired("instance-id")
	})
	qappsCmd.AddCommand(qapps_listCategoriesCmd)
}

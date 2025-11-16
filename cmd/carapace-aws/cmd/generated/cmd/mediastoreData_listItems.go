package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediastoreData_listItemsCmd = &cobra.Command{
	Use:   "list-items",
	Short: "Provides a list of metadata entries about folders and objects in the specified folder.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediastoreData_listItemsCmd).Standalone()

	mediastoreData_listItemsCmd.Flags().String("max-results", "", "The maximum number of results to return per API request.")
	mediastoreData_listItemsCmd.Flags().String("next-token", "", "The token that identifies which batch of results that you want to see.")
	mediastoreData_listItemsCmd.Flags().String("path", "", "The path in the container from which to retrieve items.")
	mediastoreDataCmd.AddCommand(mediastoreData_listItemsCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_deleteConfiguredTableCmd = &cobra.Command{
	Use:   "delete-configured-table",
	Short: "Deletes a configured table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_deleteConfiguredTableCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanrooms_deleteConfiguredTableCmd).Standalone()

		cleanrooms_deleteConfiguredTableCmd.Flags().String("configured-table-identifier", "", "The unique ID for the configured table to delete.")
		cleanrooms_deleteConfiguredTableCmd.MarkFlagRequired("configured-table-identifier")
	})
	cleanroomsCmd.AddCommand(cleanrooms_deleteConfiguredTableCmd)
}

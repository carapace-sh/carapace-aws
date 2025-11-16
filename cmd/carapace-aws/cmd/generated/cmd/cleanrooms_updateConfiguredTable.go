package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_updateConfiguredTableCmd = &cobra.Command{
	Use:   "update-configured-table",
	Short: "Updates a configured table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_updateConfiguredTableCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanrooms_updateConfiguredTableCmd).Standalone()

		cleanrooms_updateConfiguredTableCmd.Flags().String("allowed-columns", "", "The columns of the underlying table that can be used by collaborations or analysis rules.")
		cleanrooms_updateConfiguredTableCmd.Flags().String("analysis-method", "", "The analysis method for the configured table.")
		cleanrooms_updateConfiguredTableCmd.Flags().String("configured-table-identifier", "", "The identifier for the configured table to update.")
		cleanrooms_updateConfiguredTableCmd.Flags().String("description", "", "A new description for the configured table.")
		cleanrooms_updateConfiguredTableCmd.Flags().String("name", "", "A new name for the configured table.")
		cleanrooms_updateConfiguredTableCmd.Flags().String("selected-analysis-methods", "", "The selected analysis methods for the table configuration update.")
		cleanrooms_updateConfiguredTableCmd.Flags().String("table-reference", "", "")
		cleanrooms_updateConfiguredTableCmd.MarkFlagRequired("configured-table-identifier")
	})
	cleanroomsCmd.AddCommand(cleanrooms_updateConfiguredTableCmd)
}

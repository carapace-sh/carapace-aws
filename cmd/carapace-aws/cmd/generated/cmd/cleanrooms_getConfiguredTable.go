package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_getConfiguredTableCmd = &cobra.Command{
	Use:   "get-configured-table",
	Short: "Retrieves a configured table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_getConfiguredTableCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanrooms_getConfiguredTableCmd).Standalone()

		cleanrooms_getConfiguredTableCmd.Flags().String("configured-table-identifier", "", "The unique ID for the configured table to retrieve.")
		cleanrooms_getConfiguredTableCmd.MarkFlagRequired("configured-table-identifier")
	})
	cleanroomsCmd.AddCommand(cleanrooms_getConfiguredTableCmd)
}

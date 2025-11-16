package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_createConfiguredTableCmd = &cobra.Command{
	Use:   "create-configured-table",
	Short: "Creates a new configured table resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_createConfiguredTableCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanrooms_createConfiguredTableCmd).Standalone()

		cleanrooms_createConfiguredTableCmd.Flags().String("allowed-columns", "", "The columns of the underlying table that can be used by collaborations or analysis rules.")
		cleanrooms_createConfiguredTableCmd.Flags().String("analysis-method", "", "The analysis method allowed for the configured tables.")
		cleanrooms_createConfiguredTableCmd.Flags().String("description", "", "A description for the configured table.")
		cleanrooms_createConfiguredTableCmd.Flags().String("name", "", "The name of the configured table.")
		cleanrooms_createConfiguredTableCmd.Flags().String("selected-analysis-methods", "", "The analysis methods to enable for the configured table.")
		cleanrooms_createConfiguredTableCmd.Flags().String("table-reference", "", "A reference to the table being configured.")
		cleanrooms_createConfiguredTableCmd.Flags().String("tags", "", "An optional label that you can assign to a resource when you create it.")
		cleanrooms_createConfiguredTableCmd.MarkFlagRequired("allowed-columns")
		cleanrooms_createConfiguredTableCmd.MarkFlagRequired("analysis-method")
		cleanrooms_createConfiguredTableCmd.MarkFlagRequired("name")
		cleanrooms_createConfiguredTableCmd.MarkFlagRequired("table-reference")
	})
	cleanroomsCmd.AddCommand(cleanrooms_createConfiguredTableCmd)
}

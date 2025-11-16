package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector2_createFilterCmd = &cobra.Command{
	Use:   "create-filter",
	Short: "Creates a filter resource using specified filter criteria.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector2_createFilterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(inspector2_createFilterCmd).Standalone()

		inspector2_createFilterCmd.Flags().String("action", "", "Defines the action that is to be applied to the findings that match the filter.")
		inspector2_createFilterCmd.Flags().String("description", "", "A description of the filter.")
		inspector2_createFilterCmd.Flags().String("filter-criteria", "", "Defines the criteria to be used in the filter for querying findings.")
		inspector2_createFilterCmd.Flags().String("name", "", "The name of the filter.")
		inspector2_createFilterCmd.Flags().String("reason", "", "The reason for creating the filter.")
		inspector2_createFilterCmd.Flags().String("tags", "", "A list of tags for the filter.")
		inspector2_createFilterCmd.MarkFlagRequired("action")
		inspector2_createFilterCmd.MarkFlagRequired("filter-criteria")
		inspector2_createFilterCmd.MarkFlagRequired("name")
	})
	inspector2Cmd.AddCommand(inspector2_createFilterCmd)
}

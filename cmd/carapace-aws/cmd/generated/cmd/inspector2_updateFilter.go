package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector2_updateFilterCmd = &cobra.Command{
	Use:   "update-filter",
	Short: "Specifies the action that is to be applied to the findings that match the filter.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector2_updateFilterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(inspector2_updateFilterCmd).Standalone()

		inspector2_updateFilterCmd.Flags().String("action", "", "Specifies the action that is to be applied to the findings that match the filter.")
		inspector2_updateFilterCmd.Flags().String("description", "", "A description of the filter.")
		inspector2_updateFilterCmd.Flags().String("filter-arn", "", "The Amazon Resource Number (ARN) of the filter to update.")
		inspector2_updateFilterCmd.Flags().String("filter-criteria", "", "Defines the criteria to be update in the filter.")
		inspector2_updateFilterCmd.Flags().String("name", "", "The name of the filter.")
		inspector2_updateFilterCmd.Flags().String("reason", "", "The reason the filter was updated.")
		inspector2_updateFilterCmd.MarkFlagRequired("filter-arn")
	})
	inspector2Cmd.AddCommand(inspector2_updateFilterCmd)
}

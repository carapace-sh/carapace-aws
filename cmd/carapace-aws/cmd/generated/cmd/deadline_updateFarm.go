package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_updateFarmCmd = &cobra.Command{
	Use:   "update-farm",
	Short: "Updates a farm.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_updateFarmCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(deadline_updateFarmCmd).Standalone()

		deadline_updateFarmCmd.Flags().String("description", "", "The description of the farm to update.")
		deadline_updateFarmCmd.Flags().String("display-name", "", "The display name of the farm to update.")
		deadline_updateFarmCmd.Flags().String("farm-id", "", "The farm ID to update.")
		deadline_updateFarmCmd.MarkFlagRequired("farm-id")
	})
	deadlineCmd.AddCommand(deadline_updateFarmCmd)
}

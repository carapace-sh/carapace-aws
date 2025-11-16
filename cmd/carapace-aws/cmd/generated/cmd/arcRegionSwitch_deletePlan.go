package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var arcRegionSwitch_deletePlanCmd = &cobra.Command{
	Use:   "delete-plan",
	Short: "Deletes a Region switch plan.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(arcRegionSwitch_deletePlanCmd).Standalone()

	arcRegionSwitch_deletePlanCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the plan.")
	arcRegionSwitch_deletePlanCmd.MarkFlagRequired("arn")
	arcRegionSwitchCmd.AddCommand(arcRegionSwitch_deletePlanCmd)
}

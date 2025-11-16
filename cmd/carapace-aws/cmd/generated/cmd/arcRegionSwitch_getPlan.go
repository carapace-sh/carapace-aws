package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var arcRegionSwitch_getPlanCmd = &cobra.Command{
	Use:   "get-plan",
	Short: "Retrieves detailed information about a Region switch plan.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(arcRegionSwitch_getPlanCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(arcRegionSwitch_getPlanCmd).Standalone()

		arcRegionSwitch_getPlanCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the plan.")
		arcRegionSwitch_getPlanCmd.MarkFlagRequired("arn")
	})
	arcRegionSwitchCmd.AddCommand(arcRegionSwitch_getPlanCmd)
}

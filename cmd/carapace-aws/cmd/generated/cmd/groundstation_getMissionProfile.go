package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var groundstation_getMissionProfileCmd = &cobra.Command{
	Use:   "get-mission-profile",
	Short: "Returns a mission profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(groundstation_getMissionProfileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(groundstation_getMissionProfileCmd).Standalone()

		groundstation_getMissionProfileCmd.Flags().String("mission-profile-id", "", "UUID of a mission profile.")
		groundstation_getMissionProfileCmd.MarkFlagRequired("mission-profile-id")
	})
	groundstationCmd.AddCommand(groundstation_getMissionProfileCmd)
}

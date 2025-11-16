package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var groundstation_deleteMissionProfileCmd = &cobra.Command{
	Use:   "delete-mission-profile",
	Short: "Deletes a mission profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(groundstation_deleteMissionProfileCmd).Standalone()

	groundstation_deleteMissionProfileCmd.Flags().String("mission-profile-id", "", "UUID of a mission profile.")
	groundstation_deleteMissionProfileCmd.MarkFlagRequired("mission-profile-id")
	groundstationCmd.AddCommand(groundstation_deleteMissionProfileCmd)
}

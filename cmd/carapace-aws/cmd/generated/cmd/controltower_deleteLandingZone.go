package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var controltower_deleteLandingZoneCmd = &cobra.Command{
	Use:   "delete-landing-zone",
	Short: "Decommissions a landing zone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(controltower_deleteLandingZoneCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(controltower_deleteLandingZoneCmd).Standalone()

		controltower_deleteLandingZoneCmd.Flags().String("landing-zone-identifier", "", "The unique identifier of the landing zone.")
		controltower_deleteLandingZoneCmd.MarkFlagRequired("landing-zone-identifier")
	})
	controltowerCmd.AddCommand(controltower_deleteLandingZoneCmd)
}

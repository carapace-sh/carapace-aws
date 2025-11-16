package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var controltower_resetLandingZoneCmd = &cobra.Command{
	Use:   "reset-landing-zone",
	Short: "This API call resets a landing zone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(controltower_resetLandingZoneCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(controltower_resetLandingZoneCmd).Standalone()

		controltower_resetLandingZoneCmd.Flags().String("landing-zone-identifier", "", "The unique identifier of the landing zone.")
		controltower_resetLandingZoneCmd.MarkFlagRequired("landing-zone-identifier")
	})
	controltowerCmd.AddCommand(controltower_resetLandingZoneCmd)
}

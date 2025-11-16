package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var controltower_getLandingZoneCmd = &cobra.Command{
	Use:   "get-landing-zone",
	Short: "Returns details about the landing zone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(controltower_getLandingZoneCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(controltower_getLandingZoneCmd).Standalone()

		controltower_getLandingZoneCmd.Flags().String("landing-zone-identifier", "", "The unique identifier of the landing zone.")
		controltower_getLandingZoneCmd.MarkFlagRequired("landing-zone-identifier")
	})
	controltowerCmd.AddCommand(controltower_getLandingZoneCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var controltower_updateLandingZoneCmd = &cobra.Command{
	Use:   "update-landing-zone",
	Short: "This API call updates the landing zone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(controltower_updateLandingZoneCmd).Standalone()

	controltower_updateLandingZoneCmd.Flags().String("landing-zone-identifier", "", "The unique identifier of the landing zone.")
	controltower_updateLandingZoneCmd.Flags().String("manifest", "", "The manifest file (JSON) is a text file that describes your Amazon Web Services resources.")
	controltower_updateLandingZoneCmd.Flags().String("remediation-types", "", "Specifies the types of remediation actions to apply when updating the landing zone configuration.")
	controltower_updateLandingZoneCmd.Flags().String("version", "", "The landing zone version, for example, 3.2.")
	controltower_updateLandingZoneCmd.MarkFlagRequired("landing-zone-identifier")
	controltower_updateLandingZoneCmd.MarkFlagRequired("manifest")
	controltower_updateLandingZoneCmd.MarkFlagRequired("version")
	controltowerCmd.AddCommand(controltower_updateLandingZoneCmd)
}

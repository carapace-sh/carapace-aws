package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var controltower_createLandingZoneCmd = &cobra.Command{
	Use:   "create-landing-zone",
	Short: "Creates a new landing zone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(controltower_createLandingZoneCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(controltower_createLandingZoneCmd).Standalone()

		controltower_createLandingZoneCmd.Flags().String("manifest", "", "The manifest JSON file is a text file that describes your Amazon Web Services resources.")
		controltower_createLandingZoneCmd.Flags().String("remediation-types", "", "Specifies the types of remediation actions to apply when creating the landing zone, such as automatic drift correction or compliance enforcement.")
		controltower_createLandingZoneCmd.Flags().String("tags", "", "Tags to be applied to the landing zone.")
		controltower_createLandingZoneCmd.Flags().String("version", "", "The landing zone version, for example, 3.0.")
		controltower_createLandingZoneCmd.MarkFlagRequired("manifest")
		controltower_createLandingZoneCmd.MarkFlagRequired("version")
	})
	controltowerCmd.AddCommand(controltower_createLandingZoneCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_updateRadiusCmd = &cobra.Command{
	Use:   "update-radius",
	Short: "Updates the Remote Authentication Dial In User Service (RADIUS) server information for an AD Connector or Microsoft AD directory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_updateRadiusCmd).Standalone()

	ds_updateRadiusCmd.Flags().String("directory-id", "", "The identifier of the directory for which to update the RADIUS server information.")
	ds_updateRadiusCmd.Flags().String("radius-settings", "", "A [RadiusSettings]() object that contains information about the RADIUS server.")
	ds_updateRadiusCmd.MarkFlagRequired("directory-id")
	ds_updateRadiusCmd.MarkFlagRequired("radius-settings")
	dsCmd.AddCommand(ds_updateRadiusCmd)
}

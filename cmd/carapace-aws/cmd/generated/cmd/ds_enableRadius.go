package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_enableRadiusCmd = &cobra.Command{
	Use:   "enable-radius",
	Short: "Enables multi-factor authentication (MFA) with the Remote Authentication Dial In User Service (RADIUS) server for an AD Connector or Microsoft AD directory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_enableRadiusCmd).Standalone()

	ds_enableRadiusCmd.Flags().String("directory-id", "", "The identifier of the directory for which to enable MFA.")
	ds_enableRadiusCmd.Flags().String("radius-settings", "", "A [RadiusSettings]() object that contains information about the RADIUS server.")
	ds_enableRadiusCmd.MarkFlagRequired("directory-id")
	ds_enableRadiusCmd.MarkFlagRequired("radius-settings")
	dsCmd.AddCommand(ds_enableRadiusCmd)
}

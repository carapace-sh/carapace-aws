package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_disableRadiusCmd = &cobra.Command{
	Use:   "disable-radius",
	Short: "Disables multi-factor authentication (MFA) with the Remote Authentication Dial In User Service (RADIUS) server for an AD Connector or Microsoft AD directory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_disableRadiusCmd).Standalone()

	ds_disableRadiusCmd.Flags().String("directory-id", "", "The identifier of the directory for which to disable MFA.")
	ds_disableRadiusCmd.MarkFlagRequired("directory-id")
	dsCmd.AddCommand(ds_disableRadiusCmd)
}

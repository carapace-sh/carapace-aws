package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_updateSecurityControlCmd = &cobra.Command{
	Use:   "update-security-control",
	Short: "Updates the properties of a security control.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_updateSecurityControlCmd).Standalone()

	securityhub_updateSecurityControlCmd.Flags().String("last-update-reason", "", "The most recent reason for updating the properties of the security control.")
	securityhub_updateSecurityControlCmd.Flags().String("parameters", "", "An object that specifies which security control parameters to update.")
	securityhub_updateSecurityControlCmd.Flags().String("security-control-id", "", "The Amazon Resource Name (ARN) or ID of the control to update.")
	securityhub_updateSecurityControlCmd.MarkFlagRequired("parameters")
	securityhub_updateSecurityControlCmd.MarkFlagRequired("security-control-id")
	securityhubCmd.AddCommand(securityhub_updateSecurityControlCmd)
}

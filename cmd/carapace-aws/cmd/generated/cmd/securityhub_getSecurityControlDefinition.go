package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_getSecurityControlDefinitionCmd = &cobra.Command{
	Use:   "get-security-control-definition",
	Short: "Retrieves the definition of a security control.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_getSecurityControlDefinitionCmd).Standalone()

	securityhub_getSecurityControlDefinitionCmd.Flags().String("security-control-id", "", "The ID of the security control to retrieve the definition for.")
	securityhub_getSecurityControlDefinitionCmd.MarkFlagRequired("security-control-id")
	securityhubCmd.AddCommand(securityhub_getSecurityControlDefinitionCmd)
}

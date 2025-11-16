package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_updateDirectoryConfigCmd = &cobra.Command{
	Use:   "update-directory-config",
	Short: "Updates the specified Directory Config object in AppStream 2.0.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_updateDirectoryConfigCmd).Standalone()

	appstream_updateDirectoryConfigCmd.Flags().String("certificate-based-auth-properties", "", "The certificate-based authentication properties used to authenticate SAML 2.0 Identity Provider (IdP) user identities to Active Directory domain-joined streaming instances.")
	appstream_updateDirectoryConfigCmd.Flags().String("directory-name", "", "The name of the Directory Config object.")
	appstream_updateDirectoryConfigCmd.Flags().String("organizational-unit-distinguished-names", "", "The distinguished names of the organizational units for computer accounts.")
	appstream_updateDirectoryConfigCmd.Flags().String("service-account-credentials", "", "The credentials for the service account used by the fleet or image builder to connect to the directory.")
	appstream_updateDirectoryConfigCmd.MarkFlagRequired("directory-name")
	appstreamCmd.AddCommand(appstream_updateDirectoryConfigCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_createDirectoryConfigCmd = &cobra.Command{
	Use:   "create-directory-config",
	Short: "Creates a Directory Config object in AppStream 2.0.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_createDirectoryConfigCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appstream_createDirectoryConfigCmd).Standalone()

		appstream_createDirectoryConfigCmd.Flags().String("certificate-based-auth-properties", "", "The certificate-based authentication properties used to authenticate SAML 2.0 Identity Provider (IdP) user identities to Active Directory domain-joined streaming instances.")
		appstream_createDirectoryConfigCmd.Flags().String("directory-name", "", "The fully qualified name of the directory (for example, corp.example.com).")
		appstream_createDirectoryConfigCmd.Flags().String("organizational-unit-distinguished-names", "", "The distinguished names of the organizational units for computer accounts.")
		appstream_createDirectoryConfigCmd.Flags().String("service-account-credentials", "", "The credentials for the service account used by the fleet or image builder to connect to the directory.")
		appstream_createDirectoryConfigCmd.MarkFlagRequired("directory-name")
		appstream_createDirectoryConfigCmd.MarkFlagRequired("organizational-unit-distinguished-names")
	})
	appstreamCmd.AddCommand(appstream_createDirectoryConfigCmd)
}

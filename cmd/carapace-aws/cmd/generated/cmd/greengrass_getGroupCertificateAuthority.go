package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_getGroupCertificateAuthorityCmd = &cobra.Command{
	Use:   "get-group-certificate-authority",
	Short: "Retreives the CA associated with a group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_getGroupCertificateAuthorityCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(greengrass_getGroupCertificateAuthorityCmd).Standalone()

		greengrass_getGroupCertificateAuthorityCmd.Flags().String("certificate-authority-id", "", "The ID of the certificate authority.")
		greengrass_getGroupCertificateAuthorityCmd.Flags().String("group-id", "", "The ID of the Greengrass group.")
		greengrass_getGroupCertificateAuthorityCmd.MarkFlagRequired("certificate-authority-id")
		greengrass_getGroupCertificateAuthorityCmd.MarkFlagRequired("group-id")
	})
	greengrassCmd.AddCommand(greengrass_getGroupCertificateAuthorityCmd)
}

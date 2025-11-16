package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_createGroupCertificateAuthorityCmd = &cobra.Command{
	Use:   "create-group-certificate-authority",
	Short: "Creates a CA for the group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_createGroupCertificateAuthorityCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(greengrass_createGroupCertificateAuthorityCmd).Standalone()

		greengrass_createGroupCertificateAuthorityCmd.Flags().String("amzn-client-token", "", "A client token used to correlate requests and responses.")
		greengrass_createGroupCertificateAuthorityCmd.Flags().String("group-id", "", "The ID of the Greengrass group.")
		greengrass_createGroupCertificateAuthorityCmd.MarkFlagRequired("group-id")
	})
	greengrassCmd.AddCommand(greengrass_createGroupCertificateAuthorityCmd)
}

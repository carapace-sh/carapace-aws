package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_updateGroupCertificateConfigurationCmd = &cobra.Command{
	Use:   "update-group-certificate-configuration",
	Short: "Updates the Certificate expiry time for a group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_updateGroupCertificateConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(greengrass_updateGroupCertificateConfigurationCmd).Standalone()

		greengrass_updateGroupCertificateConfigurationCmd.Flags().String("certificate-expiry-in-milliseconds", "", "The amount of time remaining before the certificate expires, in milliseconds.")
		greengrass_updateGroupCertificateConfigurationCmd.Flags().String("group-id", "", "The ID of the Greengrass group.")
		greengrass_updateGroupCertificateConfigurationCmd.MarkFlagRequired("group-id")
	})
	greengrassCmd.AddCommand(greengrass_updateGroupCertificateConfigurationCmd)
}

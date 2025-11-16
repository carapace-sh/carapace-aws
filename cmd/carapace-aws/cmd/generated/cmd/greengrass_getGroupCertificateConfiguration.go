package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_getGroupCertificateConfigurationCmd = &cobra.Command{
	Use:   "get-group-certificate-configuration",
	Short: "Retrieves the current configuration for the CA used by the group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_getGroupCertificateConfigurationCmd).Standalone()

	greengrass_getGroupCertificateConfigurationCmd.Flags().String("group-id", "", "The ID of the Greengrass group.")
	greengrass_getGroupCertificateConfigurationCmd.MarkFlagRequired("group-id")
	greengrassCmd.AddCommand(greengrass_getGroupCertificateConfigurationCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_listGroupCertificateAuthoritiesCmd = &cobra.Command{
	Use:   "list-group-certificate-authorities",
	Short: "Retrieves the current CAs for a group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_listGroupCertificateAuthoritiesCmd).Standalone()

	greengrass_listGroupCertificateAuthoritiesCmd.Flags().String("group-id", "", "The ID of the Greengrass group.")
	greengrass_listGroupCertificateAuthoritiesCmd.MarkFlagRequired("group-id")
	greengrassCmd.AddCommand(greengrass_listGroupCertificateAuthoritiesCmd)
}

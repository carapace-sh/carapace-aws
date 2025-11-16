package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_listCertificatesCmd = &cobra.Command{
	Use:   "list-certificates",
	Short: "For the specified directory, lists all the certificates registered for a secure LDAP or client certificate authentication.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_listCertificatesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ds_listCertificatesCmd).Standalone()

		ds_listCertificatesCmd.Flags().String("directory-id", "", "The identifier of the directory.")
		ds_listCertificatesCmd.Flags().String("limit", "", "The number of items that should show up on one page")
		ds_listCertificatesCmd.Flags().String("next-token", "", "A token for requesting another page of certificates if the `NextToken` response element indicates that more certificates are available.")
		ds_listCertificatesCmd.MarkFlagRequired("directory-id")
	})
	dsCmd.AddCommand(ds_listCertificatesCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesWeb_listTrustStoreCertificatesCmd = &cobra.Command{
	Use:   "list-trust-store-certificates",
	Short: "Retrieves a list of trust store certificates.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesWeb_listTrustStoreCertificatesCmd).Standalone()

	workspacesWeb_listTrustStoreCertificatesCmd.Flags().String("max-results", "", "The maximum number of results to be included in the next page.")
	workspacesWeb_listTrustStoreCertificatesCmd.Flags().String("next-token", "", "The pagination token used to retrieve the next page of results for this operation.")
	workspacesWeb_listTrustStoreCertificatesCmd.Flags().String("trust-store-arn", "", "The ARN of the trust store")
	workspacesWeb_listTrustStoreCertificatesCmd.MarkFlagRequired("trust-store-arn")
	workspacesWebCmd.AddCommand(workspacesWeb_listTrustStoreCertificatesCmd)
}

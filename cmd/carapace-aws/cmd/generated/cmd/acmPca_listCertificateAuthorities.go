package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var acmPca_listCertificateAuthoritiesCmd = &cobra.Command{
	Use:   "list-certificate-authorities",
	Short: "Lists the private certificate authorities that you created by using the [CreateCertificateAuthority](https://docs.aws.amazon.com/privateca/latest/APIReference/API_CreateCertificateAuthority.html) action.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(acmPca_listCertificateAuthoritiesCmd).Standalone()

	acmPca_listCertificateAuthoritiesCmd.Flags().String("max-results", "", "Use this parameter when paginating results to specify the maximum number of items to return in the response on each page.")
	acmPca_listCertificateAuthoritiesCmd.Flags().String("next-token", "", "Use this parameter when paginating results in a subsequent request after you receive a response with truncated results.")
	acmPca_listCertificateAuthoritiesCmd.Flags().String("resource-owner", "", "Use this parameter to filter the returned set of certificate authorities based on their owner.")
	acmPcaCmd.AddCommand(acmPca_listCertificateAuthoritiesCmd)
}

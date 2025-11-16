package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var acm_listCertificatesCmd = &cobra.Command{
	Use:   "list-certificates",
	Short: "Retrieves a list of certificate ARNs and domain names.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(acm_listCertificatesCmd).Standalone()

	acm_listCertificatesCmd.Flags().String("certificate-statuses", "", "Filter the certificate list by status value.")
	acm_listCertificatesCmd.Flags().String("includes", "", "Filter the certificate list.")
	acm_listCertificatesCmd.Flags().String("max-items", "", "Use this parameter when paginating results to specify the maximum number of items to return in the response.")
	acm_listCertificatesCmd.Flags().String("next-token", "", "Use this parameter only when paginating results and only in a subsequent request after you receive a response with truncated results.")
	acm_listCertificatesCmd.Flags().String("sort-by", "", "Specifies the field to sort results by.")
	acm_listCertificatesCmd.Flags().String("sort-order", "", "Specifies the order of sorted results.")
	acmCmd.AddCommand(acm_listCertificatesCmd)
}

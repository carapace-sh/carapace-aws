package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var acmPca_describeCertificateAuthorityAuditReportCmd = &cobra.Command{
	Use:   "describe-certificate-authority-audit-report",
	Short: "Lists information about a specific audit report created by calling the [CreateCertificateAuthorityAuditReport](https://docs.aws.amazon.com/privateca/latest/APIReference/API_CreateCertificateAuthorityAuditReport.html) action.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(acmPca_describeCertificateAuthorityAuditReportCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(acmPca_describeCertificateAuthorityAuditReportCmd).Standalone()

		acmPca_describeCertificateAuthorityAuditReportCmd.Flags().String("audit-report-id", "", "The report ID returned by calling the [CreateCertificateAuthorityAuditReport](https://docs.aws.amazon.com/privateca/latest/APIReference/API_CreateCertificateAuthorityAuditReport.html) action.")
		acmPca_describeCertificateAuthorityAuditReportCmd.Flags().String("certificate-authority-arn", "", "The Amazon Resource Name (ARN) of the private CA.")
		acmPca_describeCertificateAuthorityAuditReportCmd.MarkFlagRequired("audit-report-id")
		acmPca_describeCertificateAuthorityAuditReportCmd.MarkFlagRequired("certificate-authority-arn")
	})
	acmPcaCmd.AddCommand(acmPca_describeCertificateAuthorityAuditReportCmd)
}

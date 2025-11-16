package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var acmPca_createCertificateAuthorityAuditReportCmd = &cobra.Command{
	Use:   "create-certificate-authority-audit-report",
	Short: "Creates an audit report that lists every time that your CA private key is used to issue a certificate.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(acmPca_createCertificateAuthorityAuditReportCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(acmPca_createCertificateAuthorityAuditReportCmd).Standalone()

		acmPca_createCertificateAuthorityAuditReportCmd.Flags().String("audit-report-response-format", "", "The format in which to create the report.")
		acmPca_createCertificateAuthorityAuditReportCmd.Flags().String("certificate-authority-arn", "", "The Amazon Resource Name (ARN) of the CA to be audited.")
		acmPca_createCertificateAuthorityAuditReportCmd.Flags().String("s3-bucket-name", "", "The name of the S3 bucket that will contain the audit report.")
		acmPca_createCertificateAuthorityAuditReportCmd.MarkFlagRequired("audit-report-response-format")
		acmPca_createCertificateAuthorityAuditReportCmd.MarkFlagRequired("certificate-authority-arn")
		acmPca_createCertificateAuthorityAuditReportCmd.MarkFlagRequired("s3-bucket-name")
	})
	acmPcaCmd.AddCommand(acmPca_createCertificateAuthorityAuditReportCmd)
}

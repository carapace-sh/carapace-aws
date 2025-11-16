package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_deleteCertificateCmd = &cobra.Command{
	Use:   "delete-certificate",
	Short: "Deletes the specified certificate.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_deleteCertificateCmd).Standalone()

	dms_deleteCertificateCmd.Flags().String("certificate-arn", "", "The Amazon Resource Name (ARN) of the certificate.")
	dms_deleteCertificateCmd.MarkFlagRequired("certificate-arn")
	dmsCmd.AddCommand(dms_deleteCertificateCmd)
}

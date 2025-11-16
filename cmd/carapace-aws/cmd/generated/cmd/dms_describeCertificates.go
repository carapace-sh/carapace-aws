package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_describeCertificatesCmd = &cobra.Command{
	Use:   "describe-certificates",
	Short: "Provides a description of the certificate.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_describeCertificatesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dms_describeCertificatesCmd).Standalone()

		dms_describeCertificatesCmd.Flags().String("filters", "", "Filters applied to the certificates described in the form of key-value pairs.")
		dms_describeCertificatesCmd.Flags().String("marker", "", "An optional pagination token provided by a previous request.")
		dms_describeCertificatesCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
	})
	dmsCmd.AddCommand(dms_describeCertificatesCmd)
}

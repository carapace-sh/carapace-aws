package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docdb_describeCertificatesCmd = &cobra.Command{
	Use:   "describe-certificates",
	Short: "Returns a list of certificate authority (CA) certificates provided by Amazon DocumentDB for this Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docdb_describeCertificatesCmd).Standalone()

	docdb_describeCertificatesCmd.Flags().String("certificate-identifier", "", "The user-supplied certificate identifier.")
	docdb_describeCertificatesCmd.Flags().String("filters", "", "This parameter is not currently supported.")
	docdb_describeCertificatesCmd.Flags().String("marker", "", "An optional pagination token provided by a previous `DescribeCertificates` request.")
	docdb_describeCertificatesCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
	docdbCmd.AddCommand(docdb_describeCertificatesCmd)
}

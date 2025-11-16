package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_describeCertificatesCmd = &cobra.Command{
	Use:   "describe-certificates",
	Short: "Lists the set of certificate authority (CA) certificates provided by Amazon RDS for this Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_describeCertificatesCmd).Standalone()

	rds_describeCertificatesCmd.Flags().String("certificate-identifier", "", "The user-supplied certificate identifier.")
	rds_describeCertificatesCmd.Flags().String("filters", "", "This parameter isn't currently supported.")
	rds_describeCertificatesCmd.Flags().String("marker", "", "An optional pagination token provided by a previous `DescribeCertificates` request.")
	rds_describeCertificatesCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
	rdsCmd.AddCommand(rds_describeCertificatesCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_describeHsmClientCertificatesCmd = &cobra.Command{
	Use:   "describe-hsm-client-certificates",
	Short: "Returns information about the specified HSM client certificate.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_describeHsmClientCertificatesCmd).Standalone()

	redshift_describeHsmClientCertificatesCmd.Flags().String("hsm-client-certificate-identifier", "", "The identifier of a specific HSM client certificate for which you want information.")
	redshift_describeHsmClientCertificatesCmd.Flags().String("marker", "", "An optional parameter that specifies the starting point to return a set of response records.")
	redshift_describeHsmClientCertificatesCmd.Flags().String("max-records", "", "The maximum number of response records to return in each call.")
	redshift_describeHsmClientCertificatesCmd.Flags().String("tag-keys", "", "A tag key or keys for which you want to return all matching HSM client certificates that are associated with the specified key or keys.")
	redshift_describeHsmClientCertificatesCmd.Flags().String("tag-values", "", "A tag value or values for which you want to return all matching HSM client certificates that are associated with the specified tag value or values.")
	redshiftCmd.AddCommand(redshift_describeHsmClientCertificatesCmd)
}

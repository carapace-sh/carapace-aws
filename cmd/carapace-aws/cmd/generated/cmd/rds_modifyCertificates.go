package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_modifyCertificatesCmd = &cobra.Command{
	Use:   "modify-certificates",
	Short: "Override the system-default Secure Sockets Layer/Transport Layer Security (SSL/TLS) certificate for Amazon RDS for new DB instances, or remove the override.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_modifyCertificatesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_modifyCertificatesCmd).Standalone()

		rds_modifyCertificatesCmd.Flags().String("certificate-identifier", "", "The new default certificate identifier to override the current one with.")
		rds_modifyCertificatesCmd.Flags().String("remove-customer-override", "", "Specifies whether to remove the override for the default certificate.")
	})
	rdsCmd.AddCommand(rds_modifyCertificatesCmd)
}

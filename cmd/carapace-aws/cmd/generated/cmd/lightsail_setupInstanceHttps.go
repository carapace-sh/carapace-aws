package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_setupInstanceHttpsCmd = &cobra.Command{
	Use:   "setup-instance-https",
	Short: "Creates an SSL/TLS certificate that secures traffic for your website.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_setupInstanceHttpsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lightsail_setupInstanceHttpsCmd).Standalone()

		lightsail_setupInstanceHttpsCmd.Flags().String("certificate-provider", "", "The certificate authority that issues the SSL/TLS certificate.")
		lightsail_setupInstanceHttpsCmd.Flags().String("domain-names", "", "The name of the domain and subdomains that were specified for the SSL/TLS certificate.")
		lightsail_setupInstanceHttpsCmd.Flags().String("email-address", "", "The contact method for SSL/TLS certificate renewal alerts.")
		lightsail_setupInstanceHttpsCmd.Flags().String("instance-name", "", "The name of the Lightsail instance.")
		lightsail_setupInstanceHttpsCmd.MarkFlagRequired("certificate-provider")
		lightsail_setupInstanceHttpsCmd.MarkFlagRequired("domain-names")
		lightsail_setupInstanceHttpsCmd.MarkFlagRequired("email-address")
		lightsail_setupInstanceHttpsCmd.MarkFlagRequired("instance-name")
	})
	lightsailCmd.AddCommand(lightsail_setupInstanceHttpsCmd)
}

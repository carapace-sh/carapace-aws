package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_getDomainCmd = &cobra.Command{
	Use:   "get-domain",
	Short: "Returns information about a specific domain recordset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_getDomainCmd).Standalone()

	lightsail_getDomainCmd.Flags().String("domain-name", "", "The domain name for which your want to return information about.")
	lightsail_getDomainCmd.MarkFlagRequired("domain-name")
	lightsailCmd.AddCommand(lightsail_getDomainCmd)
}

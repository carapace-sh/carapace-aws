package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_getStaticIpsCmd = &cobra.Command{
	Use:   "get-static-ips",
	Short: "Returns information about all static IPs in the user's account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_getStaticIpsCmd).Standalone()

	lightsail_getStaticIpsCmd.Flags().String("page-token", "", "The token to advance to the next page of results from your request.")
	lightsailCmd.AddCommand(lightsail_getStaticIpsCmd)
}

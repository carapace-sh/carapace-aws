package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_getDomainsCmd = &cobra.Command{
	Use:   "get-domains",
	Short: "Returns a list of all domains in the user's account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_getDomainsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lightsail_getDomainsCmd).Standalone()

		lightsail_getDomainsCmd.Flags().String("page-token", "", "The token to advance to the next page of results from your request.")
	})
	lightsailCmd.AddCommand(lightsail_getDomainsCmd)
}

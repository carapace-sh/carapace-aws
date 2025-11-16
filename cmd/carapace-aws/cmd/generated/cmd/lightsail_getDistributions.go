package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_getDistributionsCmd = &cobra.Command{
	Use:   "get-distributions",
	Short: "Returns information about one or more of your Amazon Lightsail content delivery network (CDN) distributions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_getDistributionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lightsail_getDistributionsCmd).Standalone()

		lightsail_getDistributionsCmd.Flags().String("distribution-name", "", "The name of the distribution for which to return information.")
		lightsail_getDistributionsCmd.Flags().String("page-token", "", "The token to advance to the next page of results from your request.")
	})
	lightsailCmd.AddCommand(lightsail_getDistributionsCmd)
}

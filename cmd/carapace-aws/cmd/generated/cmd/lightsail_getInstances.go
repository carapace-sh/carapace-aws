package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_getInstancesCmd = &cobra.Command{
	Use:   "get-instances",
	Short: "Returns information about all Amazon Lightsail virtual private servers, or *instances*.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_getInstancesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lightsail_getInstancesCmd).Standalone()

		lightsail_getInstancesCmd.Flags().String("page-token", "", "The token to advance to the next page of results from your request.")
	})
	lightsailCmd.AddCommand(lightsail_getInstancesCmd)
}

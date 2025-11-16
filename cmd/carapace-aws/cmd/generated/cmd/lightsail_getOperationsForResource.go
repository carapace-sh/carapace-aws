package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_getOperationsForResourceCmd = &cobra.Command{
	Use:   "get-operations-for-resource",
	Short: "Gets operations for a specific resource (an instance or a static IP).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_getOperationsForResourceCmd).Standalone()

	lightsail_getOperationsForResourceCmd.Flags().String("page-token", "", "The token to advance to the next page of results from your request.")
	lightsail_getOperationsForResourceCmd.Flags().String("resource-name", "", "The name of the resource for which you are requesting information.")
	lightsail_getOperationsForResourceCmd.MarkFlagRequired("resource-name")
	lightsailCmd.AddCommand(lightsail_getOperationsForResourceCmd)
}

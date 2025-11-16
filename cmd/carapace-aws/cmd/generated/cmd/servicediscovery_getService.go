package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicediscovery_getServiceCmd = &cobra.Command{
	Use:   "get-service",
	Short: "Gets the settings for a specified service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicediscovery_getServiceCmd).Standalone()

	servicediscovery_getServiceCmd.Flags().String("id", "", "The ID or Amazon Resource Name (ARN) of the service that you want to get settings for.")
	servicediscovery_getServiceCmd.MarkFlagRequired("id")
	servicediscoveryCmd.AddCommand(servicediscovery_getServiceCmd)
}

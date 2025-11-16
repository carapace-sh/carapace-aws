package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicediscovery_registerInstanceCmd = &cobra.Command{
	Use:   "register-instance",
	Short: "Creates or updates one or more records and, optionally, creates a health check based on the settings in a specified service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicediscovery_registerInstanceCmd).Standalone()

	servicediscovery_registerInstanceCmd.Flags().String("attributes", "", "A string map that contains the following information for the service that you specify in `ServiceId`:")
	servicediscovery_registerInstanceCmd.Flags().String("creator-request-id", "", "A unique string that identifies the request and that allows failed `RegisterInstance` requests to be retried without the risk of executing the operation twice.")
	servicediscovery_registerInstanceCmd.Flags().String("instance-id", "", "An identifier that you want to associate with the instance.")
	servicediscovery_registerInstanceCmd.Flags().String("service-id", "", "The ID or Amazon Resource Name (ARN) of the service that you want to use for settings for the instance.")
	servicediscovery_registerInstanceCmd.MarkFlagRequired("attributes")
	servicediscovery_registerInstanceCmd.MarkFlagRequired("instance-id")
	servicediscovery_registerInstanceCmd.MarkFlagRequired("service-id")
	servicediscoveryCmd.AddCommand(servicediscovery_registerInstanceCmd)
}

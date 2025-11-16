package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicediscovery_deleteServiceAttributesCmd = &cobra.Command{
	Use:   "delete-service-attributes",
	Short: "Deletes specific attributes associated with a service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicediscovery_deleteServiceAttributesCmd).Standalone()

	servicediscovery_deleteServiceAttributesCmd.Flags().String("attributes", "", "A list of keys corresponding to each attribute that you want to delete.")
	servicediscovery_deleteServiceAttributesCmd.Flags().String("service-id", "", "The ID or Amazon Resource Name (ARN) of the service from which the attributes will be deleted.")
	servicediscovery_deleteServiceAttributesCmd.MarkFlagRequired("attributes")
	servicediscovery_deleteServiceAttributesCmd.MarkFlagRequired("service-id")
	servicediscoveryCmd.AddCommand(servicediscovery_deleteServiceAttributesCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_putIntegrationCmd = &cobra.Command{
	Use:   "put-integration",
	Short: "Adds an integration between the service and a third-party service, which includes Amazon AppFlow and Amazon Connect.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_putIntegrationCmd).Standalone()

	customerProfiles_putIntegrationCmd.Flags().String("domain-name", "", "The unique name of the domain.")
	customerProfiles_putIntegrationCmd.Flags().String("event-trigger-names", "", "A list of unique names for active event triggers associated with the integration.")
	customerProfiles_putIntegrationCmd.Flags().String("flow-definition", "", "The configuration that controls how Customer Profiles retrieves data from the source.")
	customerProfiles_putIntegrationCmd.Flags().String("object-type-name", "", "The name of the profile object type.")
	customerProfiles_putIntegrationCmd.Flags().String("object-type-names", "", "A map in which each key is an event type from an external application such as Segment or Shopify, and each value is an `ObjectTypeName` (template) used to ingest the event.")
	customerProfiles_putIntegrationCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of the IAM role.")
	customerProfiles_putIntegrationCmd.Flags().String("tags", "", "The tags used to organize, track, or control access for this resource.")
	customerProfiles_putIntegrationCmd.Flags().String("uri", "", "The URI of the S3 bucket or any other type of data source.")
	customerProfiles_putIntegrationCmd.MarkFlagRequired("domain-name")
	customerProfilesCmd.AddCommand(customerProfiles_putIntegrationCmd)
}

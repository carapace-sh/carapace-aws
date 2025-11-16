package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_createDomainCmd = &cobra.Command{
	Use:   "create-domain",
	Short: "Creates a domain, which is a container for all customer data, such as customer profile attributes, object types, profile keys, and encryption keys.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_createDomainCmd).Standalone()

	customerProfiles_createDomainCmd.Flags().String("dead-letter-queue-url", "", "The URL of the SQS dead letter queue, which is used for reporting errors associated with ingesting data from third party applications.")
	customerProfiles_createDomainCmd.Flags().String("default-encryption-key", "", "The default encryption key, which is an AWS managed key, is used when no specific type of encryption key is specified.")
	customerProfiles_createDomainCmd.Flags().String("default-expiration-days", "", "The default number of days until the data within the domain expires.")
	customerProfiles_createDomainCmd.Flags().String("domain-name", "", "The unique name of the domain.")
	customerProfiles_createDomainCmd.Flags().String("matching", "", "The process of matching duplicate profiles.")
	customerProfiles_createDomainCmd.Flags().String("rule-based-matching", "", "The process of matching duplicate profiles using the Rule-Based matching.")
	customerProfiles_createDomainCmd.Flags().String("tags", "", "The tags used to organize, track, or control access for this resource.")
	customerProfiles_createDomainCmd.MarkFlagRequired("default-expiration-days")
	customerProfiles_createDomainCmd.MarkFlagRequired("domain-name")
	customerProfilesCmd.AddCommand(customerProfiles_createDomainCmd)
}

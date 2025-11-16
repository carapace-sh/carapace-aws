package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_updateDomainCmd = &cobra.Command{
	Use:   "update-domain",
	Short: "Updates the properties of a domain, including creating or selecting a dead letter queue or an encryption key.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_updateDomainCmd).Standalone()

	customerProfiles_updateDomainCmd.Flags().String("dead-letter-queue-url", "", "The URL of the SQS dead letter queue, which is used for reporting errors associated with ingesting data from third party applications.")
	customerProfiles_updateDomainCmd.Flags().String("default-encryption-key", "", "The default encryption key, which is an AWS managed key, is used when no specific type of encryption key is specified.")
	customerProfiles_updateDomainCmd.Flags().String("default-expiration-days", "", "The default number of days until the data within the domain expires.")
	customerProfiles_updateDomainCmd.Flags().String("domain-name", "", "The unique name of the domain.")
	customerProfiles_updateDomainCmd.Flags().String("matching", "", "The process of matching duplicate profiles.")
	customerProfiles_updateDomainCmd.Flags().String("rule-based-matching", "", "The process of matching duplicate profiles using the rule-Based matching.")
	customerProfiles_updateDomainCmd.Flags().String("tags", "", "The tags used to organize, track, or control access for this resource.")
	customerProfiles_updateDomainCmd.MarkFlagRequired("domain-name")
	customerProfilesCmd.AddCommand(customerProfiles_updateDomainCmd)
}

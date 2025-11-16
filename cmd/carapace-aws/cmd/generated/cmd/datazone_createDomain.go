package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_createDomainCmd = &cobra.Command{
	Use:   "create-domain",
	Short: "Creates an Amazon DataZone domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_createDomainCmd).Standalone()

	datazone_createDomainCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that is provided to ensure the idempotency of the request.")
	datazone_createDomainCmd.Flags().String("description", "", "The description of the Amazon DataZone domain.")
	datazone_createDomainCmd.Flags().String("domain-execution-role", "", "The domain execution role that is created when an Amazon DataZone domain is created.")
	datazone_createDomainCmd.Flags().String("domain-version", "", "The version of the domain that is created.")
	datazone_createDomainCmd.Flags().String("kms-key-identifier", "", "The identifier of the Amazon Web Services Key Management Service (KMS) key that is used to encrypt the Amazon DataZone domain, metadata, and reporting data.")
	datazone_createDomainCmd.Flags().String("name", "", "The name of the Amazon DataZone domain.")
	datazone_createDomainCmd.Flags().String("service-role", "", "The service role of the domain that is created.")
	datazone_createDomainCmd.Flags().String("single-sign-on", "", "The single-sign on configuration of the Amazon DataZone domain.")
	datazone_createDomainCmd.Flags().String("tags", "", "The tags specified for the Amazon DataZone domain.")
	datazone_createDomainCmd.MarkFlagRequired("domain-execution-role")
	datazone_createDomainCmd.MarkFlagRequired("name")
	datazoneCmd.AddCommand(datazone_createDomainCmd)
}

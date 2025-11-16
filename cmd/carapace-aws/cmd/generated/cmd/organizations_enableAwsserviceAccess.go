package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var organizations_enableAwsserviceAccessCmd = &cobra.Command{
	Use:   "enable-awsservice-access",
	Short: "Provides an Amazon Web Services service (the service that is specified by `ServicePrincipal`) with permissions to view the structure of an organization, create a [service-linked role](https://docs.aws.amazon.com/IAM/latest/UserGuide/using-service-linked-roles.html) in all the accounts in the organization, and allow the service to perform operations on behalf of the organization and its accounts.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(organizations_enableAwsserviceAccessCmd).Standalone()

	organizations_enableAwsserviceAccessCmd.Flags().String("service-principal", "", "The service principal name of the Amazon Web Services service for which you want to enable integration with your organization.")
	organizations_enableAwsserviceAccessCmd.MarkFlagRequired("service-principal")
	organizationsCmd.AddCommand(organizations_enableAwsserviceAccessCmd)
}

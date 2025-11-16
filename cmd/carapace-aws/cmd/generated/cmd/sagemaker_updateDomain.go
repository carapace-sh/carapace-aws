package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_updateDomainCmd = &cobra.Command{
	Use:   "update-domain",
	Short: "Updates the default settings for new user profiles in the domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_updateDomainCmd).Standalone()

	sagemaker_updateDomainCmd.Flags().String("app-network-access-type", "", "Specifies the VPC used for non-EFS traffic.")
	sagemaker_updateDomainCmd.Flags().String("app-security-group-management", "", "The entity that creates and manages the required security groups for inter-app communication in `VPCOnly` mode.")
	sagemaker_updateDomainCmd.Flags().String("default-space-settings", "", "The default settings for shared spaces that users create in the domain.")
	sagemaker_updateDomainCmd.Flags().String("default-user-settings", "", "A collection of settings.")
	sagemaker_updateDomainCmd.Flags().String("domain-id", "", "The ID of the domain to be updated.")
	sagemaker_updateDomainCmd.Flags().String("domain-settings-for-update", "", "A collection of `DomainSettings` configuration values to update.")
	sagemaker_updateDomainCmd.Flags().String("subnet-ids", "", "The VPC subnets that Studio uses for communication.")
	sagemaker_updateDomainCmd.Flags().String("tag-propagation", "", "Indicates whether custom tag propagation is supported for the domain.")
	sagemaker_updateDomainCmd.Flags().String("vpc-id", "", "The identifier for the VPC used by the domain for network communication.")
	sagemaker_updateDomainCmd.MarkFlagRequired("domain-id")
	sagemakerCmd.AddCommand(sagemaker_updateDomainCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_createDomainCmd = &cobra.Command{
	Use:   "create-domain",
	Short: "Creates a `Domain`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_createDomainCmd).Standalone()

	sagemaker_createDomainCmd.Flags().String("app-network-access-type", "", "Specifies the VPC used for non-EFS traffic.")
	sagemaker_createDomainCmd.Flags().String("app-security-group-management", "", "The entity that creates and manages the required security groups for inter-app communication in `VPCOnly` mode.")
	sagemaker_createDomainCmd.Flags().String("auth-mode", "", "The mode of authentication that members use to access the domain.")
	sagemaker_createDomainCmd.Flags().String("default-space-settings", "", "The default settings for shared spaces that users create in the domain.")
	sagemaker_createDomainCmd.Flags().String("default-user-settings", "", "The default settings to use to create a user profile when `UserSettings` isn't specified in the call to the `CreateUserProfile` API.")
	sagemaker_createDomainCmd.Flags().String("domain-name", "", "A name for the domain.")
	sagemaker_createDomainCmd.Flags().String("domain-settings", "", "A collection of `Domain` settings.")
	sagemaker_createDomainCmd.Flags().String("home-efs-file-system-kms-key-id", "", "Use `KmsKeyId`.")
	sagemaker_createDomainCmd.Flags().String("kms-key-id", "", "SageMaker AI uses Amazon Web Services KMS to encrypt EFS and EBS volumes attached to the domain with an Amazon Web Services managed key by default.")
	sagemaker_createDomainCmd.Flags().String("subnet-ids", "", "The VPC subnets that the domain uses for communication.")
	sagemaker_createDomainCmd.Flags().String("tag-propagation", "", "Indicates whether custom tag propagation is supported for the domain.")
	sagemaker_createDomainCmd.Flags().String("tags", "", "Tags to associated with the Domain.")
	sagemaker_createDomainCmd.Flags().String("vpc-id", "", "The ID of the Amazon Virtual Private Cloud (VPC) that the domain uses for communication.")
	sagemaker_createDomainCmd.MarkFlagRequired("auth-mode")
	sagemaker_createDomainCmd.MarkFlagRequired("default-user-settings")
	sagemaker_createDomainCmd.MarkFlagRequired("domain-name")
	sagemakerCmd.AddCommand(sagemaker_createDomainCmd)
}

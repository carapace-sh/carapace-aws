package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imagebuilder_createInfrastructureConfigurationCmd = &cobra.Command{
	Use:   "create-infrastructure-configuration",
	Short: "Creates a new infrastructure configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagebuilder_createInfrastructureConfigurationCmd).Standalone()

	imagebuilder_createInfrastructureConfigurationCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier you provide to ensure idempotency of the request.")
	imagebuilder_createInfrastructureConfigurationCmd.Flags().String("description", "", "The description of the infrastructure configuration.")
	imagebuilder_createInfrastructureConfigurationCmd.Flags().String("instance-metadata-options", "", "The instance metadata options that you can set for the HTTP requests that pipeline builds use to launch EC2 build and test instances.")
	imagebuilder_createInfrastructureConfigurationCmd.Flags().String("instance-profile-name", "", "The instance profile to associate with the instance used to customize your Amazon EC2 AMI.")
	imagebuilder_createInfrastructureConfigurationCmd.Flags().String("instance-types", "", "The instance types of the infrastructure configuration.")
	imagebuilder_createInfrastructureConfigurationCmd.Flags().String("key-pair", "", "The key pair of the infrastructure configuration.")
	imagebuilder_createInfrastructureConfigurationCmd.Flags().String("logging", "", "The logging configuration of the infrastructure configuration.")
	imagebuilder_createInfrastructureConfigurationCmd.Flags().String("name", "", "The name of the infrastructure configuration.")
	imagebuilder_createInfrastructureConfigurationCmd.Flags().String("placement", "", "The instance placement settings that define where the instances that are launched from your image will run.")
	imagebuilder_createInfrastructureConfigurationCmd.Flags().String("resource-tags", "", "The metadata tags to assign to the Amazon EC2 instance that Image Builder launches during the build process.")
	imagebuilder_createInfrastructureConfigurationCmd.Flags().String("security-group-ids", "", "The security group IDs to associate with the instance used to customize your Amazon EC2 AMI.")
	imagebuilder_createInfrastructureConfigurationCmd.Flags().String("sns-topic-arn", "", "The Amazon Resource Name (ARN) for the SNS topic to which we send image build event notifications.")
	imagebuilder_createInfrastructureConfigurationCmd.Flags().String("subnet-id", "", "The subnet ID in which to place the instance used to customize your Amazon EC2 AMI.")
	imagebuilder_createInfrastructureConfigurationCmd.Flags().String("tags", "", "The metadata tags to assign to the infrastructure configuration resource that Image Builder creates as output.")
	imagebuilder_createInfrastructureConfigurationCmd.Flags().String("terminate-instance-on-failure", "", "The terminate instance on failure setting of the infrastructure configuration.")
	imagebuilder_createInfrastructureConfigurationCmd.MarkFlagRequired("client-token")
	imagebuilder_createInfrastructureConfigurationCmd.MarkFlagRequired("instance-profile-name")
	imagebuilder_createInfrastructureConfigurationCmd.MarkFlagRequired("name")
	imagebuilderCmd.AddCommand(imagebuilder_createInfrastructureConfigurationCmd)
}

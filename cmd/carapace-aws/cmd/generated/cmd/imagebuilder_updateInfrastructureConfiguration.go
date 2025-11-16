package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imagebuilder_updateInfrastructureConfigurationCmd = &cobra.Command{
	Use:   "update-infrastructure-configuration",
	Short: "Updates a new infrastructure configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagebuilder_updateInfrastructureConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(imagebuilder_updateInfrastructureConfigurationCmd).Standalone()

		imagebuilder_updateInfrastructureConfigurationCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier you provide to ensure idempotency of the request.")
		imagebuilder_updateInfrastructureConfigurationCmd.Flags().String("description", "", "The description of the infrastructure configuration.")
		imagebuilder_updateInfrastructureConfigurationCmd.Flags().String("infrastructure-configuration-arn", "", "The Amazon Resource Name (ARN) of the infrastructure configuration that you want to update.")
		imagebuilder_updateInfrastructureConfigurationCmd.Flags().String("instance-metadata-options", "", "The instance metadata options that you can set for the HTTP requests that pipeline builds use to launch EC2 build and test instances.")
		imagebuilder_updateInfrastructureConfigurationCmd.Flags().String("instance-profile-name", "", "The instance profile to associate with the instance used to customize your Amazon EC2 AMI.")
		imagebuilder_updateInfrastructureConfigurationCmd.Flags().String("instance-types", "", "The instance types of the infrastructure configuration.")
		imagebuilder_updateInfrastructureConfigurationCmd.Flags().String("key-pair", "", "The key pair of the infrastructure configuration.")
		imagebuilder_updateInfrastructureConfigurationCmd.Flags().String("logging", "", "The logging configuration of the infrastructure configuration.")
		imagebuilder_updateInfrastructureConfigurationCmd.Flags().String("placement", "", "The instance placement settings that define where the instances that are launched from your image will run.")
		imagebuilder_updateInfrastructureConfigurationCmd.Flags().String("resource-tags", "", "The tags attached to the resource created by Image Builder.")
		imagebuilder_updateInfrastructureConfigurationCmd.Flags().String("security-group-ids", "", "The security group IDs to associate with the instance used to customize your Amazon EC2 AMI.")
		imagebuilder_updateInfrastructureConfigurationCmd.Flags().String("sns-topic-arn", "", "The Amazon Resource Name (ARN) for the SNS topic to which we send image build event notifications.")
		imagebuilder_updateInfrastructureConfigurationCmd.Flags().String("subnet-id", "", "The subnet ID to place the instance used to customize your Amazon EC2 AMI in.")
		imagebuilder_updateInfrastructureConfigurationCmd.Flags().String("terminate-instance-on-failure", "", "The terminate instance on failure setting of the infrastructure configuration.")
		imagebuilder_updateInfrastructureConfigurationCmd.MarkFlagRequired("client-token")
		imagebuilder_updateInfrastructureConfigurationCmd.MarkFlagRequired("infrastructure-configuration-arn")
		imagebuilder_updateInfrastructureConfigurationCmd.MarkFlagRequired("instance-profile-name")
	})
	imagebuilderCmd.AddCommand(imagebuilder_updateInfrastructureConfigurationCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_createActionConnectorCmd = &cobra.Command{
	Use:   "create-action-connector",
	Short: "Creates an action connector that enables Amazon Quick Sight to connect to external services and perform actions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_createActionConnectorCmd).Standalone()

	quicksight_createActionConnectorCmd.Flags().String("action-connector-id", "", "A unique identifier for the action connector.")
	quicksight_createActionConnectorCmd.Flags().String("authentication-config", "", "The authentication configuration for connecting to the external service.")
	quicksight_createActionConnectorCmd.Flags().String("aws-account-id", "", "The Amazon Web Services account ID associated with the action connector.")
	quicksight_createActionConnectorCmd.Flags().String("description", "", "An optional description of the action connector.")
	quicksight_createActionConnectorCmd.Flags().String("name", "", "A descriptive name for the action connector.")
	quicksight_createActionConnectorCmd.Flags().String("permissions", "", "The permissions configuration that defines which users, groups, or namespaces can access this action connector and what operations they can perform.")
	quicksight_createActionConnectorCmd.Flags().String("tags", "", "A list of tags to apply to the action connector for resource management and organization.")
	quicksight_createActionConnectorCmd.Flags().String("type", "", "The type of action connector.")
	quicksight_createActionConnectorCmd.Flags().String("vpc-connection-arn", "", "The ARN of the VPC connection to use for secure connectivity to the external service.")
	quicksight_createActionConnectorCmd.MarkFlagRequired("action-connector-id")
	quicksight_createActionConnectorCmd.MarkFlagRequired("authentication-config")
	quicksight_createActionConnectorCmd.MarkFlagRequired("aws-account-id")
	quicksight_createActionConnectorCmd.MarkFlagRequired("name")
	quicksight_createActionConnectorCmd.MarkFlagRequired("type")
	quicksightCmd.AddCommand(quicksight_createActionConnectorCmd)
}

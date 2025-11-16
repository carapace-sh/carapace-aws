package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotthingsgraph_createSystemInstanceCmd = &cobra.Command{
	Use:   "create-system-instance",
	Short: "Creates a system instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotthingsgraph_createSystemInstanceCmd).Standalone()

	iotthingsgraph_createSystemInstanceCmd.Flags().String("definition", "", "")
	iotthingsgraph_createSystemInstanceCmd.Flags().String("flow-actions-role-arn", "", "The ARN of the IAM role that AWS IoT Things Graph will assume when it executes the flow.")
	iotthingsgraph_createSystemInstanceCmd.Flags().String("greengrass-group-name", "", "The name of the Greengrass group where the system instance will be deployed.")
	iotthingsgraph_createSystemInstanceCmd.Flags().String("metrics-configuration", "", "")
	iotthingsgraph_createSystemInstanceCmd.Flags().String("s3-bucket-name", "", "The name of the Amazon Simple Storage Service bucket that will be used to store and deploy the system instance's resource file.")
	iotthingsgraph_createSystemInstanceCmd.Flags().String("tags", "", "Metadata, consisting of key-value pairs, that can be used to categorize your system instances.")
	iotthingsgraph_createSystemInstanceCmd.Flags().String("target", "", "The target type of the deployment.")
	iotthingsgraph_createSystemInstanceCmd.MarkFlagRequired("definition")
	iotthingsgraph_createSystemInstanceCmd.MarkFlagRequired("target")
	iotthingsgraphCmd.AddCommand(iotthingsgraph_createSystemInstanceCmd)
}

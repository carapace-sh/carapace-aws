package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_getFlowLogsIntegrationTemplateCmd = &cobra.Command{
	Use:   "get-flow-logs-integration-template",
	Short: "Generates a CloudFormation template that streamlines and automates the integration of VPC flow logs with Amazon Athena.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_getFlowLogsIntegrationTemplateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_getFlowLogsIntegrationTemplateCmd).Standalone()

		ec2_getFlowLogsIntegrationTemplateCmd.Flags().String("config-delivery-s3-destination-arn", "", "To store the CloudFormation template in Amazon S3, specify the location in Amazon S3.")
		ec2_getFlowLogsIntegrationTemplateCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_getFlowLogsIntegrationTemplateCmd.Flags().String("flow-log-id", "", "The ID of the flow log.")
		ec2_getFlowLogsIntegrationTemplateCmd.Flags().String("integrate-services", "", "Information about the service integration.")
		ec2_getFlowLogsIntegrationTemplateCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_getFlowLogsIntegrationTemplateCmd.MarkFlagRequired("config-delivery-s3-destination-arn")
		ec2_getFlowLogsIntegrationTemplateCmd.MarkFlagRequired("flow-log-id")
		ec2_getFlowLogsIntegrationTemplateCmd.MarkFlagRequired("integrate-services")
		ec2_getFlowLogsIntegrationTemplateCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_getFlowLogsIntegrationTemplateCmd)
}

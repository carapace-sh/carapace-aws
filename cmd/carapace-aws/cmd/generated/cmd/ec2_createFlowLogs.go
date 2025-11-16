package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createFlowLogsCmd = &cobra.Command{
	Use:   "create-flow-logs",
	Short: "Creates one or more flow logs to capture information about IP traffic for a specific network interface, subnet, or VPC.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createFlowLogsCmd).Standalone()

	ec2_createFlowLogsCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	ec2_createFlowLogsCmd.Flags().String("deliver-cross-account-role", "", "The ARN of the IAM role that allows Amazon EC2 to publish flow logs across accounts.")
	ec2_createFlowLogsCmd.Flags().String("deliver-logs-permission-arn", "", "The ARN of the IAM role that allows Amazon EC2 to publish flow logs to the log destination.")
	ec2_createFlowLogsCmd.Flags().String("destination-options", "", "The destination options.")
	ec2_createFlowLogsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createFlowLogsCmd.Flags().String("log-destination", "", "The destination for the flow log data.")
	ec2_createFlowLogsCmd.Flags().String("log-destination-type", "", "The type of destination for the flow log data.")
	ec2_createFlowLogsCmd.Flags().String("log-format", "", "The fields to include in the flow log record.")
	ec2_createFlowLogsCmd.Flags().String("log-group-name", "", "The name of a new or existing CloudWatch Logs log group where Amazon EC2 publishes your flow logs.")
	ec2_createFlowLogsCmd.Flags().String("max-aggregation-interval", "", "The maximum interval of time during which a flow of packets is captured and aggregated into a flow log record.")
	ec2_createFlowLogsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createFlowLogsCmd.Flags().String("resource-ids", "", "The IDs of the resources to monitor.")
	ec2_createFlowLogsCmd.Flags().String("resource-type", "", "The type of resource to monitor.")
	ec2_createFlowLogsCmd.Flags().String("tag-specifications", "", "The tags to apply to the flow logs.")
	ec2_createFlowLogsCmd.Flags().String("traffic-type", "", "The type of traffic to monitor (accepted traffic, rejected traffic, or all traffic).")
	ec2_createFlowLogsCmd.Flag("no-dry-run").Hidden = true
	ec2_createFlowLogsCmd.MarkFlagRequired("resource-ids")
	ec2_createFlowLogsCmd.MarkFlagRequired("resource-type")
	ec2Cmd.AddCommand(ec2_createFlowLogsCmd)
}

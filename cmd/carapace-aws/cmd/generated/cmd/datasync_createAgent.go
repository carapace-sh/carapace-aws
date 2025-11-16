package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datasync_createAgentCmd = &cobra.Command{
	Use:   "create-agent",
	Short: "Activates an DataSync agent that you deploy in your storage environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datasync_createAgentCmd).Standalone()

	datasync_createAgentCmd.Flags().String("activation-key", "", "Specifies your DataSync agent's activation key.")
	datasync_createAgentCmd.Flags().String("agent-name", "", "Specifies a name for your agent.")
	datasync_createAgentCmd.Flags().String("security-group-arns", "", "Specifies the Amazon Resource Name (ARN) of the security group that allows traffic between your agent and VPC service endpoint.")
	datasync_createAgentCmd.Flags().String("subnet-arns", "", "Specifies the ARN of the subnet where your VPC service endpoint is located.")
	datasync_createAgentCmd.Flags().String("tags", "", "Specifies labels that help you categorize, filter, and search for your Amazon Web Services resources.")
	datasync_createAgentCmd.Flags().String("vpc-endpoint-id", "", "Specifies the ID of the [VPC service endpoint](https://docs.aws.amazon.com/datasync/latest/userguide/choose-service-endpoint.html#datasync-in-vpc) that you're using.")
	datasync_createAgentCmd.MarkFlagRequired("activation-key")
	datasyncCmd.AddCommand(datasync_createAgentCmd)
}

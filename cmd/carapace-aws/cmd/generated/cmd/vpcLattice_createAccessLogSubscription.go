package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpcLattice_createAccessLogSubscriptionCmd = &cobra.Command{
	Use:   "create-access-log-subscription",
	Short: "Enables access logs to be sent to Amazon CloudWatch, Amazon S3, and Amazon Kinesis Data Firehose.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcLattice_createAccessLogSubscriptionCmd).Standalone()

	vpcLattice_createAccessLogSubscriptionCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	vpcLattice_createAccessLogSubscriptionCmd.Flags().String("destination-arn", "", "The Amazon Resource Name (ARN) of the destination.")
	vpcLattice_createAccessLogSubscriptionCmd.Flags().String("resource-identifier", "", "The ID or ARN of the service network or service.")
	vpcLattice_createAccessLogSubscriptionCmd.Flags().String("service-network-log-type", "", "The type of log that monitors your Amazon VPC Lattice service networks.")
	vpcLattice_createAccessLogSubscriptionCmd.Flags().String("tags", "", "The tags for the access log subscription.")
	vpcLattice_createAccessLogSubscriptionCmd.MarkFlagRequired("destination-arn")
	vpcLattice_createAccessLogSubscriptionCmd.MarkFlagRequired("resource-identifier")
	vpcLatticeCmd.AddCommand(vpcLattice_createAccessLogSubscriptionCmd)
}

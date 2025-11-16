package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_describeEc2InstanceLimitsCmd = &cobra.Command{
	Use:   "describe-ec2-instance-limits",
	Short: "**This API works with the following fleet types:** EC2\n\nRetrieves the instance limits and current utilization for an Amazon Web Services Region or location.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_describeEc2InstanceLimitsCmd).Standalone()

	gamelift_describeEc2InstanceLimitsCmd.Flags().String("ec2-instance-type", "", "Name of an Amazon EC2 instance type that is supported in Amazon GameLift Servers.")
	gamelift_describeEc2InstanceLimitsCmd.Flags().String("location", "", "The name of a remote location to request instance limits for, in the form of an Amazon Web Services Region code such as `us-west-2`.")
	gameliftCmd.AddCommand(gamelift_describeEc2InstanceLimitsCmd)
}

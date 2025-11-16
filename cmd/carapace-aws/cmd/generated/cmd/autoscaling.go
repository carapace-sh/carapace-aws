package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoscalingCmd = &cobra.Command{
	Use:   "autoscaling",
	Short: "Amazon EC2 Auto Scaling\n\nThe [DescribeAutoScalingGroups](https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_DescribeAutoScalingGroups.html) API operation might be throttled when retrieving details for an Auto Scaling group that contains many instances.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoscalingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(autoscalingCmd).Standalone()

	})
	rootCmd.AddCommand(autoscalingCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var computeOptimizerCmd = &cobra.Command{
	Use:   "compute-optimizer",
	Short: "Compute Optimizer is a service that analyzes the configuration and utilization metrics of your Amazon Web Services compute resources, such as Amazon EC2 instances, Amazon EC2 Auto Scaling groups, Lambda functions, Amazon EBS volumes, and Amazon ECS services on Fargate.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(computeOptimizerCmd).Standalone()

	rootCmd.AddCommand(computeOptimizerCmd)
}

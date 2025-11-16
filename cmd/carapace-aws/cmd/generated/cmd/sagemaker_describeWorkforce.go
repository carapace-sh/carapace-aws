package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_describeWorkforceCmd = &cobra.Command{
	Use:   "describe-workforce",
	Short: "Lists private workforce information, including workforce name, Amazon Resource Name (ARN), and, if applicable, allowed IP address ranges ([CIDRs](https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Subnets.html)). Allowable IP address ranges are the IP addresses that workers can use to access tasks.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_describeWorkforceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_describeWorkforceCmd).Standalone()

		sagemaker_describeWorkforceCmd.Flags().String("workforce-name", "", "The name of the private workforce whose access you want to restrict.")
		sagemaker_describeWorkforceCmd.MarkFlagRequired("workforce-name")
	})
	sagemakerCmd.AddCommand(sagemaker_describeWorkforceCmd)
}

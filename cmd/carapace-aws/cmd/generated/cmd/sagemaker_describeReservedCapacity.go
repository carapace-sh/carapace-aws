package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_describeReservedCapacityCmd = &cobra.Command{
	Use:   "describe-reserved-capacity",
	Short: "Retrieves details about a reserved capacity.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_describeReservedCapacityCmd).Standalone()

	sagemaker_describeReservedCapacityCmd.Flags().String("reserved-capacity-arn", "", "ARN of the reserved capacity to describe.")
	sagemaker_describeReservedCapacityCmd.MarkFlagRequired("reserved-capacity-arn")
	sagemakerCmd.AddCommand(sagemaker_describeReservedCapacityCmd)
}

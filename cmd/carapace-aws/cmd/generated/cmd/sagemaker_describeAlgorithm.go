package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_describeAlgorithmCmd = &cobra.Command{
	Use:   "describe-algorithm",
	Short: "Returns a description of the specified algorithm that is in your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_describeAlgorithmCmd).Standalone()

	sagemaker_describeAlgorithmCmd.Flags().String("algorithm-name", "", "The name of the algorithm to describe.")
	sagemaker_describeAlgorithmCmd.MarkFlagRequired("algorithm-name")
	sagemakerCmd.AddCommand(sagemaker_describeAlgorithmCmd)
}

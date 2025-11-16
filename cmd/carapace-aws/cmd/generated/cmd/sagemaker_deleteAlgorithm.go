package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_deleteAlgorithmCmd = &cobra.Command{
	Use:   "delete-algorithm",
	Short: "Removes the specified algorithm from your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_deleteAlgorithmCmd).Standalone()

	sagemaker_deleteAlgorithmCmd.Flags().String("algorithm-name", "", "The name of the algorithm to delete.")
	sagemaker_deleteAlgorithmCmd.MarkFlagRequired("algorithm-name")
	sagemakerCmd.AddCommand(sagemaker_deleteAlgorithmCmd)
}

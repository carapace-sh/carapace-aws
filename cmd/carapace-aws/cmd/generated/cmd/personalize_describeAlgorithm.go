package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalize_describeAlgorithmCmd = &cobra.Command{
	Use:   "describe-algorithm",
	Short: "Describes the given algorithm.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalize_describeAlgorithmCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(personalize_describeAlgorithmCmd).Standalone()

		personalize_describeAlgorithmCmd.Flags().String("algorithm-arn", "", "The Amazon Resource Name (ARN) of the algorithm to describe.")
		personalize_describeAlgorithmCmd.MarkFlagRequired("algorithm-arn")
	})
	personalizeCmd.AddCommand(personalize_describeAlgorithmCmd)
}

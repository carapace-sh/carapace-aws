package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var evidently_testSegmentPatternCmd = &cobra.Command{
	Use:   "test-segment-pattern",
	Short: "Use this operation to test a rules pattern that you plan to use to create an audience segment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(evidently_testSegmentPatternCmd).Standalone()

	evidently_testSegmentPatternCmd.Flags().String("pattern", "", "The pattern to test.")
	evidently_testSegmentPatternCmd.Flags().String("payload", "", "A sample `evaluationContext` JSON block to test against the specified pattern.")
	evidently_testSegmentPatternCmd.MarkFlagRequired("pattern")
	evidently_testSegmentPatternCmd.MarkFlagRequired("payload")
	evidentlyCmd.AddCommand(evidently_testSegmentPatternCmd)
}

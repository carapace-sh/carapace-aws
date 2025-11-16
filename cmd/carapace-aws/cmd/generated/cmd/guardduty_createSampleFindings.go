package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var guardduty_createSampleFindingsCmd = &cobra.Command{
	Use:   "create-sample-findings",
	Short: "Generates sample findings of types specified by the list of finding types.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(guardduty_createSampleFindingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(guardduty_createSampleFindingsCmd).Standalone()

		guardduty_createSampleFindingsCmd.Flags().String("detector-id", "", "The ID of the detector for which you need to create sample findings.")
		guardduty_createSampleFindingsCmd.Flags().String("finding-types", "", "The types of sample findings to generate.")
		guardduty_createSampleFindingsCmd.MarkFlagRequired("detector-id")
	})
	guarddutyCmd.AddCommand(guardduty_createSampleFindingsCmd)
}

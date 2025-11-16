package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_describeOfferingCmd = &cobra.Command{
	Use:   "describe-offering",
	Short: "Get details for an offering.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_describeOfferingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(medialive_describeOfferingCmd).Standalone()

		medialive_describeOfferingCmd.Flags().String("offering-id", "", "Unique offering ID, e.g. '87654321'")
		medialive_describeOfferingCmd.MarkFlagRequired("offering-id")
	})
	medialiveCmd.AddCommand(medialive_describeOfferingCmd)
}

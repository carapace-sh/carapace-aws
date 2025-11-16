package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var directconnect_describeInterconnectsCmd = &cobra.Command{
	Use:   "describe-interconnects",
	Short: "Lists the interconnects owned by the Amazon Web Services account or only the specified interconnect.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(directconnect_describeInterconnectsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(directconnect_describeInterconnectsCmd).Standalone()

		directconnect_describeInterconnectsCmd.Flags().String("interconnect-id", "", "The ID of the interconnect.")
		directconnect_describeInterconnectsCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
		directconnect_describeInterconnectsCmd.Flags().String("next-token", "", "The token for the next page of results.")
	})
	directconnectCmd.AddCommand(directconnect_describeInterconnectsCmd)
}

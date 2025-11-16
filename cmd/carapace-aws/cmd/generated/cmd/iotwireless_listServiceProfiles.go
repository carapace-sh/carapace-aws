package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_listServiceProfilesCmd = &cobra.Command{
	Use:   "list-service-profiles",
	Short: "Lists the service profiles registered to your AWS account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_listServiceProfilesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotwireless_listServiceProfilesCmd).Standalone()

		iotwireless_listServiceProfilesCmd.Flags().String("max-results", "", "The maximum number of results to return in this operation.")
		iotwireless_listServiceProfilesCmd.Flags().String("next-token", "", "To retrieve the next set of results, the `nextToken` value from a previous response; otherwise **null** to receive the first set of results.")
	})
	iotwirelessCmd.AddCommand(iotwireless_listServiceProfilesCmd)
}

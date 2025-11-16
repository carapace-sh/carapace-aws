package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsecuretunneling_listTunnelsCmd = &cobra.Command{
	Use:   "list-tunnels",
	Short: "List all tunnels for an Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsecuretunneling_listTunnelsCmd).Standalone()

	iotsecuretunneling_listTunnelsCmd.Flags().String("max-results", "", "The maximum number of results to return at once.")
	iotsecuretunneling_listTunnelsCmd.Flags().String("next-token", "", "To retrieve the next set of results, the nextToken value from a previous response; otherwise null to receive the first set of results.")
	iotsecuretunneling_listTunnelsCmd.Flags().String("thing-name", "", "The name of the IoT thing associated with the destination device.")
	iotsecuretunnelingCmd.AddCommand(iotsecuretunneling_listTunnelsCmd)
}

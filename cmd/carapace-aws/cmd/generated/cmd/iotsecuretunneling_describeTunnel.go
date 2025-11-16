package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsecuretunneling_describeTunnelCmd = &cobra.Command{
	Use:   "describe-tunnel",
	Short: "Gets information about a tunnel identified by the unique tunnel id.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsecuretunneling_describeTunnelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotsecuretunneling_describeTunnelCmd).Standalone()

		iotsecuretunneling_describeTunnelCmd.Flags().String("tunnel-id", "", "The tunnel to describe.")
		iotsecuretunneling_describeTunnelCmd.MarkFlagRequired("tunnel-id")
	})
	iotsecuretunnelingCmd.AddCommand(iotsecuretunneling_describeTunnelCmd)
}

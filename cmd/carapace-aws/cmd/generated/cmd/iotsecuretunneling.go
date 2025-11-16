package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsecuretunnelingCmd = &cobra.Command{
	Use:   "iotsecuretunneling",
	Short: "IoT Secure Tunneling",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsecuretunnelingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotsecuretunnelingCmd).Standalone()

	})
	rootCmd.AddCommand(iotsecuretunnelingCmd)
}

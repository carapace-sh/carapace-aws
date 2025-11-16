package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsecuretunnelingCmd = &cobra.Command{
	Use:   "iotsecuretunneling",
	Short: "IoT Secure Tunneling\n\nIoT Secure Tunneling creates remote connections to devices deployed in the field.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsecuretunnelingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotsecuretunnelingCmd).Standalone()

	})
	rootCmd.AddCommand(iotsecuretunnelingCmd)
}

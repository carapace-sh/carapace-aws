package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var directconnectCmd = &cobra.Command{
	Use:   "directconnect",
	Short: "Direct Connect links your internal network to an Direct Connect location over a standard Ethernet fiber-optic cable.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(directconnectCmd).Standalone()

	rootCmd.AddCommand(directconnectCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconnectCmd = &cobra.Command{
	Use:   "mediaconnect",
	Short: "Welcome to the Elemental MediaConnect API reference.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconnectCmd).Standalone()

	rootCmd.AddCommand(mediaconnectCmd)
}

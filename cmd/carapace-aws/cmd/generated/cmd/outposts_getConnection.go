package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var outposts_getConnectionCmd = &cobra.Command{
	Use:   "get-connection",
	Short: "Amazon Web Services uses this action to install Outpost servers.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(outposts_getConnectionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(outposts_getConnectionCmd).Standalone()

		outposts_getConnectionCmd.Flags().String("connection-id", "", "The ID of the connection.")
		outposts_getConnectionCmd.MarkFlagRequired("connection-id")
	})
	outpostsCmd.AddCommand(outposts_getConnectionCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_listChannelsCmd = &cobra.Command{
	Use:   "list-channels",
	Short: "Produces list of channels that have been created",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_listChannelsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(medialive_listChannelsCmd).Standalone()

		medialive_listChannelsCmd.Flags().String("max-results", "", "")
		medialive_listChannelsCmd.Flags().String("next-token", "", "")
	})
	medialiveCmd.AddCommand(medialive_listChannelsCmd)
}

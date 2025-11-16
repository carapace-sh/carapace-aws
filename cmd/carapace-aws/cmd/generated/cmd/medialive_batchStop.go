package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_batchStopCmd = &cobra.Command{
	Use:   "batch-stop",
	Short: "Stops running resources",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_batchStopCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(medialive_batchStopCmd).Standalone()

		medialive_batchStopCmd.Flags().String("channel-ids", "", "List of channel IDs")
		medialive_batchStopCmd.Flags().String("multiplex-ids", "", "List of multiplex IDs")
	})
	medialiveCmd.AddCommand(medialive_batchStopCmd)
}

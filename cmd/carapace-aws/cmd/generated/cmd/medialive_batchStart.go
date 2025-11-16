package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_batchStartCmd = &cobra.Command{
	Use:   "batch-start",
	Short: "Starts existing resources",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_batchStartCmd).Standalone()

	medialive_batchStartCmd.Flags().String("channel-ids", "", "List of channel IDs")
	medialive_batchStartCmd.Flags().String("multiplex-ids", "", "List of multiplex IDs")
	medialiveCmd.AddCommand(medialive_batchStartCmd)
}

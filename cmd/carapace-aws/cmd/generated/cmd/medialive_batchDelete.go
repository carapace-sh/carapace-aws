package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_batchDeleteCmd = &cobra.Command{
	Use:   "batch-delete",
	Short: "Starts delete of resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_batchDeleteCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(medialive_batchDeleteCmd).Standalone()

		medialive_batchDeleteCmd.Flags().String("channel-ids", "", "List of channel IDs")
		medialive_batchDeleteCmd.Flags().String("input-ids", "", "List of input IDs")
		medialive_batchDeleteCmd.Flags().String("input-security-group-ids", "", "List of input security group IDs")
		medialive_batchDeleteCmd.Flags().String("multiplex-ids", "", "List of multiplex IDs")
	})
	medialiveCmd.AddCommand(medialive_batchDeleteCmd)
}

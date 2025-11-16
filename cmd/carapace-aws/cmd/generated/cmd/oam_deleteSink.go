package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var oam_deleteSinkCmd = &cobra.Command{
	Use:   "delete-sink",
	Short: "Deletes a sink.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(oam_deleteSinkCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(oam_deleteSinkCmd).Standalone()

		oam_deleteSinkCmd.Flags().String("identifier", "", "The ARN of the sink to delete.")
		oam_deleteSinkCmd.MarkFlagRequired("identifier")
	})
	oamCmd.AddCommand(oam_deleteSinkCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rbin_getRuleCmd = &cobra.Command{
	Use:   "get-rule",
	Short: "Gets information about a Recycle Bin retention rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rbin_getRuleCmd).Standalone()

	rbin_getRuleCmd.Flags().String("identifier", "", "The unique ID of the retention rule.")
	rbin_getRuleCmd.MarkFlagRequired("identifier")
	rbinCmd.AddCommand(rbin_getRuleCmd)
}

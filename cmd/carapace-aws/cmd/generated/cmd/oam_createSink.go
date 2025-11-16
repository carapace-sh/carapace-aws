package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var oam_createSinkCmd = &cobra.Command{
	Use:   "create-sink",
	Short: "Use this to create a *sink* in the current account, so that it can be used as a monitoring account in CloudWatch cross-account observability.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(oam_createSinkCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(oam_createSinkCmd).Standalone()

		oam_createSinkCmd.Flags().String("name", "", "A name for the sink.")
		oam_createSinkCmd.Flags().String("tags", "", "Assigns one or more tags (key-value pairs) to the link.")
		oam_createSinkCmd.MarkFlagRequired("name")
	})
	oamCmd.AddCommand(oam_createSinkCmd)
}

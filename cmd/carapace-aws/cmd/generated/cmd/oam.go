package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var oamCmd = &cobra.Command{
	Use:   "oam",
	Short: "Use Amazon CloudWatch Observability Access Manager to create and manage links between source accounts and monitoring accounts by using *CloudWatch cross-account observability*.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(oamCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(oamCmd).Standalone()

	})
	rootCmd.AddCommand(oamCmd)
}

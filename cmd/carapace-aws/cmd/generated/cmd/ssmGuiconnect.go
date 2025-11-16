package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmGuiconnectCmd = &cobra.Command{
	Use:   "ssm-guiconnect",
	Short: "AWS Systems Manager GUI Connect",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmGuiconnectCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssmGuiconnectCmd).Standalone()

	})
	rootCmd.AddCommand(ssmGuiconnectCmd)
}

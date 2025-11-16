package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmGuiconnectCmd = &cobra.Command{
	Use:   "ssm-guiconnect",
	Short: "AWS Systems Manager GUI Connect\n\nSystems Manager GUI Connect, a component of Fleet Manager, lets you connect to your Window Server-type Amazon Elastic Compute Cloud (Amazon EC2) instances using the Remote Desktop Protocol (RDP).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmGuiconnectCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssmGuiconnectCmd).Standalone()

	})
	rootCmd.AddCommand(ssmGuiconnectCmd)
}

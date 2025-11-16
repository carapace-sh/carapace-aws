package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codestarConnections_updateHostCmd = &cobra.Command{
	Use:   "update-host",
	Short: "Updates a specified host with the provided configurations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codestarConnections_updateHostCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codestarConnections_updateHostCmd).Standalone()

		codestarConnections_updateHostCmd.Flags().String("host-arn", "", "The Amazon Resource Name (ARN) of the host to be updated.")
		codestarConnections_updateHostCmd.Flags().String("provider-endpoint", "", "The URL or endpoint of the host to be updated.")
		codestarConnections_updateHostCmd.Flags().String("vpc-configuration", "", "The VPC configuration of the host to be updated.")
		codestarConnections_updateHostCmd.MarkFlagRequired("host-arn")
	})
	codestarConnectionsCmd.AddCommand(codestarConnections_updateHostCmd)
}

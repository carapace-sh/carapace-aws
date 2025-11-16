package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_getConnectionStatusCmd = &cobra.Command{
	Use:   "get-connection-status",
	Short: "Retrieves the Session Manager connection status for a managed node to determine whether it is running and ready to receive Session Manager connections.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_getConnectionStatusCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssm_getConnectionStatusCmd).Standalone()

		ssm_getConnectionStatusCmd.Flags().String("target", "", "The managed node ID.")
		ssm_getConnectionStatusCmd.MarkFlagRequired("target")
	})
	ssmCmd.AddCommand(ssm_getConnectionStatusCmd)
}

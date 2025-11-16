package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codestarConnections_deleteHostCmd = &cobra.Command{
	Use:   "delete-host",
	Short: "The host to be deleted.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codestarConnections_deleteHostCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codestarConnections_deleteHostCmd).Standalone()

		codestarConnections_deleteHostCmd.Flags().String("host-arn", "", "The Amazon Resource Name (ARN) of the host to be deleted.")
		codestarConnections_deleteHostCmd.MarkFlagRequired("host-arn")
	})
	codestarConnectionsCmd.AddCommand(codestarConnections_deleteHostCmd)
}

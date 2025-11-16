package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var drs_disconnectSourceServerCmd = &cobra.Command{
	Use:   "disconnect-source-server",
	Short: "Disconnects a specific Source Server from Elastic Disaster Recovery.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(drs_disconnectSourceServerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(drs_disconnectSourceServerCmd).Standalone()

		drs_disconnectSourceServerCmd.Flags().String("source-server-id", "", "The ID of the Source Server to disconnect.")
		drs_disconnectSourceServerCmd.MarkFlagRequired("source-server-id")
	})
	drsCmd.AddCommand(drs_disconnectSourceServerCmd)
}

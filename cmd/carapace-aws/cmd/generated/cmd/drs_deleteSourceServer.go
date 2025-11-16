package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var drs_deleteSourceServerCmd = &cobra.Command{
	Use:   "delete-source-server",
	Short: "Deletes a single Source Server by ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(drs_deleteSourceServerCmd).Standalone()

	drs_deleteSourceServerCmd.Flags().String("source-server-id", "", "The ID of the Source Server to be deleted.")
	drs_deleteSourceServerCmd.MarkFlagRequired("source-server-id")
	drsCmd.AddCommand(drs_deleteSourceServerCmd)
}

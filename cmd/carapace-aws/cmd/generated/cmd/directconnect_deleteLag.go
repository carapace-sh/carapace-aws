package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var directconnect_deleteLagCmd = &cobra.Command{
	Use:   "delete-lag",
	Short: "Deletes the specified link aggregation group (LAG).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(directconnect_deleteLagCmd).Standalone()

	directconnect_deleteLagCmd.Flags().String("lag-id", "", "The ID of the LAG.")
	directconnect_deleteLagCmd.MarkFlagRequired("lag-id")
	directconnectCmd.AddCommand(directconnect_deleteLagCmd)
}

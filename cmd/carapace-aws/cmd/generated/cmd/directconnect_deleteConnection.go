package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var directconnect_deleteConnectionCmd = &cobra.Command{
	Use:   "delete-connection",
	Short: "Deletes the specified connection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(directconnect_deleteConnectionCmd).Standalone()

	directconnect_deleteConnectionCmd.Flags().String("connection-id", "", "The ID of the connection.")
	directconnect_deleteConnectionCmd.MarkFlagRequired("connection-id")
	directconnectCmd.AddCommand(directconnect_deleteConnectionCmd)
}

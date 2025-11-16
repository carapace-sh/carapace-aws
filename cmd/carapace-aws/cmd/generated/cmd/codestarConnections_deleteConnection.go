package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codestarConnections_deleteConnectionCmd = &cobra.Command{
	Use:   "delete-connection",
	Short: "The connection to be deleted.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codestarConnections_deleteConnectionCmd).Standalone()

	codestarConnections_deleteConnectionCmd.Flags().String("connection-arn", "", "The Amazon Resource Name (ARN) of the connection to be deleted.")
	codestarConnections_deleteConnectionCmd.MarkFlagRequired("connection-arn")
	codestarConnectionsCmd.AddCommand(codestarConnections_deleteConnectionCmd)
}

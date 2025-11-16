package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codestarConnections_getConnectionCmd = &cobra.Command{
	Use:   "get-connection",
	Short: "Returns the connection ARN and details such as status, owner, and provider type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codestarConnections_getConnectionCmd).Standalone()

	codestarConnections_getConnectionCmd.Flags().String("connection-arn", "", "The Amazon Resource Name (ARN) of a connection.")
	codestarConnections_getConnectionCmd.MarkFlagRequired("connection-arn")
	codestarConnectionsCmd.AddCommand(codestarConnections_getConnectionCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeconnections_getConnectionCmd = &cobra.Command{
	Use:   "get-connection",
	Short: "Returns the connection ARN and details such as status, owner, and provider type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeconnections_getConnectionCmd).Standalone()

	codeconnections_getConnectionCmd.Flags().String("connection-arn", "", "The Amazon Resource Name (ARN) of a connection.")
	codeconnections_getConnectionCmd.MarkFlagRequired("connection-arn")
	codeconnectionsCmd.AddCommand(codeconnections_getConnectionCmd)
}

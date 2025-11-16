package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeconnections_deleteConnectionCmd = &cobra.Command{
	Use:   "delete-connection",
	Short: "The connection to be deleted.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeconnections_deleteConnectionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codeconnections_deleteConnectionCmd).Standalone()

		codeconnections_deleteConnectionCmd.Flags().String("connection-arn", "", "The Amazon Resource Name (ARN) of the connection to be deleted.")
		codeconnections_deleteConnectionCmd.MarkFlagRequired("connection-arn")
	})
	codeconnectionsCmd.AddCommand(codeconnections_deleteConnectionCmd)
}

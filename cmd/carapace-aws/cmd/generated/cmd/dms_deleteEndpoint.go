package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_deleteEndpointCmd = &cobra.Command{
	Use:   "delete-endpoint",
	Short: "Deletes the specified endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_deleteEndpointCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dms_deleteEndpointCmd).Standalone()

		dms_deleteEndpointCmd.Flags().String("endpoint-arn", "", "The Amazon Resource Name (ARN) string that uniquely identifies the endpoint.")
		dms_deleteEndpointCmd.MarkFlagRequired("endpoint-arn")
	})
	dmsCmd.AddCommand(dms_deleteEndpointCmd)
}

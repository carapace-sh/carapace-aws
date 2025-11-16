package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3outposts_deleteEndpointCmd = &cobra.Command{
	Use:   "delete-endpoint",
	Short: "Deletes an endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3outposts_deleteEndpointCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3outposts_deleteEndpointCmd).Standalone()

		s3outposts_deleteEndpointCmd.Flags().String("endpoint-id", "", "The ID of the endpoint.")
		s3outposts_deleteEndpointCmd.Flags().String("outpost-id", "", "The ID of the Outposts.")
		s3outposts_deleteEndpointCmd.MarkFlagRequired("endpoint-id")
		s3outposts_deleteEndpointCmd.MarkFlagRequired("outpost-id")
	})
	s3outpostsCmd.AddCommand(s3outposts_deleteEndpointCmd)
}

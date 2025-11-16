package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconvert_describeEndpointsCmd = &cobra.Command{
	Use:   "describe-endpoints",
	Short: "Send a request with an empty body to the regional API endpoint to get your account API endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconvert_describeEndpointsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediaconvert_describeEndpointsCmd).Standalone()

		mediaconvert_describeEndpointsCmd.Flags().String("max-results", "", "Optional.")
		mediaconvert_describeEndpointsCmd.Flags().String("mode", "", "Optional field, defaults to DEFAULT.")
		mediaconvert_describeEndpointsCmd.Flags().String("next-token", "", "Use this string, provided with the response to a previous request, to request the next batch of endpoints.")
	})
	mediaconvertCmd.AddCommand(mediaconvert_describeEndpointsCmd)
}

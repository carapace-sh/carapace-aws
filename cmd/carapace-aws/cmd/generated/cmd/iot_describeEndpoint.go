package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_describeEndpointCmd = &cobra.Command{
	Use:   "describe-endpoint",
	Short: "Returns or creates a unique endpoint specific to the Amazon Web Services account making the call.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_describeEndpointCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_describeEndpointCmd).Standalone()

		iot_describeEndpointCmd.Flags().String("endpoint-type", "", "The endpoint type.")
	})
	iotCmd.AddCommand(iot_describeEndpointCmd)
}

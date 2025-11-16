package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_describeEndpointCmd = &cobra.Command{
	Use:   "describe-endpoint",
	Short: "Gets the properties associated with a specific endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_describeEndpointCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(comprehend_describeEndpointCmd).Standalone()

		comprehend_describeEndpointCmd.Flags().String("endpoint-arn", "", "The Amazon Resource Number (ARN) of the endpoint being described.")
		comprehend_describeEndpointCmd.MarkFlagRequired("endpoint-arn")
	})
	comprehendCmd.AddCommand(comprehend_describeEndpointCmd)
}

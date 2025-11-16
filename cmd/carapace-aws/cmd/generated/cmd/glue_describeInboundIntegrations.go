package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_describeInboundIntegrationsCmd = &cobra.Command{
	Use:   "describe-inbound-integrations",
	Short: "Returns a list of inbound integrations for the specified integration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_describeInboundIntegrationsCmd).Standalone()

	glue_describeInboundIntegrationsCmd.Flags().String("integration-arn", "", "The Amazon Resource Name (ARN) of the integration.")
	glue_describeInboundIntegrationsCmd.Flags().String("marker", "", "A token to specify where to start paginating.")
	glue_describeInboundIntegrationsCmd.Flags().String("max-records", "", "The total number of items to return in the output.")
	glue_describeInboundIntegrationsCmd.Flags().String("target-arn", "", "The Amazon Resource Name (ARN) of the target resource in the integration.")
	glueCmd.AddCommand(glue_describeInboundIntegrationsCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_describeInboundIntegrationsCmd = &cobra.Command{
	Use:   "describe-inbound-integrations",
	Short: "Returns a list of inbound integrations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_describeInboundIntegrationsCmd).Standalone()

	redshift_describeInboundIntegrationsCmd.Flags().String("integration-arn", "", "The Amazon Resource Name (ARN) of the inbound integration.")
	redshift_describeInboundIntegrationsCmd.Flags().String("marker", "", "An optional parameter that specifies the starting point to return a set of response records.")
	redshift_describeInboundIntegrationsCmd.Flags().String("max-records", "", "The maximum number of response records to return in each call.")
	redshift_describeInboundIntegrationsCmd.Flags().String("target-arn", "", "The Amazon Resource Name (ARN) of the target of an inbound integration.")
	redshiftCmd.AddCommand(redshift_describeInboundIntegrationsCmd)
}

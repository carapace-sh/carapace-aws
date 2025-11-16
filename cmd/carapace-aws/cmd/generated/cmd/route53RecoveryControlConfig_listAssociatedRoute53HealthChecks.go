package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53RecoveryControlConfig_listAssociatedRoute53HealthChecksCmd = &cobra.Command{
	Use:   "list-associated-route53-health-checks",
	Short: "Returns an array of all Amazon Route 53 health checks associated with a specific routing control.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53RecoveryControlConfig_listAssociatedRoute53HealthChecksCmd).Standalone()

	route53RecoveryControlConfig_listAssociatedRoute53HealthChecksCmd.Flags().String("max-results", "", "The number of objects that you want to return with this call.")
	route53RecoveryControlConfig_listAssociatedRoute53HealthChecksCmd.Flags().String("next-token", "", "The token that identifies which batch of results you want to see.")
	route53RecoveryControlConfig_listAssociatedRoute53HealthChecksCmd.Flags().String("routing-control-arn", "", "The Amazon Resource Name (ARN) of the routing control.")
	route53RecoveryControlConfig_listAssociatedRoute53HealthChecksCmd.MarkFlagRequired("routing-control-arn")
	route53RecoveryControlConfigCmd.AddCommand(route53RecoveryControlConfig_listAssociatedRoute53HealthChecksCmd)
}

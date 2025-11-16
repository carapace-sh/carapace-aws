package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53RecoveryReadiness_listResourceSetsCmd = &cobra.Command{
	Use:   "list-resource-sets",
	Short: "Lists the resource sets in an account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53RecoveryReadiness_listResourceSetsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53RecoveryReadiness_listResourceSetsCmd).Standalone()

		route53RecoveryReadiness_listResourceSetsCmd.Flags().String("max-results", "", "The number of objects that you want to return with this call.")
		route53RecoveryReadiness_listResourceSetsCmd.Flags().String("next-token", "", "The token that identifies which batch of results you want to see.")
	})
	route53RecoveryReadinessCmd.AddCommand(route53RecoveryReadiness_listResourceSetsCmd)
}

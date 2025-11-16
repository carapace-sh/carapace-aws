package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53RecoveryControlConfig_listControlPanelsCmd = &cobra.Command{
	Use:   "list-control-panels",
	Short: "Returns an array of control panels in an account or in a cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53RecoveryControlConfig_listControlPanelsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53RecoveryControlConfig_listControlPanelsCmd).Standalone()

		route53RecoveryControlConfig_listControlPanelsCmd.Flags().String("cluster-arn", "", "The Amazon Resource Name (ARN) of a cluster.")
		route53RecoveryControlConfig_listControlPanelsCmd.Flags().String("max-results", "", "The number of objects that you want to return with this call.")
		route53RecoveryControlConfig_listControlPanelsCmd.Flags().String("next-token", "", "The token that identifies which batch of results you want to see.")
	})
	route53RecoveryControlConfigCmd.AddCommand(route53RecoveryControlConfig_listControlPanelsCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var drs_startSourceNetworkRecoveryCmd = &cobra.Command{
	Use:   "start-source-network-recovery",
	Short: "Deploy VPC for the specified Source Network and modify launch templates to use this network.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(drs_startSourceNetworkRecoveryCmd).Standalone()

	drs_startSourceNetworkRecoveryCmd.Flags().Bool("deploy-as-new", false, "Don't update existing CloudFormation Stack, recover the network using a new stack.")
	drs_startSourceNetworkRecoveryCmd.Flags().Bool("no-deploy-as-new", false, "Don't update existing CloudFormation Stack, recover the network using a new stack.")
	drs_startSourceNetworkRecoveryCmd.Flags().String("source-networks", "", "The Source Networks that we want to start a Recovery Job for.")
	drs_startSourceNetworkRecoveryCmd.Flags().String("tags", "", "The tags to be associated with the Source Network recovery Job.")
	drs_startSourceNetworkRecoveryCmd.Flag("no-deploy-as-new").Hidden = true
	drs_startSourceNetworkRecoveryCmd.MarkFlagRequired("source-networks")
	drsCmd.AddCommand(drs_startSourceNetworkRecoveryCmd)
}

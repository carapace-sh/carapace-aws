package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var drs_associateSourceNetworkStackCmd = &cobra.Command{
	Use:   "associate-source-network-stack",
	Short: "Associate a Source Network to an existing CloudFormation Stack and modify launch templates to use this network.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(drs_associateSourceNetworkStackCmd).Standalone()

	drs_associateSourceNetworkStackCmd.Flags().String("cfn-stack-name", "", "CloudFormation template to associate with a Source Network.")
	drs_associateSourceNetworkStackCmd.Flags().String("source-network-id", "", "The Source Network ID to associate with CloudFormation template.")
	drs_associateSourceNetworkStackCmd.MarkFlagRequired("cfn-stack-name")
	drs_associateSourceNetworkStackCmd.MarkFlagRequired("source-network-id")
	drsCmd.AddCommand(drs_associateSourceNetworkStackCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var drs_createSourceNetworkCmd = &cobra.Command{
	Use:   "create-source-network",
	Short: "Create a new Source Network resource for a provided VPC ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(drs_createSourceNetworkCmd).Standalone()

	drs_createSourceNetworkCmd.Flags().String("origin-account-id", "", "Account containing the VPC to protect.")
	drs_createSourceNetworkCmd.Flags().String("origin-region", "", "Region containing the VPC to protect.")
	drs_createSourceNetworkCmd.Flags().String("tags", "", "A set of tags to be associated with the Source Network resource.")
	drs_createSourceNetworkCmd.Flags().String("vpc-id", "", "Which VPC ID to protect.")
	drs_createSourceNetworkCmd.MarkFlagRequired("origin-account-id")
	drs_createSourceNetworkCmd.MarkFlagRequired("origin-region")
	drs_createSourceNetworkCmd.MarkFlagRequired("vpc-id")
	drsCmd.AddCommand(drs_createSourceNetworkCmd)
}

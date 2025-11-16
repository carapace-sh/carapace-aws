package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpcLattice_listRulesCmd = &cobra.Command{
	Use:   "list-rules",
	Short: "Lists the rules for the specified listener.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcLattice_listRulesCmd).Standalone()

	vpcLattice_listRulesCmd.Flags().String("listener-identifier", "", "The ID or ARN of the listener.")
	vpcLattice_listRulesCmd.Flags().String("max-results", "", "The maximum number of results to return.")
	vpcLattice_listRulesCmd.Flags().String("next-token", "", "A pagination token for the next page of results.")
	vpcLattice_listRulesCmd.Flags().String("service-identifier", "", "The ID or ARN of the service.")
	vpcLattice_listRulesCmd.MarkFlagRequired("listener-identifier")
	vpcLattice_listRulesCmd.MarkFlagRequired("service-identifier")
	vpcLatticeCmd.AddCommand(vpcLattice_listRulesCmd)
}

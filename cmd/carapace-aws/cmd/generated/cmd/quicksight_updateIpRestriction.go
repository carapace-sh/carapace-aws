package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_updateIpRestrictionCmd = &cobra.Command{
	Use:   "update-ip-restriction",
	Short: "Updates the content and status of IP rules.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_updateIpRestrictionCmd).Standalone()

	quicksight_updateIpRestrictionCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contains the IP rules.")
	quicksight_updateIpRestrictionCmd.Flags().String("enabled", "", "A value that specifies whether IP rules are turned on.")
	quicksight_updateIpRestrictionCmd.Flags().String("ip-restriction-rule-map", "", "A map that describes the updated IP rules with CIDR ranges and descriptions.")
	quicksight_updateIpRestrictionCmd.Flags().String("vpc-endpoint-id-restriction-rule-map", "", "A map of allowed VPC endpoint IDs and their corresponding rule descriptions.")
	quicksight_updateIpRestrictionCmd.Flags().String("vpc-id-restriction-rule-map", "", "A map of VPC IDs and their corresponding rules.")
	quicksight_updateIpRestrictionCmd.MarkFlagRequired("aws-account-id")
	quicksightCmd.AddCommand(quicksight_updateIpRestrictionCmd)
}

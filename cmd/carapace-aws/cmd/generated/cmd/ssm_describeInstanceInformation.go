package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_describeInstanceInformationCmd = &cobra.Command{
	Use:   "describe-instance-information",
	Short: "Provides information about one or more of your managed nodes, including the operating system platform, SSM Agent version, association status, and IP address.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_describeInstanceInformationCmd).Standalone()

	ssm_describeInstanceInformationCmd.Flags().String("filters", "", "One or more filters.")
	ssm_describeInstanceInformationCmd.Flags().String("instance-information-filter-list", "", "This is a legacy method.")
	ssm_describeInstanceInformationCmd.Flags().String("max-results", "", "The maximum number of items to return for this call.")
	ssm_describeInstanceInformationCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
	ssmCmd.AddCommand(ssm_describeInstanceInformationCmd)
}

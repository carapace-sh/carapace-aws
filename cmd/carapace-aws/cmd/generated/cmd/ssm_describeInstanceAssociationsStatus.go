package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_describeInstanceAssociationsStatusCmd = &cobra.Command{
	Use:   "describe-instance-associations-status",
	Short: "The status of the associations for the managed nodes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_describeInstanceAssociationsStatusCmd).Standalone()

	ssm_describeInstanceAssociationsStatusCmd.Flags().String("instance-id", "", "The managed node IDs for which you want association status information.")
	ssm_describeInstanceAssociationsStatusCmd.Flags().String("max-results", "", "The maximum number of items to return for this call.")
	ssm_describeInstanceAssociationsStatusCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
	ssm_describeInstanceAssociationsStatusCmd.MarkFlagRequired("instance-id")
	ssmCmd.AddCommand(ssm_describeInstanceAssociationsStatusCmd)
}

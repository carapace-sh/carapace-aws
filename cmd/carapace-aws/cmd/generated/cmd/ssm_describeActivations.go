package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_describeActivationsCmd = &cobra.Command{
	Use:   "describe-activations",
	Short: "Describes details about the activation, such as the date and time the activation was created, its expiration date, the Identity and Access Management (IAM) role assigned to the managed nodes in the activation, and the number of nodes registered by using this activation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_describeActivationsCmd).Standalone()

	ssm_describeActivationsCmd.Flags().String("filters", "", "A filter to view information about your activations.")
	ssm_describeActivationsCmd.Flags().String("max-results", "", "The maximum number of items to return for this call.")
	ssm_describeActivationsCmd.Flags().String("next-token", "", "A token to start the list.")
	ssmCmd.AddCommand(ssm_describeActivationsCmd)
}

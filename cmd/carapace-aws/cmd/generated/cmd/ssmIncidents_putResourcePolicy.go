package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmIncidents_putResourcePolicyCmd = &cobra.Command{
	Use:   "put-resource-policy",
	Short: "Adds a resource policy to the specified response plan.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmIncidents_putResourcePolicyCmd).Standalone()

	ssmIncidents_putResourcePolicyCmd.Flags().String("policy", "", "Details of the resource policy.")
	ssmIncidents_putResourcePolicyCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the response plan to add the resource policy to.")
	ssmIncidents_putResourcePolicyCmd.MarkFlagRequired("policy")
	ssmIncidents_putResourcePolicyCmd.MarkFlagRequired("resource-arn")
	ssmIncidentsCmd.AddCommand(ssmIncidents_putResourcePolicyCmd)
}

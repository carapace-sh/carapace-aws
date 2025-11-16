package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_putDestinationPolicyCmd = &cobra.Command{
	Use:   "put-destination-policy",
	Short: "Creates or updates an access policy associated with an existing destination.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_putDestinationPolicyCmd).Standalone()

	logs_putDestinationPolicyCmd.Flags().String("access-policy", "", "An IAM policy document that authorizes cross-account users to deliver their log events to the associated destination.")
	logs_putDestinationPolicyCmd.Flags().String("destination-name", "", "A name for an existing destination.")
	logs_putDestinationPolicyCmd.Flags().String("force-update", "", "Specify true if you are updating an existing destination policy to grant permission to an organization ID instead of granting permission to individual Amazon Web Services accounts.")
	logs_putDestinationPolicyCmd.MarkFlagRequired("access-policy")
	logs_putDestinationPolicyCmd.MarkFlagRequired("destination-name")
	logsCmd.AddCommand(logs_putDestinationPolicyCmd)
}

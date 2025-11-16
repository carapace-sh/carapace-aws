package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var shield_associateDrtroleCmd = &cobra.Command{
	Use:   "associate-drtrole",
	Short: "Authorizes the Shield Response Team (SRT) using the specified role, to access your Amazon Web Services account to assist with DDoS attack mitigation during potential attacks.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(shield_associateDrtroleCmd).Standalone()

	shield_associateDrtroleCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of the role the SRT will use to access your Amazon Web Services account.")
	shield_associateDrtroleCmd.MarkFlagRequired("role-arn")
	shieldCmd.AddCommand(shield_associateDrtroleCmd)
}

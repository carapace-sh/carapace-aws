package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmContacts_describeEngagementCmd = &cobra.Command{
	Use:   "describe-engagement",
	Short: "Incident Manager uses engagements to engage contacts and escalation plans during an incident.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmContacts_describeEngagementCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssmContacts_describeEngagementCmd).Standalone()

		ssmContacts_describeEngagementCmd.Flags().String("engagement-id", "", "The Amazon Resource Name (ARN) of the engagement you want the details of.")
		ssmContacts_describeEngagementCmd.MarkFlagRequired("engagement-id")
	})
	ssmContactsCmd.AddCommand(ssmContacts_describeEngagementCmd)
}

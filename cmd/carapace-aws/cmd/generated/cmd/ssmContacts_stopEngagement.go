package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmContacts_stopEngagementCmd = &cobra.Command{
	Use:   "stop-engagement",
	Short: "Stops an engagement before it finishes the final stage of the escalation plan or engagement plan.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmContacts_stopEngagementCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssmContacts_stopEngagementCmd).Standalone()

		ssmContacts_stopEngagementCmd.Flags().String("engagement-id", "", "The Amazon Resource Name (ARN) of the engagement.")
		ssmContacts_stopEngagementCmd.Flags().String("reason", "", "The reason that you're stopping the engagement.")
		ssmContacts_stopEngagementCmd.MarkFlagRequired("engagement-id")
	})
	ssmContactsCmd.AddCommand(ssmContacts_stopEngagementCmd)
}

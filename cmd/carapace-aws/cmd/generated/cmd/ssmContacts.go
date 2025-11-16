package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmContactsCmd = &cobra.Command{
	Use:   "ssm-contacts",
	Short: "Systems Manager Incident Manager is an incident management console designed to help users mitigate and recover from incidents affecting their Amazon Web Services-hosted applications.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmContactsCmd).Standalone()

	rootCmd.AddCommand(ssmContactsCmd)
}

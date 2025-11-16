package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var grafana_createWorkspaceServiceAccountTokenCmd = &cobra.Command{
	Use:   "create-workspace-service-account-token",
	Short: "Creates a token that can be used to authenticate and authorize Grafana HTTP API operations for the given [workspace service account](https://docs.aws.amazon.com/grafana/latest/userguide/service-accounts.html).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(grafana_createWorkspaceServiceAccountTokenCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(grafana_createWorkspaceServiceAccountTokenCmd).Standalone()

		grafana_createWorkspaceServiceAccountTokenCmd.Flags().String("name", "", "A name for the token to create.")
		grafana_createWorkspaceServiceAccountTokenCmd.Flags().String("seconds-to-live", "", "Sets how long the token will be valid, in seconds.")
		grafana_createWorkspaceServiceAccountTokenCmd.Flags().String("service-account-id", "", "The ID of the service account for which to create a token.")
		grafana_createWorkspaceServiceAccountTokenCmd.Flags().String("workspace-id", "", "The ID of the workspace the service account resides within.")
		grafana_createWorkspaceServiceAccountTokenCmd.MarkFlagRequired("name")
		grafana_createWorkspaceServiceAccountTokenCmd.MarkFlagRequired("seconds-to-live")
		grafana_createWorkspaceServiceAccountTokenCmd.MarkFlagRequired("service-account-id")
		grafana_createWorkspaceServiceAccountTokenCmd.MarkFlagRequired("workspace-id")
	})
	grafanaCmd.AddCommand(grafana_createWorkspaceServiceAccountTokenCmd)
}

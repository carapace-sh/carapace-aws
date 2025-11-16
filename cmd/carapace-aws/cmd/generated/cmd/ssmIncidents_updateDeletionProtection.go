package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmIncidents_updateDeletionProtectionCmd = &cobra.Command{
	Use:   "update-deletion-protection",
	Short: "Update deletion protection to either allow or deny deletion of the final Region in a replication set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmIncidents_updateDeletionProtectionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssmIncidents_updateDeletionProtectionCmd).Standalone()

		ssmIncidents_updateDeletionProtectionCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the replication set to update.")
		ssmIncidents_updateDeletionProtectionCmd.Flags().String("client-token", "", "A token that ensures that the operation is called only once with the specified details.")
		ssmIncidents_updateDeletionProtectionCmd.Flags().Bool("deletion-protected", false, "Specifies if deletion protection is turned on or off in your account.")
		ssmIncidents_updateDeletionProtectionCmd.Flags().Bool("no-deletion-protected", false, "Specifies if deletion protection is turned on or off in your account.")
		ssmIncidents_updateDeletionProtectionCmd.MarkFlagRequired("arn")
		ssmIncidents_updateDeletionProtectionCmd.MarkFlagRequired("deletion-protected")
		ssmIncidents_updateDeletionProtectionCmd.Flag("no-deletion-protected").Hidden = true
		ssmIncidents_updateDeletionProtectionCmd.MarkFlagRequired("no-deletion-protected")
	})
	ssmIncidentsCmd.AddCommand(ssmIncidents_updateDeletionProtectionCmd)
}

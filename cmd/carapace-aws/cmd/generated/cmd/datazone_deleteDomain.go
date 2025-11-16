package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_deleteDomainCmd = &cobra.Command{
	Use:   "delete-domain",
	Short: "Deletes a Amazon DataZone domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_deleteDomainCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_deleteDomainCmd).Standalone()

		datazone_deleteDomainCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that is provided to ensure the idempotency of the request.")
		datazone_deleteDomainCmd.Flags().String("identifier", "", "The identifier of the Amazon Web Services domain that is to be deleted.")
		datazone_deleteDomainCmd.Flags().Bool("no-skip-deletion-check", false, "Specifies the optional flag to delete all child entities within the domain.")
		datazone_deleteDomainCmd.Flags().Bool("skip-deletion-check", false, "Specifies the optional flag to delete all child entities within the domain.")
		datazone_deleteDomainCmd.MarkFlagRequired("identifier")
		datazone_deleteDomainCmd.Flag("no-skip-deletion-check").Hidden = true
	})
	datazoneCmd.AddCommand(datazone_deleteDomainCmd)
}

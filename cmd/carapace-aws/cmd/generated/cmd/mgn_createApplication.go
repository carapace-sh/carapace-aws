package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgn_createApplicationCmd = &cobra.Command{
	Use:   "create-application",
	Short: "Create application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgn_createApplicationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mgn_createApplicationCmd).Standalone()

		mgn_createApplicationCmd.Flags().String("account-id", "", "Account ID.")
		mgn_createApplicationCmd.Flags().String("description", "", "Application description.")
		mgn_createApplicationCmd.Flags().String("name", "", "Application name.")
		mgn_createApplicationCmd.Flags().String("tags", "", "Application tags.")
		mgn_createApplicationCmd.MarkFlagRequired("name")
	})
	mgnCmd.AddCommand(mgn_createApplicationCmd)
}

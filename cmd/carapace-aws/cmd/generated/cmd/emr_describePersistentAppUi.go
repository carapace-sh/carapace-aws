package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emr_describePersistentAppUiCmd = &cobra.Command{
	Use:   "describe-persistent-app-ui",
	Short: "Describes a persistent application user interface.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emr_describePersistentAppUiCmd).Standalone()

	emr_describePersistentAppUiCmd.Flags().String("persistent-app-uiid", "", "The identifier for the persistent application user interface.")
	emr_describePersistentAppUiCmd.MarkFlagRequired("persistent-app-uiid")
	emrCmd.AddCommand(emr_describePersistentAppUiCmd)
}

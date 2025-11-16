package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auditmanager_listControlsCmd = &cobra.Command{
	Use:   "list-controls",
	Short: "Returns a list of controls from Audit Manager.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auditmanager_listControlsCmd).Standalone()

	auditmanager_listControlsCmd.Flags().String("control-catalog-id", "", "A filter that narrows the list of controls to a specific resource from the Amazon Web Services Control Catalog.")
	auditmanager_listControlsCmd.Flags().String("control-type", "", "A filter that narrows the list of controls to a specific type.")
	auditmanager_listControlsCmd.Flags().String("max-results", "", "The maximum number of results on a page or for an API request call.")
	auditmanager_listControlsCmd.Flags().String("next-token", "", "The pagination token that's used to fetch the next set of results.")
	auditmanager_listControlsCmd.MarkFlagRequired("control-type")
	auditmanagerCmd.AddCommand(auditmanager_listControlsCmd)
}

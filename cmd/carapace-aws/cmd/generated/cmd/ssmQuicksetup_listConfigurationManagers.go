package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmQuicksetup_listConfigurationManagersCmd = &cobra.Command{
	Use:   "list-configuration-managers",
	Short: "Returns Quick Setup configuration managers.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmQuicksetup_listConfigurationManagersCmd).Standalone()

	ssmQuicksetup_listConfigurationManagersCmd.Flags().String("filters", "", "Filters the results returned by the request.")
	ssmQuicksetup_listConfigurationManagersCmd.Flags().String("max-items", "", "Specifies the maximum number of configuration managers that are returned by the request.")
	ssmQuicksetup_listConfigurationManagersCmd.Flags().String("starting-token", "", "The token to use when requesting a specific set of items from a list.")
	ssmQuicksetupCmd.AddCommand(ssmQuicksetup_listConfigurationManagersCmd)
}

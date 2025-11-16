package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector_listRulesPackagesCmd = &cobra.Command{
	Use:   "list-rules-packages",
	Short: "Lists all available Amazon Inspector rules packages.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector_listRulesPackagesCmd).Standalone()

	inspector_listRulesPackagesCmd.Flags().String("max-results", "", "You can use this parameter to indicate the maximum number of items you want in the response.")
	inspector_listRulesPackagesCmd.Flags().String("next-token", "", "You can use this parameter when paginating results.")
	inspectorCmd.AddCommand(inspector_listRulesPackagesCmd)
}

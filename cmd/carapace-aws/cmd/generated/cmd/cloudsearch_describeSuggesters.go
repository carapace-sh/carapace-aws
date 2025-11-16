package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudsearch_describeSuggestersCmd = &cobra.Command{
	Use:   "describe-suggesters",
	Short: "Gets the suggesters configured for a domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudsearch_describeSuggestersCmd).Standalone()

	cloudsearch_describeSuggestersCmd.Flags().Bool("deployed", false, "Whether to display the deployed configuration (`true`) or include any pending changes (`false`).")
	cloudsearch_describeSuggestersCmd.Flags().String("domain-name", "", "The name of the domain you want to describe.")
	cloudsearch_describeSuggestersCmd.Flags().Bool("no-deployed", false, "Whether to display the deployed configuration (`true`) or include any pending changes (`false`).")
	cloudsearch_describeSuggestersCmd.Flags().String("suggester-names", "", "The suggesters you want to describe.")
	cloudsearch_describeSuggestersCmd.MarkFlagRequired("domain-name")
	cloudsearch_describeSuggestersCmd.Flag("no-deployed").Hidden = true
	cloudsearchCmd.AddCommand(cloudsearch_describeSuggestersCmd)
}

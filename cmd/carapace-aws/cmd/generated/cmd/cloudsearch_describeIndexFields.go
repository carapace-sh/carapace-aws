package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudsearch_describeIndexFieldsCmd = &cobra.Command{
	Use:   "describe-index-fields",
	Short: "Gets information about the index fields configured for the search domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudsearch_describeIndexFieldsCmd).Standalone()

	cloudsearch_describeIndexFieldsCmd.Flags().Bool("deployed", false, "Whether to display the deployed configuration (`true`) or include any pending changes (`false`).")
	cloudsearch_describeIndexFieldsCmd.Flags().String("domain-name", "", "The name of the domain you want to describe.")
	cloudsearch_describeIndexFieldsCmd.Flags().String("field-names", "", "A list of the index fields you want to describe.")
	cloudsearch_describeIndexFieldsCmd.Flags().Bool("no-deployed", false, "Whether to display the deployed configuration (`true`) or include any pending changes (`false`).")
	cloudsearch_describeIndexFieldsCmd.MarkFlagRequired("domain-name")
	cloudsearch_describeIndexFieldsCmd.Flag("no-deployed").Hidden = true
	cloudsearchCmd.AddCommand(cloudsearch_describeIndexFieldsCmd)
}

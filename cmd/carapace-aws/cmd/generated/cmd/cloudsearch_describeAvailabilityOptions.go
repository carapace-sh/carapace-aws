package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudsearch_describeAvailabilityOptionsCmd = &cobra.Command{
	Use:   "describe-availability-options",
	Short: "Gets the availability options configured for a domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudsearch_describeAvailabilityOptionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudsearch_describeAvailabilityOptionsCmd).Standalone()

		cloudsearch_describeAvailabilityOptionsCmd.Flags().Bool("deployed", false, "Whether to display the deployed configuration (`true`) or include any pending changes (`false`).")
		cloudsearch_describeAvailabilityOptionsCmd.Flags().String("domain-name", "", "The name of the domain you want to describe.")
		cloudsearch_describeAvailabilityOptionsCmd.Flags().Bool("no-deployed", false, "Whether to display the deployed configuration (`true`) or include any pending changes (`false`).")
		cloudsearch_describeAvailabilityOptionsCmd.MarkFlagRequired("domain-name")
		cloudsearch_describeAvailabilityOptionsCmd.Flag("no-deployed").Hidden = true
	})
	cloudsearchCmd.AddCommand(cloudsearch_describeAvailabilityOptionsCmd)
}

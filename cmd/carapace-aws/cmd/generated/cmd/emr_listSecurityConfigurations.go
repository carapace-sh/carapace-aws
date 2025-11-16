package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emr_listSecurityConfigurationsCmd = &cobra.Command{
	Use:   "list-security-configurations",
	Short: "Lists all the security configurations visible to this account, providing their creation dates and times, and their names.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emr_listSecurityConfigurationsCmd).Standalone()

	emr_listSecurityConfigurationsCmd.Flags().String("marker", "", "The pagination token that indicates the set of results to retrieve.")
	emrCmd.AddCommand(emr_listSecurityConfigurationsCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_batchGetCustomDataIdentifiersCmd = &cobra.Command{
	Use:   "batch-get-custom-data-identifiers",
	Short: "Retrieves information about one or more custom data identifiers.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_batchGetCustomDataIdentifiersCmd).Standalone()

	macie2_batchGetCustomDataIdentifiersCmd.Flags().String("ids", "", "An array of custom data identifier IDs, one for each custom data identifier to retrieve information about.")
	macie2Cmd.AddCommand(macie2_batchGetCustomDataIdentifiersCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_createSampleFindingsCmd = &cobra.Command{
	Use:   "create-sample-findings",
	Short: "Creates sample findings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_createSampleFindingsCmd).Standalone()

	macie2_createSampleFindingsCmd.Flags().String("finding-types", "", "An array of finding types, one for each type of sample finding to create.")
	macie2Cmd.AddCommand(macie2_createSampleFindingsCmd)
}

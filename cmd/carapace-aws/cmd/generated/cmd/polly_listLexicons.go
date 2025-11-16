package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var polly_listLexiconsCmd = &cobra.Command{
	Use:   "list-lexicons",
	Short: "Returns a list of pronunciation lexicons stored in an Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(polly_listLexiconsCmd).Standalone()

	polly_listLexiconsCmd.Flags().String("next-token", "", "An opaque pagination token returned from previous `ListLexicons` operation.")
	pollyCmd.AddCommand(polly_listLexiconsCmd)
}

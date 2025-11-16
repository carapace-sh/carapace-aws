package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityIr_listCommentsCmd = &cobra.Command{
	Use:   "list-comments",
	Short: "Returns comments for a designated case.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityIr_listCommentsCmd).Standalone()

	securityIr_listCommentsCmd.Flags().String("case-id", "", "Required element for ListComments to designate the case to query.")
	securityIr_listCommentsCmd.Flags().String("max-results", "", "Optional element for ListComments to limit the number of responses.")
	securityIr_listCommentsCmd.Flags().String("next-token", "", "An optional string that, if supplied, must be copied from the output of a previous call to ListComments.")
	securityIr_listCommentsCmd.MarkFlagRequired("case-id")
	securityIrCmd.AddCommand(securityIr_listCommentsCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmContacts_listRotationsCmd = &cobra.Command{
	Use:   "list-rotations",
	Short: "Retrieves a list of on-call rotations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmContacts_listRotationsCmd).Standalone()

	ssmContacts_listRotationsCmd.Flags().String("max-results", "", "The maximum number of items to return for this call.")
	ssmContacts_listRotationsCmd.Flags().String("next-token", "", "A token to start the list.")
	ssmContacts_listRotationsCmd.Flags().String("rotation-name-prefix", "", "A filter to include rotations in list results based on their common prefix.")
	ssmContactsCmd.AddCommand(ssmContacts_listRotationsCmd)
}

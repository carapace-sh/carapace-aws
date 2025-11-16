package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecommit_listRepositoriesCmd = &cobra.Command{
	Use:   "list-repositories",
	Short: "Gets information about one or more repositories.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecommit_listRepositoriesCmd).Standalone()

	codecommit_listRepositoriesCmd.Flags().String("next-token", "", "An enumeration token that allows the operation to batch the results of the operation.")
	codecommit_listRepositoriesCmd.Flags().String("order", "", "The order in which to sort the results of a list repositories operation.")
	codecommit_listRepositoriesCmd.Flags().String("sort-by", "", "The criteria used to sort the results of a list repositories operation.")
	codecommitCmd.AddCommand(codecommit_listRepositoriesCmd)
}

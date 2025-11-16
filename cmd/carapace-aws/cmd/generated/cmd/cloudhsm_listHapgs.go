package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudhsm_listHapgsCmd = &cobra.Command{
	Use:   "list-hapgs",
	Short: "This is documentation for **AWS CloudHSM Classic**.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudhsm_listHapgsCmd).Standalone()

	cloudhsm_listHapgsCmd.Flags().String("next-token", "", "The `NextToken` value from a previous call to `ListHapgs`.")
	cloudhsmCmd.AddCommand(cloudhsm_listHapgsCmd)
}

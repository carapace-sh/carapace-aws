package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_startAssociationsOnceCmd = &cobra.Command{
	Use:   "start-associations-once",
	Short: "Runs an association immediately and only one time.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_startAssociationsOnceCmd).Standalone()

	ssm_startAssociationsOnceCmd.Flags().String("association-ids", "", "The association IDs that you want to run immediately and only one time.")
	ssm_startAssociationsOnceCmd.MarkFlagRequired("association-ids")
	ssmCmd.AddCommand(ssm_startAssociationsOnceCmd)
}

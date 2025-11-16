package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgn_changeServerLifeCycleStateCmd = &cobra.Command{
	Use:   "change-server-life-cycle-state",
	Short: "Allows the user to set the SourceServer.LifeCycle.state property for specific Source Server IDs to one of the following: READY\\_FOR\\_TEST or READY\\_FOR\\_CUTOVER.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgn_changeServerLifeCycleStateCmd).Standalone()

	mgn_changeServerLifeCycleStateCmd.Flags().String("account-id", "", "The request to change the source server migration account ID.")
	mgn_changeServerLifeCycleStateCmd.Flags().String("life-cycle", "", "The request to change the source server migration lifecycle state.")
	mgn_changeServerLifeCycleStateCmd.Flags().String("source-server-id", "", "The request to change the source server migration lifecycle state by source server ID.")
	mgn_changeServerLifeCycleStateCmd.MarkFlagRequired("life-cycle")
	mgn_changeServerLifeCycleStateCmd.MarkFlagRequired("source-server-id")
	mgnCmd.AddCommand(mgn_changeServerLifeCycleStateCmd)
}

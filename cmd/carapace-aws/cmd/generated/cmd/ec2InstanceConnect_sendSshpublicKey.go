package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2InstanceConnect_sendSshpublicKeyCmd = &cobra.Command{
	Use:   "send-sshpublic-key",
	Short: "Pushes an SSH public key to the specified EC2 instance for use by the specified user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2InstanceConnect_sendSshpublicKeyCmd).Standalone()

	ec2InstanceConnect_sendSshpublicKeyCmd.Flags().String("availability-zone", "", "The Availability Zone in which the EC2 instance was launched.")
	ec2InstanceConnect_sendSshpublicKeyCmd.Flags().String("instance-id", "", "The ID of the EC2 instance.")
	ec2InstanceConnect_sendSshpublicKeyCmd.Flags().String("instance-osuser", "", "The OS user on the EC2 instance for whom the key can be used to authenticate.")
	ec2InstanceConnect_sendSshpublicKeyCmd.Flags().String("sshpublic-key", "", "The public key material.")
	ec2InstanceConnect_sendSshpublicKeyCmd.MarkFlagRequired("instance-id")
	ec2InstanceConnect_sendSshpublicKeyCmd.MarkFlagRequired("instance-osuser")
	ec2InstanceConnect_sendSshpublicKeyCmd.MarkFlagRequired("sshpublic-key")
	ec2InstanceConnectCmd.AddCommand(ec2InstanceConnect_sendSshpublicKeyCmd)
}

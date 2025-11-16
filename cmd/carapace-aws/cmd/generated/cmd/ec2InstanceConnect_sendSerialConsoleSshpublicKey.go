package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2InstanceConnect_sendSerialConsoleSshpublicKeyCmd = &cobra.Command{
	Use:   "send-serial-console-sshpublic-key",
	Short: "Pushes an SSH public key to the specified EC2 instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2InstanceConnect_sendSerialConsoleSshpublicKeyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2InstanceConnect_sendSerialConsoleSshpublicKeyCmd).Standalone()

		ec2InstanceConnect_sendSerialConsoleSshpublicKeyCmd.Flags().String("instance-id", "", "The ID of the EC2 instance.")
		ec2InstanceConnect_sendSerialConsoleSshpublicKeyCmd.Flags().String("serial-port", "", "The serial port of the EC2 instance.")
		ec2InstanceConnect_sendSerialConsoleSshpublicKeyCmd.Flags().String("sshpublic-key", "", "The public key material.")
		ec2InstanceConnect_sendSerialConsoleSshpublicKeyCmd.MarkFlagRequired("instance-id")
		ec2InstanceConnect_sendSerialConsoleSshpublicKeyCmd.MarkFlagRequired("sshpublic-key")
	})
	ec2InstanceConnectCmd.AddCommand(ec2InstanceConnect_sendSerialConsoleSshpublicKeyCmd)
}

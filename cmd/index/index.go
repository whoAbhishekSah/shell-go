package index

const (
	ExitCmd = "exit"
	EchoCmd = "echo"
	TypeCmd = "type"
)

var ValidCommands = []string{ExitCmd, EchoCmd, TypeCmd}
var ShellBuiltInCommands = []string{ExitCmd, EchoCmd, TypeCmd}

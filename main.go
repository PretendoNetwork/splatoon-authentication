package main

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	nexauth "github.com/PretendoNetwork/nex-protocols-common-go/authentication"
)

var nexServer *nex.Server

func main() {
	nexServer = nex.NewServer()
	nexServer.SetPrudpVersion(1)
	nexServer.SetNexVersion(30500)
	nexServer.SetKerberosKeySize(32)
	nexServer.SetAccessKey("6f599f81")
	nexServer.SetPingTimeout(20)

	nexServer.On("Data", func(packet *nex.PacketV1) {
		request := packet.RMCRequest()

		fmt.Println("==Splatoon - Auth==")
		fmt.Printf("Protocol ID: %#v\n", request.ProtocolID())
		fmt.Printf("Method ID: %#v\n", request.MethodID())
		fmt.Println("===============")
	})

	authenticationServer := nexauth.NewCommonAuthenticationProtocol(nexServer)
	authenticationServer.SetSecureStationURL(nex.NewStationURL("prudps:/address=159.203.102.56;port=61003;CID=1;PID=2;sid=1;stream=10;type=2"))
	authenticationServer.SetBuildName("Pretendo Splatoon - Commit")
	authenticationServer.PasswordFromPID(getNEXAccountByPID)
	_ = authenticationServer

	nexServer.Listen(":61002")
}

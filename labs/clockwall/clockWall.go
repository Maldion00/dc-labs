import (
	"fmt"
	"os"
	"strings"
	"net"
)

func main() {
	
	Channel := make(chan string, len(os.Args[1:]))
	c := make([]string, len(os.Args[1:]))

	for num, part := range os.Args[1:] {

		div := strings.Split(part, "=")
		c[num] = div[1]

	}
	for _, con := range c{

		conn, err := net.Dial("tcp", con)
		if err != nil {
			fmt.Println("error connecting", err.Error())
			os.Exit(1)
		}

		go readBuffer(conn, Channel)
	}

	for i := range Channel{
		fmt.Printf("\r%v", i)
	}

	close(Channel)
}

func readBuffer(a net.Conn, channel chan string){
	defer c.Close()

	buffer := make([]byte, 1024)
	bytes, err := c.Read(buffer)
	if err != nil{
		fmt.Println("error" err.Error())
		
		os.Exit(2)
	}
	if bytes > 0{
		ch <- string(buffer[:])
	}
}
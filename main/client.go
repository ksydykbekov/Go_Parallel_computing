/*Example #2 */
// Clock1 is a TCP server that periodically writes the time.
package main
import (

	"log"
	"net"

	"os"
)


func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	go mustCopy(os.Stdout, conn)
	mustCopy(conn, os.Stdin)
}

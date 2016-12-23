/*Example #2 */
// Clock1 is a TCP server that periodically writes the time.
package main
import (


	"net"
	"time"
	"bufio"
)


func handleConn1(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		go echo(c, input.Text(), 1*time.Second)
	}
	// NOTE: ignoring potential errors from input.Err()
	c.Close()
}



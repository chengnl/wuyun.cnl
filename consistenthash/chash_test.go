package consistenthash

import "testing"
import "fmt"

type IPNode struct {
	ip   string
	port int
}

func NewNode(ip string, port int) *IPNode {
	return &IPNode{ip: ip, port: port}
}
func (node *IPNode) String() string {
	return fmt.Sprintf("%s-%d", node.ip, node.port)
}
func TestString(t *testing.T) {
	nodes := make([]interface{}, 3)
	nodes[0] = "192.168.0.1-1001"
	nodes[1] = "192.168.0.2-1002"
	nodes[2] = "192.168.0.3-1003"
	chash := NewConsistenHash(6, nodes)

	fmt.Println(chash.members)

	fmt.Println(chash.GetNode("23457"))
	fmt.Println(chash.GetNode("testfdsafdsafdsa"))
	fmt.Println(chash.GetNode("9876782222"))

	fmt.Println("DeleteNode.......192.168.0.3-1003")
	chash.DeleteNode("192.168.0.3-1003")

	fmt.Println(chash.GetNode("23457"))
	fmt.Println(chash.GetNode("testfdsafdsafdsa"))
	fmt.Println(chash.GetNode("9876782222"))
	fmt.Println("AddNode.......192.168.0.3-1003")
	chash.AddNode("192.168.0.3-1003")

	fmt.Println(chash.GetNode("23457"))
	fmt.Println(chash.GetNode("testfdsafdsafdsa"))
	fmt.Println(chash.GetNode("9876782222"))
}

func TestInt(t *testing.T) {
	nodes := make([]interface{}, 3)
	nodes[0] = 1001
	nodes[1] = 1002
	nodes[2] = 1003
	chash := NewConsistenHash(6, nodes)

	fmt.Println(chash.members)

	fmt.Println(chash.GetNode("23457"))
	fmt.Println(chash.GetNode("testfdsafdsafdsa"))
	fmt.Println(chash.GetNode("9876782222"))

	fmt.Println("DeleteNode.......1003")
	chash.DeleteNode(1003)

	fmt.Println(chash.GetNode("23457"))
	fmt.Println(chash.GetNode("testfdsafdsafdsa"))
	fmt.Println(chash.GetNode("9876782222"))
	fmt.Println("AddNode.......1003")
	chash.AddNode(1003)

	fmt.Println(chash.GetNode("23457"))
	fmt.Println(chash.GetNode("testfdsafdsafdsa"))
	fmt.Println(chash.GetNode("9876782222"))
}

func TestNode(t *testing.T) {
	nodes := []interface{}{
		NewNode("192.168.0.1", 1001),
		NewNode("192.168.0.2", 1002),
		NewNode("192.168.0.3", 1003),
	}
	chash := NewConsistenHash(6, nodes)

	fmt.Println(chash.members)

	fmt.Println(chash.GetNode("23457"))
	fmt.Println(chash.GetNode("test"))
	fmt.Println(chash.GetNode("987678"))
	fmt.Println("DeleteNode.......192.168.0.1-1001")
	chash.DeleteNode(NewNode("192.168.0.1", 1001))

	fmt.Println(chash.GetNode("23457"))
	fmt.Println(chash.GetNode("test"))
	fmt.Println(chash.GetNode("987678"))
	fmt.Println("AddNode.......192.168.0.1-1001")
	chash.AddNode(NewNode("192.168.0.1", 1001))

	fmt.Println(chash.GetNode("23457"))
	fmt.Println(chash.GetNode("test"))
	fmt.Println(chash.GetNode("987678"))
}

type Server struct {
	name string
	addr string
	port int
}

func TestInterface(t *testing.T) {
	nodes := []interface{}{
		Server{name: "server1", addr: "192.168.0.1", port: 1001},
		Server{name: "server2", addr: "192.168.0.2", port: 1002},
		Server{name: "server3", addr: "192.168.0.3", port: 1003},
	}
	chash := NewConsistenHash(6, nodes)

	fmt.Println(chash.members)

	fmt.Println(chash.GetNode("23457"))
	fmt.Println(chash.GetNode("test2222"))
	fmt.Println(chash.GetNode("987678fdasfdsa"))
	fmt.Println(chash.GetNode("9999"))
	fmt.Println("DeleteNode.......192.168.0.1-1001")
	chash.DeleteNode(Server{name: "server1", addr: "192.168.0.1", port: 1001})

	fmt.Println(chash.GetNode("23457"))
	fmt.Println(chash.GetNode("test2222"))
	fmt.Println(chash.GetNode("987678fdasfdsa"))
	fmt.Println(chash.GetNode("9999"))
	fmt.Println("AddNode.......192.168.0.1-1001")
	chash.AddNode(Server{name: "server1", addr: "192.168.0.1", port: 1001})

	fmt.Println(chash.GetNode("23457"))
	fmt.Println(chash.GetNode("test2222"))
	fmt.Println(chash.GetNode("987678fdasfdsa"))
	fmt.Println(chash.GetNode("9999"))
}
func TestHFunc(t *testing.T) {
	nodes := []interface{}{
		Server{name: "server1", addr: "192.168.0.1", port: 1001},
		Server{name: "server2", addr: "192.168.0.2", port: 1002},
		Server{name: "server3", addr: "192.168.0.3", port: 1003},
	}
	chash := NewConsistenHashWithHFunc(6, func(hashString string) int {
		//FNV1A_64_HASH
		var FNV_64_INIT uint64 = 0xcbf29ce484222325
		var FNV_64_PRIME uint64 = 0x100000001b3
		hashCode := FNV_64_INIT
		for _, r := range hashString {
			hashCode = (hashCode ^ uint64(r)) * FNV_64_PRIME
		}
		//fmt.Printf("hashCode:%d\n", hashCode&0xffffffff)
		return int(hashCode & 0xffffffff)
	}, nodes)

	fmt.Println(chash.members)

	fmt.Println(chash.GetNode("23457"))
	fmt.Println(chash.GetNode("test2222"))
	fmt.Println(chash.GetNode("987678fdasfdsa"))
	fmt.Println(chash.GetNode("9999"))
	fmt.Println("DeleteNode.......192.168.0.1-1001")
	chash.DeleteNode(Server{name: "server1", addr: "192.168.0.1", port: 1001})

	fmt.Println(chash.GetNode("23457"))
	fmt.Println(chash.GetNode("test2222"))
	fmt.Println(chash.GetNode("987678fdasfdsa"))
	fmt.Println(chash.GetNode("9999"))
	fmt.Println("AddNode.......192.168.0.1-1001")
	chash.AddNode(Server{name: "server1", addr: "192.168.0.1", port: 1001})

	fmt.Println(chash.GetNode("23457"))
	fmt.Println(chash.GetNode("test2222"))
	fmt.Println(chash.GetNode("987678fdasfdsa"))
	fmt.Println(chash.GetNode("9999"))

	server := chash.GetNode("9999")
	fmt.Println(server.(Server).name)
}
